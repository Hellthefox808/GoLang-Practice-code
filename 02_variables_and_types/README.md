# 02 - Variables, Static Types, and Zero Values

**Author:** Ravi Ranjan Singh  
**Repository:** [GoLang-Practice-code](https://github.com/Hellthefox808/GoLang-Practice-code)  
**License:** MIT  

---

## 1. Overview
Go is a statically typed, compiled programming language with explicit type enforcement. All variables must possess a known type at compilation time. Unlike C, Go variables initialized without an explicit value are guaranteed to be initialized to their respective **Zero Value** by the runtime, eradicating uninitialized memory access vulnerabilities.

---

## 2. Core Concepts & Zero Values Table
| Type Category | Specific Types | Zero Value |
|---|---|---|
| **Numeric (Integer)** | `int`, `int8`..`int64`, `uint`, `uint8`..`uint64`, `uintptr` | `0` |
| **Numeric (Floating)** | `float32`, `float64` | `0.0` |
| **Complex** | `complex64`, `complex128` | `(0+0i)` |
| **Boolean** | `bool` | `false` |
| **String** | `string` | `""` (empty string) |
| **Reference / Pointer** | Pointer (`*T`), Slice (`[]T`), Map (`map[K]V`), Channel (`chan T`), Interface (`any`), Function | `nil` |

---

## 3. Variable Declaration Styles
1. **Explicit Type:** `var x int = 10`
2. **Type Inference:** `var x = 10` (compiler infers `int` from the literal)
3. **Short Declaration:** `x := 10` (valid only inside function bodies)
4. **Block Declaration:**
   ```go
   var (
       host = "localhost"
       port = 8080
   )
   ```

---

## 4. How to Run & Verify
```bash
go run main.go
go test -v ./...
```
