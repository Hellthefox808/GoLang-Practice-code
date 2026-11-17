# 03 - Robust User & Stream Input Handling

**Author:** Ravi Ranjan Singh  
**Repository:** [GoLang-Practice-code](https://github.com/Hellthefox808/GoLang-Practice-code)  
**License:** MIT  

---

## 1. Overview
Input handling in CLI utilities, HTTP servers, or socket connections requires meticulous handling of delimiters, buffer boundaries, character encodings, and platform-specific line breaks (`\r\n` on Windows vs `\n` on Linux/macOS).

---

## 2. Key Techniques & Design
1. **`bufio.Reader` vs `bufio.Scanner`:**
   - `bufio.Reader`: Ideal for unbounded streams, fine-grained byte/rune manipulation, or reading up to specific byte delimiters.
   - `bufio.Scanner`: Higher level, token-based interface with built-in tokenizers (`ScanLines`, `ScanWords`, `ScanRunes`). Note: `bufio.Scanner` has a default max buffer limit of 64KB unless expanded with `Buffer()`.
2. **Cross-Platform CRLF Normalization:**
   ```go
   cleaned := strings.TrimRight(raw, "\r\n")
   ```
3. **Decoupled Testability:**
   By designing functions to accept `io.Reader` instead of binding strictly to `os.Stdin`, functions can be unit-tested deterministically using `strings.NewReader` or `bytes.Buffer`.

---

## 3. How to Run & Verify
```bash
go run main.go
go test -v ./...
```
