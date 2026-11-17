# 24 - Enterprise REST API with Gorilla Mux

**Author:** Ravi Ranjan Singh  
**Repository:** [GoLang-Practice-code](https://github.com/Hellthefox808/GoLang-Practice-code)  
**License:** MIT  

---

## 1. Overview & Clean Architecture
This module demonstrates an enterprise-grade RESTful API built on Gorilla Mux, structured following layered architectural separation:

```
[ Client HTTP Request ]
          |
          v
[ Gorilla Mux Router ]
          |
          v
[ HTTP Controller Handlers ]  <--->  [ Strict JSON Decoder / Response Envelopes ]
          |
          v
[ Mutex-Guarded Repository ]  <--->  [ sync.RWMutex Concurrency Protection ]
          |
          v
[ In-Memory Datastore ]
```

---

## 2. Eliminating Critical Concurrency Race Conditions
Introductory tutorials frequently store collections in global unprotected slices (`var courses []Course`). In Go, `http.ListenAndServe` spawns an independent Goroutine for each incoming HTTP connection.

### The Race Condition Disaster:
If two clients submit concurrent POST requests or one client writes while another reads, the slice triggers a simultaneous memory race condition, crashing the server process with:
```
fatal error: concurrent map read and map write / slice bounds out of range
```

### The Solution: `sync.RWMutex`
Our repository synchronizes all reads and writes:
```go
type CourseRepository struct {
    mu      sync.RWMutex
    courses map[string]Course
}

func (r *CourseRepository) GetAll() []Course {
    r.mu.RLock()
    defer r.mu.RUnlock()
    // Multiple readers execute concurrently
}

func (r *CourseRepository) Create(c Course) (Course, error) {
    r.mu.Lock()
    defer r.mu.Unlock()
    // Exclusive write lock guarantees atomic append
}
```

---

## 3. Standard Unified JSON Envelope & HTTP Status Codes
All responses conform to an enterprise contract:
```json
{
    "success": true,
    "message": "Course registered successfully",
    "data": {
        "id": "103",
        "title": "Go Concurrency Mastery",
        "price_usd": 99.0
    },
    "error": ""
}
```

### Semantic Status Code Mapping:
- `200 OK`: Successful reads (`GET`), updates (`PUT`).
- `201 Created`: Successful creation (`POST`).
- `204 No Content`: Successful deletion (`DELETE`).
- `400 Bad Request`: Malformed JSON or nil request body.
- `404 Not Found`: Target entity ID absent from datastore.
- `409 Conflict`: Business validation violation (e.g. duplicate course title).
- `422 Unprocessable Entity`: Missing required domain fields.

---

## 4. Verification & Testing
```bash
go run main.go
go test -v ./...
# Verify zero race conditions
go test -race ./...
```
