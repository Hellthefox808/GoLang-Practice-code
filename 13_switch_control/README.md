# 13 - Switch Statements, Tagless Ladders, and Type Switches

**Author:** Ravi Ranjan Singh  
**Repository:** [GoLang-Practice-code](https://github.com/Hellthefox808/GoLang-Practice-code)  
**License:** MIT  

---

## 1. Overview
Go's `switch` statement is significantly more flexible and safer than its C, C++, or Java counterparts:
- **Automatic Break:** In Go, cases break automatically upon completion. No explicit `break` keyword is required.
- **Multiple Values:** Multiple match candidates can be grouped on a single line separated by commas (`case 1, 2, 3:`).
- **Tagless Switches:** Omitting the expression (`switch { case ...: }`) evaluates cases as boolean predicates, serving as an idiomatic replacement for long `if-else if-else` chains.
- **Type Switches:** Inspect the underlying dynamic type of an `interface{}` / `any` variable at runtime using `switch v := i.(type)`.

---

## 2. How to Run & Verify
```bash
go run main.go
go test -v ./...
```
