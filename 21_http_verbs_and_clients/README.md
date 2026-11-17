# 21 - HTTP Verbs, Headers, and Content-Type Handling

**Author:** Ravi Ranjan Singh  
**Repository:** [GoLang-Practice-code](https://github.com/Hellthefox808/GoLang-Practice-code)  
**License:** MIT  

---

## 1. Overview
Modern REST and RPC APIs interact via distinct HTTP verbs (`GET`, `POST`, `PUT`, `PATCH`, `DELETE`). Go's `net/http` package supports arbitrary methods and payloads via `http.NewRequestWithContext`.

---

## 2. Payload Encodings
1. **JSON (`application/json`):**
   - Serialized via `json.Marshal(struct)`.
   - Streamed via `bytes.NewBuffer(jsonBytes)` or `bytes.NewReader()`.
2. **URL Encoded Forms (`application/x-www-form-urlencoded`):**
   - Constructed using `url.Values`.
   - Serialized using `url.Values.Encode()`.
   - Streamed via `strings.NewReader(formData.Encode())`.
3. **In-Memory Testing:**
   Never hardcode local ports (`localhost:8000`) for unit and integration testing. Always rely on `net/http/httptest.NewServer` to allocate ephemeral ports dynamically, allowing test execution in parallel across CI/CD runners.

---

## 3. How to Run & Verify
```bash
go run main.go
go test -v ./...
```
