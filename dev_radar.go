package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"regexp"
	"runtime"
	"sort"
	"strings"
	"sync"
	"time"
)

type ServiceInfo struct {
	Port       int
	Service    string
	Technology string
	Title      string
	Server     string
	Framework  string
	Status     string
	URL        string
}

type HostResult struct {
	IP       string
	Hostname string
	Services []ServiceInfo
}

func main() {
	fmt.Println("Development Server Network Scanner")
	fmt.Println("==================================")

	// Get local network range
	localIP, subnet, err := getLocalNetwork()
	if err != nil {
		fmt.Printf("Error getting local network: %v\n", err)
		return
	}

	fmt.Printf("Local IP: %s\n", localIP)
	fmt.Printf("Scanning network: %s\n", subnet)
	fmt.Println()

	// Discover hosts
	fmt.Println("Discovering hosts...")
	hosts := discoverHosts(subnet)

	if len(hosts) == 0 {
		fmt.Println("No hosts found on the network")
		return
	}

	fmt.Printf("Found %d active hosts\n", len(hosts))
	fmt.Println()

	// Scan for development servers
	fmt.Println("Scanning for development servers...")
	results := scanDevelopmentServers(hosts)

	// Display results
	displayResults(results)
}

func getLocalNetwork() (string, string, error) {
	// Get local IP address
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "", "", err
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	localIP := localAddr.IP.String()

	// Get network interfaces
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", "", err
	}

	for _, iface := range interfaces {
		if iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagLoopback != 0 {
			continue
		}

		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}

		for _, addr := range addrs {
			ipNet, ok := addr.(*net.IPNet)
			if !ok {
				continue
			}

			if ipNet.IP.String() == localIP {
				network := ipNet.IP.Mask(ipNet.Mask)
				subnet := fmt.Sprintf("%s/%d", network.String(), getMaskBits(ipNet.Mask))
				return localIP, subnet, nil
			}
		}
	}

	// Fallback: assume /24 network
	parts := strings.Split(localIP, ".")
	if len(parts) == 4 {
		subnet := fmt.Sprintf("%s.%s.%s.0/24", parts[0], parts[1], parts[2])
		return localIP, subnet, nil
	}

	return "", "", fmt.Errorf("could not determine network range")
}

func getMaskBits(mask net.IPMask) int {
	ones, _ := mask.Size()
	return ones
}

func discoverHosts(subnet string) []string {
	_, ipNet, err := net.ParseCIDR(subnet)
	if err != nil {
		fmt.Printf("Error parsing CIDR: %v\n", err)
		return nil
	}

	var hosts []string
	var wg sync.WaitGroup
	var mu sync.Mutex

	semaphore := make(chan struct{}, 100)

	for ip := ipNet.IP.Mask(ipNet.Mask); ipNet.Contains(ip); inc(ip) {
		wg.Add(1)
		go func(ip string) {
			defer wg.Done()
			semaphore <- struct{}{}
			defer func() { <-semaphore }()

			if isHostAlive(ip) {
				mu.Lock()
				hosts = append(hosts, ip)
				mu.Unlock()
			}
		}(ip.String())
	}

	wg.Wait()
	sort.Strings(hosts)
	return hosts
}

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

func isHostAlive(ip string) bool {
	timeout := time.Millisecond * 500

	// Try common development ports first
	devPorts := []int{3000, 8080, 8000, 5000, 4200, 8888}
	for _, port := range devPorts {
		conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", ip, port), timeout)
		if err == nil {
			conn.Close()
			return true
		}
	}

	// Try other common ports
	commonPorts := []int{22, 80, 443, 135, 139, 445}
	for _, port := range commonPorts {
		conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", ip, port), timeout)
		if err == nil {
			conn.Close()
			return true
		}
	}

	return false
}

func scanDevelopmentServers(hosts []string) []HostResult {
	var results []HostResult
	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, host := range hosts {
		wg.Add(1)
		go func(host string) {
			defer wg.Done()

			fmt.Printf("Scanning %s for development servers...\n", host)
			services := scanDevPorts(host)

			if len(services) > 0 {
				mu.Lock()
				results = append(results, HostResult{
					IP:       host,
					Hostname: getHostname(host),
					Services: services,
				})
				mu.Unlock()
			}
		}(host)
	}

	wg.Wait()
	return results
}

func scanDevPorts(host string) []ServiceInfo {
	// Extended list of development and common web server ports
	devPorts := []int{
		// Node.js/React/Next.js common ports
		3000, 3001, 3002, 3003, 3004, 3005,
		// Python (Django/Flask) common ports
		8000, 8001, 8002, 5000, 5001, 5002,
		// Java (Spring Boot) common ports
		8080, 8081, 8082, 8090, 9000, 9001,
		// .NET common ports
		5000, 5001, 7000, 7001, 44300, 44301,
		// Angular CLI
		4200, 4201, 4202,
		// Vue.js
		8080, 8081,
		// General development
		8888, 9090, 9999, 10000,
		// Jupyter/JupyterLab
		8888, 8889,
		// Webpack dev server
		8080, 3000,
		// PHP
		8000, 8080,
		// Ruby on Rails
		3000, 4000,
		// Go
		8080, 8000,
		// Other common dev ports
		6000, 6001, 7000, 7001, 7777, 8888,
	}

	var services []ServiceInfo
	var wg sync.WaitGroup
	var mu sync.Mutex

	semaphore := make(chan struct{}, 20)

	for _, port := range devPorts {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			semaphore <- struct{}{}
			defer func() { <-semaphore }()

			if service := identifyService(host, port); service.Port != 0 {
				mu.Lock()
				services = append(services, service)
				mu.Unlock()
			}
		}(port)
	}

	wg.Wait()

	// Sort by port number
	sort.Slice(services, func(i, j int) bool {
		return services[i].Port < services[j].Port
	})

	return services
}

func identifyService(host string, port int) ServiceInfo {
	// First check if port is open
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", host, port), time.Second)
	if err != nil {
		return ServiceInfo{}
	}
	conn.Close()

	service := ServiceInfo{
		Port: port,
		URL:  fmt.Sprintf("http://%s:%d", host, port),
	}

	// Try HTTP request to identify the service
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get(service.URL)
	if err != nil {
		// Try HTTPS
		service.URL = fmt.Sprintf("https://%s:%d", host, port)
		resp, err = client.Get(service.URL)
		if err != nil {
			// Port is open but not HTTP/HTTPS
			service.Service = "Unknown TCP Service"
			service.Status = "Open (Non-HTTP)"
			return service
		}
	}
	defer resp.Body.Close()

	service.Status = fmt.Sprintf("%d %s", resp.StatusCode, resp.Status)

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		service.Service = "HTTP Service"
		return service
	}

	bodyStr := string(body)

	// Extract page title
	service.Title = extractTitle(bodyStr)

	// Get server header
	service.Server = resp.Header.Get("Server")

	// Identify technology and framework
	service.Technology, service.Framework = identifyTechnology(bodyStr, resp.Header)

	// Determine service type
	service.Service = determineServiceType(service.Technology, service.Framework, service.Title, port)

	return service
}

func extractTitle(html string) string {
	titleRegex := regexp.MustCompile(`<title[^>]*>([^<]+)</title>`)
	matches := titleRegex.FindStringSubmatch(html)
	if len(matches) > 1 {
		return strings.TrimSpace(matches[1])
	}
	return ""
}

func identifyTechnology(body string, headers http.Header) (string, string) {
	bodyLower := strings.ToLower(body)

	// Check for Next.js
	if strings.Contains(bodyLower, "next.js") || strings.Contains(bodyLower, "_next/") ||
		strings.Contains(bodyLower, "__next") {
		return "Next.js", "React"
	}

	// Check for React
	if strings.Contains(bodyLower, "react") || strings.Contains(bodyLower, "reactdom") ||
		strings.Contains(bodyLower, "react-dom") {
		return "React", "React"
	}

	// Check for Vue.js
	if strings.Contains(bodyLower, "vue.js") || strings.Contains(bodyLower, "vue") ||
		strings.Contains(bodyLower, "__vue__") {
		return "Vue.js", "Vue"
	}

	// Check for Angular
	if strings.Contains(bodyLower, "angular") || strings.Contains(bodyLower, "ng-version") {
		return "Angular", "Angular"
	}

	// Check for Django
	if strings.Contains(bodyLower, "django") || strings.Contains(bodyLower, "csrfmiddlewaretoken") {
		return "Python", "Django"
	}

	// Check for Flask
	if strings.Contains(bodyLower, "flask") {
		return "Python", "Flask"
	}

	// Check for Spring Boot
	if strings.Contains(bodyLower, "spring") || strings.Contains(bodyLower, "spring-boot") {
		return "Java", "Spring Boot"
	}

	// Check for .NET
	if strings.Contains(bodyLower, "asp.net") || strings.Contains(bodyLower, ".net") ||
		headers.Get("Server") != "" && strings.Contains(strings.ToLower(headers.Get("Server")), "kestrel") {
		return ".NET", "ASP.NET"
	}

	// Check for Node.js indicators
	if strings.Contains(bodyLower, "express") || strings.Contains(bodyLower, "node.js") ||
		headers.Get("X-Powered-By") != "" && strings.Contains(strings.ToLower(headers.Get("X-Powered-By")), "express") {
		return "Node.js", "Express"
	}

	// Check for Jupyter
	if strings.Contains(bodyLower, "jupyter") || strings.Contains(bodyLower, "notebook") {
		return "Python", "Jupyter"
	}

	// Check for Webpack dev server
	if strings.Contains(bodyLower, "webpack") {
		return "JavaScript", "Webpack Dev Server"
	}

	// Check for PHP
	if strings.Contains(bodyLower, "php") {
		return "PHP", "PHP"
	}

	// Check for Ruby on Rails
	if strings.Contains(bodyLower, "rails") || strings.Contains(bodyLower, "ruby") {
		return "Ruby", "Rails"
	}

	// Check for Go
	if strings.Contains(bodyLower, "go") && strings.Contains(bodyLower, "server") {
		return "Go", "Go HTTP Server"
	}

	return "Unknown", ""
}

func determineServiceType(tech, framework, title string, port int) string {
	if framework != "" {
		return fmt.Sprintf("%s Development Server", framework)
	}

	if tech != "" {
		return fmt.Sprintf("%s Application", tech)
	}

	if title != "" {
		return fmt.Sprintf("Web Application (%s)", title)
	}

	// Port-based identification
	switch port {
	case 3000:
		return "Node.js/React Dev Server"
	case 8000:
		return "Python/Django Dev Server"
	case 5000:
		return "Flask/ASP.NET Dev Server"
	case 8080:
		return "Java/Spring Boot Server"
	case 4200:
		return "Angular Dev Server"
	case 8888:
		return "Jupyter Notebook"
	default:
		return "Development Server"
	}
}

func getHostname(ip string) string {
	names, err := net.LookupAddr(ip)
	if err != nil || len(names) == 0 {
		return ""
	}
	return strings.TrimSuffix(names[0], ".")
}

func displayResults(results []HostResult) {
	fmt.Println("\n" + strings.Repeat("=", 80))
	fmt.Println("DEVELOPMENT SERVERS FOUND")
	fmt.Println(strings.Repeat("=", 80))

	if len(results) == 0 {
		fmt.Println("No development servers found on the network.")
		fmt.Println("Make sure the servers are running and accessible.")
		return
	}

	for _, result := range results {
		fmt.Printf("\n🖥️  HOST: %s", result.IP)
		if result.Hostname != "" {
			fmt.Printf(" (%s)", result.Hostname)
		}
		fmt.Println()
		fmt.Println(strings.Repeat("-", 60))

		for _, service := range result.Services {
			fmt.Printf("  🌐 Port %d: %s\n", service.Port, service.Service)

			if service.Technology != "" {
				fmt.Printf("     Technology: %s", service.Technology)
				if service.Framework != "" && service.Framework != service.Technology {
					fmt.Printf(" (%s)", service.Framework)
				}
				fmt.Println()
			}

			if service.Title != "" {
				fmt.Printf("     Title: %s\n", service.Title)
			}

			if service.Server != "" {
				fmt.Printf("     Server: %s\n", service.Server)
			}

			fmt.Printf("     URL: %s\n", service.URL)
			fmt.Printf("     Status: %s\n", service.Status)
			fmt.Println()
		}
	}

	fmt.Printf("\n📊 Summary: Found %d hosts with %d development servers\n",
		len(results), getTotalServices(results))

	// Group by technology
	techCount := make(map[string]int)
	for _, result := range results {
		for _, service := range result.Services {
			if service.Technology != "" {
				techCount[service.Technology]++
			}
		}
	}

	if len(techCount) > 0 {
		fmt.Println("\n📈 Technology Distribution:")
		for tech, count := range techCount {
			fmt.Printf("  %s: %d server(s)\n", tech, count)
		}
	}
}

func getTotalServices(results []HostResult) int {
	total := 0
	for _, result := range results {
		total += len(result.Services)
	}
	return total
}

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}
