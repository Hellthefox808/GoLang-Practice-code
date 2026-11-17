# 17 - Defer Mechanics, LIFO Execution, and Panic/Recover

**Author:** Ravi Ranjan Singh  
**Repository:** [GoLang-Practice-code](https://github.com/Hellthefox808/GoLang-Practice-code)  
**License:** MIT  

---

## 1. Overview
The `defer` keyword schedules a function call to be executed immediately after the surrounding function returns.

---

## 2. Fundamental Defer Rules
1. **LIFO Ordering:** Deferred calls are placed onto a call stack and popped in Last-In, First-Out order.
2. **Immediate Argument Evaluation:** The parameters supplied to a deferred function are evaluated **at the moment the defer statement is encountered**, not when the function finally runs.
3. **Closure Exceptions:** If the deferred function is an anonymous closure accessing variables from the outer scope, it reads their values at the moment of execution.
4. **Panic and Recover:** `recover()` is useful strictly inside deferred functions. It halts an active panic unwinding sequence and restores normal control flow.

---

## 3. How to Run & Verify
```bash
go run main.go
go test -v ./...
```
