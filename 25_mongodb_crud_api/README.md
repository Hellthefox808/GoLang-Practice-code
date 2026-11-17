# 25 - Production MongoDB CRUD Microservice

**Author:** Ravi Ranjan Singh  
**Repository:** [GoLang-Practice-code](https://github.com/Hellthefox808/GoLang-Practice-code)  
**License:** MIT  

---

## 1. Overview
This module demonstrates enterprise MongoDB driver integration in Go (`go.mongodb.org/mongo-driver`). It adheres to 12-factor application standards, separating configuration from code and strictly avoiding credential hardcoding.

---

## 2. Critical Audits & Fixes Applied
1. **Zero Credential Hardcoding:** Database connection strings are loaded dynamically from `MONGO_URI`.
2. **Elimination of Process-Killing `log.Fatal`:** Controller handlers return HTTP error statuses (`500`, `400`, `404`) rather than terminating the Go runtime process when database queries fail.
3. **Correct Content-Type:** Headers declare `application/json`, replacing flawed form-encoded headers.
4. **Hex ObjectID Validation:** Incoming IDs are validated via `primitive.ObjectIDFromHex(id)`. Invalid formats return `400 Bad Request`.
5. **Cursor Lifecycle Safety:** Cursors are closed immediately via `defer cur.Close(ctx)` and verified with `cur.Err()`.

---

## 3. How to Run & Verify
```bash
# Optional: Set local or cloud Mongo URI
export MONGO_URI="mongodb://localhost:27017"

go run main.go
go test -v ./...
```
