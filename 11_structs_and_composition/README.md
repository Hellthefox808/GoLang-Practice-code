# 11 - Structs, Composition, and Memory Alignment

**Author:** Ravi Ranjan Singh  
**Repository:** [GoLang-Practice-code](https://github.com/Hellthefox808/GoLang-Practice-code)  
**License:** MIT  

---

## 1. Overview & Architectural Philosophy
Go deliberately rejects classical Object-Oriented inheritance (`extends`, `super`, virtual method tables). Instead, Go embraces **Composition over Inheritance** through struct embedding and field promotion.

---

## 2. Struct Embedding & Method Promotion
When a struct type is embedded anonymously inside another struct, all of its fields and methods are automatically promoted to the outer struct:

```go
type BaseEntity struct {
    ID        string
    CreatedAt time.Time
}

type User struct {
    BaseEntity // Embedded struct (Composition)
    Email      string
}
```

Promoted fields can be accessed directly:
```go
u := User{BaseEntity: BaseEntity{ID: "usr_101"}}
fmt.Println(u.ID) // Direct access without u.BaseEntity.ID
```

---

## 3. Hardware Memory Alignment & Struct Padding

CPUs read and write memory in discrete word boundaries (8 bytes on 64-bit platforms). When struct fields are ordered haphazardly, the compiler is forced to inject zero-padding bytes to align memory addresses.

### Unoptimized Struct (24 Bytes)
```go
type Unoptimized struct {
    A bool  // 1 byte  + 7 bytes padding
    B int64 // 8 bytes (aligned at offset 8)
    C bool  // 1 byte  + 7 bytes padding
} // Total: 24 bytes
```

### Optimized Struct (16 Bytes — 33% Memory Reduction)
```go
type Optimized struct {
    B int64 // 8 bytes (offset 0)
    A bool  // 1 byte  (offset 8)
    C bool  // 1 byte  (offset 9)
            // + 6 bytes trailing padding to round to word boundary
} // Total: 16 bytes
```

**Rule of Thumb:** Order struct fields from largest byte size to smallest (e.g., `int64`, `float64`, `pointers` first -> `int32`, `float32` -> `int16` -> `bool`, `byte`).

---

## 4. Verification & Testing
```bash
go run main.go
go test -v ./...
```
