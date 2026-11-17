# 10 - Hash Tables, Map Semantics, and Safety

**Author:** Ravi Ranjan Singh  
**Repository:** [GoLang-Practice-code](https://github.com/Hellthefox808/GoLang-Practice-code)  
**License:** MIT  

---

## 1. Overview
Go maps are built upon dynamically growing hash tables (`runtime.hmap`), consisting of buckets each holding up to 8 key-value pairs.

---

## 2. Critical Safety Laws
1. **The Nil Map Rule:**
   - Reading `nilMap[key]` is completely safe and evaluates to the zero value of the value type.
   - Writing `nilMap[key] = val` causes an immediate, unrecoverable runtime panic.
2. **The Comma-Ok Idiom:**
   Because non-existent keys return zero values, `val, ok := m[key]` is mandatory when a zero-value could be a valid value.
3. **Iteration Non-Determinism:**
   Go intentionally randomizes map iteration order starting at a random bucket offset to prevent applications from relying on specific iteration sequences.
4. **Concurrency Hazard:**
   Standard maps are NOT safe for concurrent reads and writes. A concurrent write without synchronization triggers a fatal runtime error (`fatal error: concurrent map writes`). Use `sync.RWMutex` or `sync.Map` for concurrent access.

---

## 3. How to Run & Verify
```bash
go run main.go
go test -v ./...
```
