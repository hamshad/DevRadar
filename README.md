```
  _____             _____           _            
 |  __ \           |  __ \         | |           
 | |  | | _____   _| |__) |__ _  __| | __ _ _ __ 
 | |  | |/ _ \ \ / /  _  // _` |/ _` |/ _` | '__|
 | |__| |  __/\ V /| | \ \ (_| | (_| | (_| | |   
 |_____/ \___| \_/ |_|  \_\__,_|\__,_|\__,_|_|   

       Development Server & Stack Scanner
```


# 🧪 DevRadar

**DevRadar** is a powerful **Development-Focused Network Scanner** built in Go to help teams discover and identify development servers and hosted applications across a local network. Perfect for DevOps teams, developers working in shared environments, hackathons, and security audits.

---

## 🚀 Key Features

### 🎯 Comprehensive Development Server Detection
- **Auto-Discovery**: Automatically detects local network range and discovers active hosts
- **Concurrent Scanning**: Multi-threaded scanning for fast network discovery
- **Smart Port Detection**: Scans 40+ common development ports including:
  - **Node.js/React**: 3000-3005
  - **Python/Django/Flask**: 8000-8002, 5000-5002
  - **Java/Spring Boot**: 8080-8082, 8090, 9000-9001
  - **Angular**: 4200-4202
  - **.NET**: 5000-5001, 7000-7001, 44300-44301
  - **Jupyter**: 8888-8889
  - **General Dev**: 6000-6001, 7777, 9090, 9999, 10000

### 🔍 Advanced Technology Detection
- **HTTP/HTTPS Probing**: Automatically tries both HTTP and HTTPS protocols
- **Content Analysis**: Deep inspection of HTML content and response headers
- **Framework Detection**: Identifies popular frameworks and technologies:
  - **Frontend**: Next.js, React, Vue.js, Angular
  - **Backend**: Express, Django, Flask, Spring Boot, ASP.NET, Rails, Go
  - **Development Tools**: Jupyter Notebook, Webpack Dev Server
- **Title Extraction**: Extracts page titles using regex pattern matching
- **Server Identification**: Analyzes server headers for additional insights

### 📊 Beautiful Console Output
- **Styled Terminal Interface**: Rich ANSI color-coded output with Unicode characters
- **Organized Display**: Clear hierarchical presentation of results
- **Real-time Progress**: Live scanning progress updates
- **Detailed Reporting**: Comprehensive information for each discovered service
- **Technology Statistics**: Summary with technology distribution and visual charts

---

## 🧠 Technologies It Can Identify

| Technology | Detection Method | Frameworks/Tools |
|------------|------------------|------------------|
| **JavaScript** | Content & Headers | Next.js, React, Vue.js, Angular, Express, Webpack |
| **Python** | Content Analysis | Django, Flask, Jupyter Notebook |
| **Java** | Content & Patterns | Spring Boot |
| **.NET** | Headers & Content | ASP.NET, Kestrel Server |
| **Node.js** | Headers & Content | Express.js |
| **PHP** | Content Detection | Generic PHP Applications |
| **Ruby** | Content Analysis | Ruby on Rails |
| **Go** | Content Patterns | Go HTTP Servers |

---

## 🖥️ Example Output

```
  _____             _____           _            
 |  __ \           |  __ \         | |           
 | |  | | _____   _| |__) |__ _  __| | __ _ _ __ 
 | |  | |/ _ \ \ / /  _  // _` |/ _` |/ _` | '__|
 | |__| |  __/\ V /| | \ \ (_| | (_| | (_| | |   
 |_____/ \___| \_/ |_|  \_\__,_|\__,_|\__,_|_|   
                                                 
         Development Server & Stack Scanner

═══════════════════════════════════════════════════════════════════════════════

┌─ Network Information ────────────────────────────────────────────────────────┐
│ Local IP:            192.168.1.100
│ Scanning network:    192.168.1.0/24

┌─ Host Discovery ─────────────────────────────────────────────────────────────┐
│ 🔍 Discovering hosts...
│ ✅ Found 5 active hosts

┌─ Development Server Scanning ────────────────────────────────────────────────┐
│ 🔍 Scanning 192.168.1.105 for development servers...

╔══════════════════════════════════════════════════════════════════════════════╗
║                     🎯 DEVELOPMENT SERVERS DISCOVERED                        ║
╚══════════════════════════════════════════════════════════════════════════════╝

┌─ 🖥️  HOST: 192.168.1.105 (johns-laptop) ──────────────────────────────────┐
│  🌐 Port 3000: Next.js Development Server
│     🔧 Technology: Next.js (React)
│     📄 Title: My Awesome App
│     🔗 URL: http://192.168.1.105:3000
│     📊 Status: 200 OK
│
│  🌐 Port 8000: Python Application
│     🔧 Technology: Python (Django)
│     📄 Title: Django Administration
│     ⚙️  Server: WSGIServer/0.2
│     🔗 URL: http://192.168.1.105:8000
│     📊 Status: 200 OK
└──────────────────────────────────────────────────────────────────────────────┘

╔══════════════════════════════════════════════════════════════════════════════╗
║                              📊 SCAN SUMMARY                                 ║
╚══════════════════════════════════════════════════════════════════════════════╝

🏠 Total Hosts: 1
🌐 Total Services: 2

📈 Technology Distribution:
  Next.js      ██ (1)
  Python       ██ (1)

🎉 Scan completed successfully! Happy developing! 🚀
```

---

## 📦 Installation & Usage

### Prerequisites
- **Go 1.18+** installed on your system
- Network access to local subnet
- Terminal/Command prompt access

### Quick Start
1. **Clone or download** the repository:
   ```bash
   git clone https://github.com/hamshad/DevRadar.git
   cd DevRadar
   ```

2. **Run the scanner**:
   ```bash
   go run dev_radar.go
   ```

3. **Alternative - Build and run**:
   ```bash
   go build -o devradar dev_radar.go
   ./devradar
   ```

### What Happens During Scan
1. **Network Detection**: Automatically detects your local IP and network range
2. **Host Discovery**: Pings common development and system ports to find active hosts
3. **Service Scanning**: Tests 40+ development ports on each discovered host
4. **Technology Identification**: Analyzes HTTP responses to identify frameworks
5. **Report Generation**: Displays organized results with technology statistics

---

## 🔍 Ideal Use Cases

- **Team Development**: Discover what teammates are running on shared networks
- **DevOps Audits**: Quickly inventory development services across infrastructure
- **Hackathons**: Find and catalog running applications during collaborative events
- **Security Assessment**: Identify exposed development servers for security review
- **Network Documentation**: Generate reports of active development services
- **Troubleshooting**: Locate services that might be conflicting or causing issues

---

## ⚡ Technical Details

### Architecture
- **Concurrent Design**: Uses Go routines with semaphores for controlled concurrency
- **Network Discovery**: CIDR-based subnet scanning with automatic network detection
- **Service Detection**: TCP connection testing followed by HTTP/HTTPS probing
- **Content Analysis**: Regex-based pattern matching for technology identification
- **Error Handling**: Graceful handling of network timeouts and connection failures

### Performance Optimizations
- **Parallel Processing**: Utilizes all CPU cores with `GOMAXPROCS`
- **Connection Pooling**: Efficient HTTP client with proper timeout handling
- **Selective Scanning**: Prioritizes common development ports for faster discovery
- **Memory Efficient**: Streaming response processing to minimize memory usage

### Network Behavior
- **Non-Invasive**: Only performs lightweight TCP connections and HTTP GET requests
- **Timeout Management**: 500ms for port checks, 5s for HTTP requests
- **Rate Limiting**: Semaphore-controlled concurrency to prevent network flooding
- **Protocol Detection**: Automatic fallback from HTTPS to HTTP

---

## 🛡️ Security Considerations

- **Local Network Only**: Designed for local network scanning, not internet-wide discovery
- **Read-Only Operations**: Only performs GET requests, no data modification
- **No Authentication**: Does not attempt to bypass authentication mechanisms
- **Respectful Scanning**: Built-in rate limiting to avoid overwhelming target services

---

## 🤝 Contributing

DevRadar is open for contributions! Whether it's adding new technology detection patterns, improving the UI, or optimizing performance, all contributions are welcome.

---

## 📁 License

MIT License - Feel free to use, modify, and distribute!

---

**Made with ❤️ for developers who want to discover the hidden gems running on their network!**

---

## 🏷️ Tags

`go` `golang` `network-scanner` `development-tools` `devops` `port-scanner` `technology-detection` `local-network` `dev-server` `stack-detection`
