# grape-shared

Shared Go module for Grape microservices. Consumed by each service via `go.mod`:

```go
require github.com/your-moon/grape-shared v0.1.0
```

## Contents

- **proto/** – Protobuf definitions and generated Go (auth, engagement, product, jobs, market, chat, feed).
- **domain/** – Shared domain types and interfaces (Product, Job, Market, Chat, User, etc.).
- **dto/** – Request/response DTOs and mappers (dmo: bounded-context to shared conversion).
- **common/** – config, constants, database, messaging, types, utils.
- **entities/** – GORM models (shared table structs).
- **bc/** – Bounded-context domain structs (jobs, product, market) used by dmo.

## Versioning

Tag releases with semver (e.g. v0.1.0). Services pin a version in their go.mod.

## Regenerate proto

```bash
make proto
```

## Build

```bash
go build ./...
```
