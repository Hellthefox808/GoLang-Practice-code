# Mock Server (Go & Node.js)

**Author:** Ravi Ranjan Singh  
**Repository:** [GoLang-Practice-code](https://github.com/Hellthefox808/GoLang-Practice-code)  
**License:** MIT  

---

## 1. Overview
This directory provides a local test server capable of serving `GET`, `POST` (JSON), and `POST` (Form URL-encoded) routes for testing Go HTTP clients (such as modules 19 and 21) without external cloud dependencies.

---

## 2. Running Options

### Option A: Pure Go (Zero External Dependencies)
```bash
go run main.go
```
Listens on `http://localhost:8000`.

### Option B: Node.js Express
```bash
cd node_server
npm install
node index.js
```
