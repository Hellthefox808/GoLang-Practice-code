# 20 - RFC 3986 URL Parsing & Safe Query Manipulation

**Author:** Ravi Ranjan Singh  
**Repository:** [GoLang-Practice-code](https://github.com/Hellthefox808/GoLang-Practice-code)  
**License:** MIT  

---

## 1. Overview
URL construction via string concatenation (`"http://" + host + "/" + path + "?" + params`) is fragile and introduces injection vulnerabilities. The standard `net/url` package provides RFC 3986 compliant encoding, percent-escaping, and component resolution.

---

## 2. Common Pitfall: `RawPath` vs `RawQuery`
Beginner tutorials frequently confuse `RawPath` with `RawQuery`:
- **`Path` vs `RawPath`:** `Path` is the unescaped path string (e.g. `/my files/`). `RawPath` is populated only when the path contains escaped characters like `%2F` that would otherwise alter the path structure when unescaped.
- **`RawQuery`:** The raw, encoded query parameter string (e.g. `foo=bar&num=42`).
- **`Query()` Method:** Decodes the query string into a `url.Values` (which is `map[string][]string`). `url.Values.Encode()` serializes parameters in deterministic sorted key order.

---

## 3. How to Run & Verify
```bash
go run main.go
go test -v ./...
```
