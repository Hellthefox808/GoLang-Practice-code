# 07 - Pointers, Memory Model, and Mutation Semantics

**Author:** Ravi Ranjan Singh  
**Repository:** [GoLang-Practice-code](https://github.com/Hellthefox808/GoLang-Practice-code)  
**License:** MIT  

---

## 1. Overview
Go features pointers but intentionally omits pointer arithmetic (unlike C/C++). This provides direct control over memory layout and pass-by-reference behavior without the buffer-overrun and arbitrary memory offset vulnerabilities inherent in pointer arithmetic.

---

## 2. Core Tenets
1. **Everything is Pass-by-Value:**
   In Go, arguments are always copied into function parameters. When a pointer `*T` is passed, the memory address itself is copied by value.
2. **Operators:**
   - `&`: Address-of operator (returns `*T`).
   - `*`: Dereference operator (accesses underlying value `T`).
3. **Escape Analysis:**
   The Go compiler analyzes variable lifecycles. If a variable created inside a function has its pointer returned or captured in a closure, it "escapes to the heap" to ensure safety. Otherwise, it is allocated on the fast stack and discarded when the stack frame returns.

---

## 3. How to Run & Verify
```bash
go run main.go
go test -v ./...
# View escape analysis diagnostics
go build -gcflags="-m" main.go
```
