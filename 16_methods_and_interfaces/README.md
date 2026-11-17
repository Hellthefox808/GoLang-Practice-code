# 16 - Methods, Value vs Pointer Receivers, and Interfaces

**Author:** Ravi Ranjan Singh  
**Repository:** [GoLang-Practice-code](https://github.com/Hellthefox808/GoLang-Practice-code)  
**License:** MIT  

---

## 1. Overview
In Go, interfaces are satisfied **implicitly**. If a concrete type implements all the method signatures declared by an interface, it satisfies that interface automatically with no `implements` keyword needed.

---

## 2. Choosing Between Receiver Types
| Criterion | Value Receiver `(t T)` | Pointer Receiver `(t *T)` |
|---|---|---|
| **State Mutation** | Cannot modify caller | Modifies caller directly |
| **Data Size** | Small, fixed structs | Large structs (avoids copy overhead) |
| **Consistency** | Use when all methods are values | If any method needs a pointer, make all methods pointers |
| **Interface Satisfaction** | Satisfies both `T` and `*T` | Satisfies only `*T` |

---

## 3. The Interface Pair Under the Hood
An interface variable is internally a 2-word structure `(type, value)`:
1. Pointer to the concrete type information (itab).
2. Pointer to the underlying value.
*The Nil Interface Trap:* An interface holding a concrete pointer whose value is `nil` is **NOT nil** because its type descriptor is populated. Always be mindful when returning concrete `nil` pointers as interface types.

---

## 4. How to Run & Verify
```bash
go run main.go
go test -v ./...
```
