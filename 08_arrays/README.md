# 08 - Fixed-Size Arrays & Value Semantics

**Author:** Ravi Ranjan Singh  
**Repository:** [GoLang-Practice-code](https://github.com/Hellthefox808/GoLang-Practice-code)  
**License:** MIT  

---

## 1. Overview
In Go, an array is a contiguous block of memory with a fixed size determined at compile time.

### Fundamental Rules:
1. **Length is in the Type:** `[4]int` and `[5]int` are distinct, incompatible types.
2. **Value Semantics:** Arrays are values. Assigning `a = b` or passing an array to a function copies every element. In contrast to C (where arrays decay into pointers), Go arrays do not decay.
3. **Slices vs Arrays:** Slices (`[]T`) are built on top of arrays and are used for 95%+ of dynamic collection use cases in Go. Arrays are reserved for low-level memory buffers, cryptographic blocks (e.g. `[32]byte` for SHA-256), or cache-locality critical structures.

---

## 2. How to Run & Verify
```bash
go run main.go
go test -v ./...
```
