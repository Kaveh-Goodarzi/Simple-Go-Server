# Simple Go Server

A minimal HTTP server written in Go for serving static files and basic API endpoints.
Built using only Go’s standard library.

---

## ✨ Features

- Serve static HTML and CSS files
- Health check endpoint
- Echo API endpoint
- No frameworks, no dependencies
- Clean and beginner-friendly code
- Unit tests and CI with GitHub Actions

---

## 📦 Requirements

- Go 1.21+

Check your Go version:

```bash
go version
```

---

## 🚀 Running the Server

Clone the repository:

```bash
git clone https://github.com/Kaveh-Goodarzi/Simple-Go-Server.git
cd Simple-Go-Server
```

Run the server:

```bash
go run main.go
```

Open in browser:

```bash
http://localhost:8080
```

---

## 🧪 Running Tests

```bash
go test -v
```

---

## 🗂 Project Structure

```
    Simple-Go-Server/
    ├── .github/workflows/go.yml
    ├── main.go
    ├── main_test.go
    ├── static/
    │   ├── index.html
    │   ├── about.html
    │   └── css/style.css
    ├── go.mod
    └── README.md
```