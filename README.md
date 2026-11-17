# GoLang-Practice-Code: Production-Grade Go Engineering Masterclass

[![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Author](https://img.shields.io/badge/Author-Ravi%20Ranjan%20Singh-blue)](https://github.com/Hellthefox808)
[![Build Status](https://img.shields.io/badge/Build-Passing-brightgreen)]()

An enterprise-grade, highly structured, production-hardened reference curriculum for mastering **Go (Golang)**. Designed, enhanced, organized, and authored by **Ravi Ranjan Singh**.

Every module adheres to the **Zero-Placeholder Law**: no `// TODO` stubs, complete memory-safe implementations, explicit error handling, automated table-driven tests, and deep architectural explanations.

---

## 📑 Curriculum & Architecture Roadmap

The repository is modularized into 28 progressive engineering disciplines, managed as a unified Go Workspace (`go.work`):

| # | Module Directory | Core Technical Topics & Idioms | Status & Tests |
|---|---|---|---|
| **01** | [`01_hello_world`](./01_hello_world) | Entrypoint contract, `init()` sequence, `runtime` diagnostics, `io.Writer` decoupling | Complete ✅ |
| **02** | [`02_variables_and_types`](./02_variables_and_types) | Static type system, zero-value guarantees, `iota` enumerations, memory bounds | Complete ✅ |
| **03** | [`03_user_input`](./03_user_input) | `bufio.Reader`, `bufio.Scanner`, cross-platform CRLF normalization, mockable streams | Complete ✅ |
| **04** | [`04_type_conversions`](./04_type_conversions) | `strconv` parsing, safe numeric downcasting, overflow interception, byte slice casting | Complete ✅ |
| **05** | [`05_math_and_crypto`](./05_math_and_crypto) | PRNG vs CSPRNG, `crypto/rand` token generation, `math/big` arbitrary precision | Complete ✅ |
| **06** | [`06_time_handling`](./06_time_handling) | Reference date `Mon Jan 2 15:04:05 MST 2006`, IANA timezones, Monotonic clock | Complete ✅ |
| **07** | [`07_pointers_and_memory`](./07_pointers_and_memory) | Pointer mechanics, pass-by-value fundamentals, escape analysis, nil pointer guards | Complete ✅ |
| **08** | [`08_arrays`](./08_arrays) | Fixed `[N]T` value semantics, type identity, matrix inversion, cache locality | Complete ✅ |
| **09** | [`09_slices_and_internals`](./09_slices_and_internals) | SliceHeader (ptr, len, cap), capacity growth, shared backing arrays, `slices` package | Complete ✅ |
| **10** | [`10_maps_and_hash_tables`](./10_maps_and_hash_tables) | `runtime.hmap` buckets, comma-ok idiom, non-deterministic iteration, `maps` package | Complete ✅ |
| **11** | [`11_structs_and_composition`](./11_structs_and_composition) | Composition over inheritance, struct embedding, promotion, memory padding alignment | Complete ✅ |
| **12** | [`12_conditionals`](./12_conditionals) | Guard clauses (early return pattern), short-statement `if`, boolean short-circuiting | Complete ✅ |
| **13** | [`13_switch_control`](./13_switch_control) | Expression switches, automatic breaks, tagless boolean ladders, dynamic type switches | Complete ✅ |
| **14** | [`14_loops_and_iterations`](./14_loops_and_iterations) | `for` variants, UTF-8 rune iteration, labeled breaks, Go 1.22+ per-iteration scoping | Complete ✅ |
| **15** | [`15_functions_and_variadics`](./15_functions_and_variadics) | First-class functions, stateful closures, variadic unpacking (`...T`), generic map | Complete ✅ |
| **16** | [`16_methods_and_interfaces`](./16_methods_and_interfaces) | Value vs Pointer receivers, implicit duck typing, interface `(type, value)` internals | Complete ✅ |
| **17** | [`17_defer_and_cleanup`](./17_defer_and_cleanup) | LIFO stack execution, immediate argument binding, named return mutation, `recover()` | Complete ✅ |
| **18** | [`18_file_io_operations`](./18_file_io_operations) | Atomic file writes (`CreateTemp` -> `Sync` -> `Rename`), buffered streaming, `os.ReadFile` | Complete ✅ |
| **19** | [`19_web_requests`](./19_web_requests) | Custom `http.Client` timeouts, context deadline propagation, socket leak prevention | Complete ✅ |
| **20** | [`20_url_parsing`](./20_url_parsing) | RFC 3986 parsing, `RawPath` vs `RawQuery` resolution, `url.Values` query encoding | Complete ✅ |
| **21** | [`21_http_verbs_and_clients`](./21_http_verbs_and_clients) | GET, POST (JSON), POST (Form), custom headers, in-memory `httptest.Server` mocks | Complete ✅ |
| **22** | [`22_json_handling`](./22_json_handling) | Struct tags (`omitempty`, `-`), streaming `json.Decoder`, `DisallowUnknownFields` guard | Complete ✅ |
| **23** | [`23_modules_and_vendoring`](./23_modules_and_vendoring) | `go.mod`/`go.sum`, Go 1.22+ stdlib routing (`GET /path`), middleware, graceful shutdown | Complete ✅ |
| **24** | [`24_rest_api_gorilla_mux`](./24_rest_api_gorilla_mux) | Thread-safe CRUD (`sync.RWMutex`), standard JSON envelopes, HTTP semantic statuses | Complete ✅ |
| **25** | [`25_mongodb_crud_api`](./25_mongodb_crud_api) | Enterprise MongoDB driver (`v1.15+`), 12-Factor `MONGO_URI`, ObjectID validation, health ping | Complete ✅ |
| **26** | [`26_goroutines_concurrency`](./26_goroutines_concurrency) | M:N scheduler (G, M, P), `sync.WaitGroup` mechanics, bounded semaphore concurrency | Complete ✅ |
| **27** | [`27_mutex_and_waitgroups`](./27_mutex_and_waitgroups) | `sync.Mutex` vs `sync.RWMutex`, critical section minimizing, lock-free `sync/atomic` | Complete ✅ |
| **28** | [`28_channels_and_select`](./28_channels_and_select) | CSP communication, channel directionality, multiplexing `select`, Worker Pool pattern | Complete ✅ |
| **Mock** | [`mock_server`](./mock_server) | Dual Go & Node.js test mock servers for offline local client development | Complete ✅ |

---

## 🚀 Getting Started

### Prerequisites
- **Go 1.22+** installed on your operating system. Check your version with:
  ```bash
  go version
  ```

### Workspace Setup
This repository uses standard **Go Workspaces (`go.work`)** to link all 28 modules seamlessly:
```bash
# Clone the repository
git clone https://github.com/Hellthefox808/GoLang-Practice-code.git
cd GoLang-Practice-code

# Verify workspace sync
go work sync
```

---

## 🧪 Testing & Verification

Every module contains isolated table-driven unit tests. You can run all tests across the entire repository simultaneously:

```bash
# Run all unit tests across all workspace modules
go test ./... -v

# Run race condition verification on concurrent modules
go test -race ./24_rest_api_gorilla_mux/... ./26_goroutines_concurrency/... ./27_mutex_and_waitgroups/... ./28_channels_and_select/...
```

---

## 🛡️ Code Quality & Engineering Enforcement
- **Zero-Placeholder Law:** Every function, struct, and handler is fully written and immediately functional.
- **Memory Safety:** Pointer operations, slices, and maps are defended against nil pointer dereferences and concurrent write panics.
- **Resource Discipline:** All open file handles, database cursors, HTTP response bodies, and ticker channels are explicitly deferred, drained, and closed.
- **Security Hardening:** MongoDB connection strings and secret keys are externalized through environment variables; JSON decoders disallow unexpected field injection.

---

## 👤 Author & Ownership
- **Lead Architect & Engineer:** **Ravi Ranjan Singh**
- **GitHub:** [@Hellthefox808](https://github.com/Hellthefox808)
- **Repository:** [https://github.com/Hellthefox808/GoLang-Practice-code](https://github.com/Hellthefox808/GoLang-Practice-code)
- **License:** [MIT License](./LICENSE)
