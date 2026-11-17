# 14 - Loops, Range Iterators, and Labeled Jumps

**Author:** Ravi Ranjan Singh  
**Repository:** [GoLang-Practice-code](https://github.com/Hellthefox808/GoLang-Practice-code)  
**License:** MIT  

---

## 1. Overview
Go deliberately provides only one looping keyword: `for`. It replaces `while`, `do-while`, and enhanced `for-each` through varied syntactical forms.

---

## 2. Iteration Forms
1. **Classic 3-Component:** `for i := 0; i < N; i++`
2. **Condition ("While"):** `for condition`
3. **Infinite Loop:** `for { ... }`
4. **Range Loop:** `for key, val := range collection`
   - Slices/Arrays: `index, element`
   - Maps: `key, value`
   - Strings: `byteOffset, rune`
   - Channels: `element` (until closed)

---

## 3. Go 1.22+ Per-Iteration Scoping Revolution
Historically in Go (prior to 1.22), loop variables were allocated once per loop and updated in place. Capturing a loop variable inside a closure or goroutine without an explicit local copy (`v := v`) led to subtle race conditions where all closures read the final iteration value.
Starting in **Go 1.22**, each iteration introduces a brand new variable binding, making closures inside loops completely safe and intuitive.

---

## 4. How to Run & Verify
```bash
go run main.go
go test -v ./...
```
