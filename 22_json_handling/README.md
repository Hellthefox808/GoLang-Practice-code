# 22 - Advanced JSON Serialization, Tags, and Strict Decoding

**Author:** Ravi Ranjan Singh  
**Repository:** [GoLang-Practice-code](https://github.com/Hellthefox808/GoLang-Practice-code)  
**License:** MIT  

---

## 1. Overview
JSON serialization is at the heart of web backends and microservices. Go's `encoding/json` package uses struct reflection and struct tags to map Go types to JSON keys.

---

## 2. Essential Struct Tags
- `json:"name"`: Custom JSON key name.
- `json:",omitempty"`: Omits the field from JSON if it holds its zero value (e.g. empty string, 0, false, nil slice/pointer).
- `json:"-"`: Permanently prevents serialization or deserialization (critical for passwords, tokens, API keys).
- `json:",string"`: Encodes numbers or booleans inside quotes as strings.

---

## 3. Production API Security
When ingesting client payloads in REST endpoints, always use `json.NewDecoder(r.Body).DisallowUnknownFields()`. By default, `json.Unmarshal` silently ignores extraneous JSON fields, which can mask client typing mistakes or facilitate mass-assignment vulnerabilities.

---

## 4. How to Run & Verify
```bash
go run main.go
go test -v ./...
```
