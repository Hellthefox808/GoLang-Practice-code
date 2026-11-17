# 19 - Production HTTP Clients & Connection Pool Reuse

**Author:** Ravi Ranjan Singh  
**Repository:** [GoLang-Practice-code](https://github.com/Hellthefox808/GoLang-Practice-code)  
**License:** MIT  

---

## 1. Overview
In Go, `http.DefaultClient` has **no timeout**. Using `http.Get()` or `http.DefaultClient` in production services is a major anti-pattern because hung connections can exhaust OS file descriptors and thread pools.

---

## 2. The Golden Rules of Go HTTP Clients
1. **Always Configure Timeouts:** Set client-level timeouts and attach `context.WithTimeout` to each `http.Request`.
2. **Drain and Close the Body:**
   ```go
   defer func() {
       io.Copy(io.Discard, resp.Body)
       resp.Body.Close()
   }()
   ```
   If the body is closed without reading all bytes to EOF, the underlying TCP connection is reset and terminated instead of being recycled in the keep-alive connection pool.
3. **Deterministic Unit Testing:** Use `net/http/httptest.NewServer` to test HTTP clients offline without flakiness or external internet dependency.

---

## 3. How to Run & Verify
```bash
go run main.go
go test -v ./...
```
