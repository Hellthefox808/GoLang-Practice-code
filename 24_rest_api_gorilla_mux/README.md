# 24 - Enterprise REST API with Gorilla Mux

**Author:** Ravi Ranjan Singh  
**Repository:** [GoLang-Practice-code](https://github.com/Hellthefox808/GoLang-Practice-code)  
**License:** MIT  

---

## 1. Overview
This module demonstrates production RESTful web services in Go. It completely resolves the race conditions, status code omissions, and placeholder defects common in introductory tutorial code.

---

## 2. Architectural Enhancements
1. **Thread-Safe Repository (`sync.RWMutex`):**
   HTTP handlers execute concurrently across distinct Goroutines. Shared state must be synchronized to prevent memory corruption and race condition crashes.
2. **Unified Response Envelope:**
   ```json
   {
       "success": true,
       "message": "Course registered successfully",
       "data": { ... },
       "error": ""
   }
   ```
3. **Semantic Status Codes:**
   - `200 OK` for successful queries and updates.
   - `201 Created` on resource creation.
   - `204 No Content` on successful deletions.
   - `400 Bad Request` for malformed payloads.
   - `404 Not Found` for nonexistent resource IDs.
   - `409 Conflict` for duplicate unique constraints.
   - `422 Unprocessable Entity` for domain validation failures.

---

## 3. How to Run & Verify
```bash
go run main.go
go test -v ./...
```
