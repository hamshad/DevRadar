
  _____             _____           _            
 |  __ \           |  __ \         | |           
 | |  | | _____   _| |__) |__ _  __| | __ _ _ __ 
 | |  | |/ _ \ \ / /  _  // _` |/ _` |/ _` | '__|
 | |__| |  __/\ V /| | \ \ (_| | (_| | (_| | |   
 |_____/ \___| \_/ |_|  \_\__,_|\__,_|\__,_|_|   
                                                 
                                                 

         Development Server & Stack Scanner


# 🧪 DevRadar

**DevRadar** is a powerful **Development-Focused Network Scanner** to help teams discover and identify development servers and hosted applications across a local network. This enhanced version is tailored specifically for modern dev environments.

---

## 🚀 Key Features
### 🎯 Development-Focused Scanning
- Scans common development ports:
  - `3000`, `8000`, `5000`, `8080`, `4200`, and others.
- Identifies popular **technology stacks**:
  - **Frontend**: Next.js, React, Angular, Vue.js
  - **Backend**: Express, Django, Flask, Spring Boot, .NET, Rails, Go
- Framework & Language Detection:
  - Node.js, Python, Java, C#, Ruby, PHP, Go

### 🔍 Smart Detection
- **HTTP/HTTPS Probing**:
  - Sends requests to detect running services
- **Content Analysis**:
  - Analyzes HTML, response headers, and metadata
- **Title Extraction**:
  - Extracts page titles to help recognize applications
- **Server Header Parsing**:
  - Identifies server software from headers

### 📊 Detailed Reporting
- Lists all discovered hosts and applications
- Groups services by technology type
- Displays:
  - Port
  - Technology detected
  - App title
  - Access URL
  - HTTP status code
  - Server header info

---

## 🧠 Technologies It Can Identify

| Language      | Frameworks/Technologies             |
|---------------|--------------------------------------|
| JavaScript    | Next.js, React, Vue.js, Angular, Express |
| Python        | Django, Flask, Jupyter Notebook     |
| Java          | Spring Boot                         |
| .NET          | ASP.NET                             |
| PHP           | Generic PHP Applications            |
| Ruby          | Ruby on Rails                       |
| Go            | Go HTTP Servers                     |

---

## 🖥️ Example Output

```
HOST: 192.168.1.105 (johns-laptop)
------------------------------------------------------------
🌐 Port 3000: Next.js Development Server  
   Technology: Next.js (React)  
   Title: My Awesome App  
   URL: http://192.168.1.105:3000  
   Status: 200 OK  

🌐 Port 8000: Python Application  
   Technology: Python (Django)  
   Title: Django Admin  
   URL: http://192.168.1.105:8000  
   Status: 200 OK  
```

---

## 📦 How to Use

1. **Run the scanner**:
   ```bash
   go run dev_radar.go
   ```
2. **Wait for scanning** to complete.
3. **View the results** and discover:
   - Who’s running which app
   - URLs to access them
   - Technology behind each one

---

## 🔍 Ideal Use Cases

- Internal developer teams working on shared networks
- DevOps looking to audit running dev services
- Quickly identify running apps during hackathons or collaborative sessions

---

## 🛠️ Requirements

- Go 1.18 or later
- Network access to local subnet

---

## 📁 License

MIT License

---

Made with ❤️ to help dev teams discover hidden gems on their network!
