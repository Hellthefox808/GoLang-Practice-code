# 16 - Methods, Value vs Pointer Receivers, and Interfaces

**Author:** Ravi Ranjan Singh  
**Repository:** [GoLang-Practice-code](https://github.com/Hellthefox808/GoLang-Practice-code)  
**License:** MIT  

---

## 1. Overview & Method Receivers
In Go, methods are functions declared with a **receiver** parameter binding them to a specific type. Go provides two receiver categories with fundamentally distinct behaviors:

| Receiver Category | Declaration Syntax | Mutation Behavior | Performance Implication |
|---|---|---|---|
| **Value Receiver** | `func (u User) Get()` | Operates on an isolated copy; mutations **do not** affect the caller | Suitable for small, immutable types ($< 64$ bytes) |
| **Pointer Receiver** | `func (u *User) Set()` | Operates directly on caller's memory; mutations **modify** original struct | Avoids large struct copy overhead; mandatory for mutation |

---

## 2. Implicit Interface Satisfaction (Duck Typing Done Right)
Go interfaces are satisfied **implicitly**. A concrete type satisfies an interface by implementing all required methods; no `implements` keyword exists:

```
        Notifier Interface
         +--------------+
         | Notify(...)  |
         +--------------+
                ^
     +----------+----------+
     |                     |
UserAccount           SlackChannel
(implements Notify)   (implements Notify)
```

### The Interface Internal Pair: `(itab, data)`
Under the hood, an interface variable is a 2-word structure:
1. `itab`: Pointer to metadata describing the concrete type and its method table.
2. `data`: Pointer to the underlying concrete value.

### The Dangerous "Nil Interface Trap"
An interface is considered `nil` **only when both its type and value are nil**. If a function returns a concrete pointer whose value is `nil` as an interface, the interface itself is **NOT nil**:
```go
func GetNotifier() Notifier {
    var u *UserAccount = nil
    return u // BUG: Returns non-nil interface holding a nil pointer!
}

n := GetNotifier()
if n != nil { // Evaluates to TRUE!
    n.Notify("alert") // PANIC: nil pointer dereference inside method
}
```
**Fix:** Always return the untyped `nil` literal when returning nil interface values:
```go
if u == nil {
    return nil // Clean untyped nil
}
```

---

## 3. Verification & Testing
```bash
go run main.go
go test -v ./...
```
