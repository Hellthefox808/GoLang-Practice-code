# 23 - Go Modules, Standard Router (Go 1.22+), and Graceful Shutdown

**Author:** Ravi Ranjan Singh  
**Repository:** [GoLang-Practice-code](https://github.com/Hellthefox808/GoLang-Practice-code)  
**License:** MIT  

---

## 1. Overview
In modern Go (1.22+), `net/http.ServeMux` received routing enhancements: method matching (`"GET /path"`) and wildcard matching (`"/items/{id}"`) are now first-class built-ins, reducing dependency on external routers like Gorilla Mux.

---

## 2. Production Graceful Shutdown
In containerized Kubernetes environments, pods receive `SIGTERM` when scaling down or deploying updates. If an HTTP server immediately exits, active client requests are terminated abruptly with 502/504 errors.
The production pattern:
1. Intercept `os.Interrupt` and `syscall.SIGTERM` with `signal.Notify`.
2. Stop accepting new connections while active connections drain using `server.Shutdown(ctx)`.
3. Allow a grace period (e.g. 5-30s) before terminating.

---

## 3. How to Run & Verify
```bash
go run main.go
go test -v ./...
```
