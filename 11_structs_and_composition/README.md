# 11 - Structs, Composition, and Memory Alignment

**Author:** Ravi Ranjan Singh  
**Repository:** [GoLang-Practice-code](https://github.com/Hellthefox808/GoLang-Practice-code)  
**License:** MIT  

---

## 1. Overview
Go rejects classical object-oriented class hierarchies (`extends`, `super`, virtual method tables). Instead, Go embraces **Composition over Inheritance** through anonymous struct embedding and field promotion.

---

## 2. Key Paradigms
1. **Struct Embedding:**
   Placing an un-named struct type inside another struct automatically promotes all fields and methods of the embedded type to the outer struct.
2. **Constructors as Pure Functions:**
   Because Go lacks automatic constructor methods, the convention is `New<Type>(args...) (*<Type>, error)`. This ensures invariants and validation rules are satisfied before memory references escape.
3. **Hardware Memory Alignment (Padding):**
   CPUs read memory in word boundaries (8 bytes on 64-bit systems). Misaligned fields force the compiler to inject zero-padding bytes. Arranging fields in decreasing order of byte size (e.g., 8-byte types first, followed by 4, 2, and 1-byte types) significantly reduces memory footprints in large data pipelines.

---

## 3. How to Run & Verify
```bash
go run main.go
go test -v ./...
```
