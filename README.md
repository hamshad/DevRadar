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

[![Build and Release](https://github.com/hamshad/DevRadar/actions/workflows/build.yml/badge.svg)](https://github.com/hamshad/DevRadar/actions/workflows/build.yml)
[![Go Version](https://img.shields.io/badge/Go-1.21+-blue.svg)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![Platform](https://img.shields.io/badge/Platform-Linux%20%7C%20macOS%20%7C%20Windows-lightgrey.svg)](#-installation--usage)

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

### 🎯 Quick Install (Recommended)

**Option 1: Download Pre-built Binaries** *(Easiest)*

Visit the [**Releases Page**](https://github.com/hamshad/DevRadar/releases/latest) and download the appropriate executable for your system:

| Platform | Architecture | Download |
|----------|-------------|----------|
| **Linux** | AMD64 (x86_64) | `devradar-linux-amd64` |
| **Linux** | ARM64 | `devradar-linux-arm64` |
| **macOS** | Intel (x86_64) | `devradar-macos-amd64` |
| **macOS** | Apple Silicon (M1/M2) | `devradar-macos-arm64` |
| **Windows** | AMD64 (x86_64) | `devradar-windows-amd64.exe` |
| **Windows** | ARM64 | `devradar-windows-arm64.exe` |

**Linux/macOS:**
```bash
# Download the appropriate binary
wget https://github.com/hamshad/DevRadar/releases/latest/download/devradar-linux-amd64

# Make it executable
chmod +x devradar-linux-amd64

# Run the scanner
./devradar-linux-amd64
```

**Windows:**
```cmd
# Download devradar-windows-amd64.exe from releases
# Run directly from command prompt or PowerShell
devradar-windows-amd64.exe
```

### 🔧 Build from Source

**Prerequisites:**
- **Go 1.21+** installed on your system
- Git (optional, for cloning)

**Steps:**
```bash
# Clone the repository
git clone https://github.com/hamshad/DevRadar.git
cd DevRadar

# Build for your platform
go build -ldflags="-s -w" -o devradar dev_radar.go

# Run the scanner
./devradar
```

**Cross-compile for other platforms:**
```bash
# Linux AMD64
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o devradar-linux dev_radar.go

# macOS ARM64 (Apple Silicon)
GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o devradar-macos dev_radar.go

# Windows AMD64
GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o devradar.exe dev_radar.go
```

### 🚀 Running DevRadar

Once you have the executable:

1. **Simply run the binary:**
   ```bash
   ./devradar-linux-amd64        # Linux
   ./devradar-macos-arm64        # macOS
   devradar-windows-amd64.exe    # Windows
   ```

2. **What happens automatically:**
   - Detects your local network range
   - Discovers active hosts
   - Scans for development services
   - Displays results with beautiful formatting

### 🔐 Verification

All releases include checksums for security verification:

```bash
# Download checksums.txt from the release
wget https://github.com/hamshad/DevRadar/releases/latest/download/checksums.txt

# Verify integrity (Linux/macOS)
sha256sum -c checksums.txt

# Verify integrity (Windows PowerShell)
Get-FileHash devradar-windows-amd64.exe -Algorithm SHA256
```

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

## 🔄 Continuous Integration

DevRadar uses **GitHub Actions** for automated building and releasing:

### 🏗️ Automated Builds
- **Triggers**: Every push to `main`, PR creation, and tag creation
- **Platforms**: Builds for 6 platform/architecture combinations
- **Artifacts**: Stores build artifacts for 90 days
- **Checksums**: Automatically generates SHA256 checksums for verification

### 📋 Build Matrix
| Platform | Architecture | Binary Name |
|----------|-------------|-------------|
| Linux | AMD64 | `devradar-linux-amd64` |
| Linux | ARM64 | `devradar-linux-arm64` |
| macOS | AMD64 | `devradar-macos-amd64` |
| macOS | ARM64 | `devradar-macos-arm64` |
| Windows | AMD64 | `devradar-windows-amd64.exe` |
| Windows | ARM64 | `devradar-windows-arm64.exe` |

### 🚀 Release Process
1. **Tag Creation**: Create a version tag (e.g., `v1.0.0`)
2. **Automatic Building**: GitHub Actions builds all platform binaries
3. **Release Creation**: Automatically creates a GitHub release with:
   - All platform binaries
   - SHA256 checksums
   - Release notes
   - Download instructions

---

## 🛡️ Security Considerations

- **Local Network Only**: Designed for local network scanning, not internet-wide discovery
- **Read-Only Operations**: Only performs GET requests, no data modification
- **No Authentication**: Does not attempt to bypass authentication mechanisms
- **Respectful Scanning**: Built-in rate limiting to avoid overwhelming target services
- **Binary Verification**: All releases include SHA256 checksums for integrity verification

---

## 🤝 Contributing

DevRadar welcomes contributions! Here's how you can help:

### 🐛 Reporting Issues
- Use GitHub Issues to report bugs or request features
- Include your OS, architecture, and DevRadar version
- Provide steps to reproduce any issues

### 💻 Development
1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Test across platforms (CI will help with this)
5. Submit a pull request

### 🔧 Areas for Contribution
- **New Technology Detection**: Add support for more frameworks
- **Performance Improvements**: Optimize scanning algorithms
- **UI Enhancements**: Improve terminal output formatting
- **Platform Support**: Add support for additional architectures
- **Documentation**: Improve README, add examples, write tutorials

---

## 📁 License

MIT License - Feel free to use, modify, and distribute!

---

## 📊 Project Stats

- **Language**: Go 1.21+
- **Platforms**: Linux, macOS, Windows
- **Architectures**: AMD64, ARM64
- **CI/CD**: GitHub Actions
- **License**: MIT
- **Maintainer**: [@hamshad](https://github.com/hamshad)

---

**Made with ❤️ for developers who want to discover the hidden gems running on their network!**

---

## 🏷️ Tags

`go` `golang` `network-scanner` `development-tools` `devops` `port-scanner` `technology-detection` `local-network` `dev-server` `stack-detection` `cross-platform` `github-actions` `automated-builds`
