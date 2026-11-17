# 12 - Conditionals, Guard Clauses, and Scoped If Statements

**Author:** Ravi Ranjan Singh  
**Repository:** [GoLang-Practice-code](https://github.com/Hellthefox808/GoLang-Practice-code)  
**License:** MIT  

---

## 1. Overview
In Go, conditionals are designed to avoid defensive nesting (the "arrow-code" anti-pattern) and restrict variable scope tightly to the decision block.

---

## 2. Key Language Idioms
1. **Short Variable Initialization:**
   ```go
   if err := executeOperation(); err != nil {
       return fmt.Errorf("operation failed: %w", err)
   }
   ```
   `err` exists only inside the `if` body, avoiding namespace pollution.
2. **Guard Clauses (Early Returns):**
   Evaluate edge cases and failure preconditions first and return immediately. The "happy path" remains unindented at the root level of the function.

---

## 3. How to Run & Verify
```bash
go run main.go
go test -v ./...
```
