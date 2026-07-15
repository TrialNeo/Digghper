# AGENTS.md

This file provides guidance to Codex when working with code in this repository.

## Development Commands

### Backend (Go, module: `Diggpher`, go 1.24.0)
- **Start dev server**: `go run cmd/main.go` (listens on `:9090`)
- **Add dependencies**: `go get <pkg>` then `go mod tidy`
- **Build binary**: `cd cmd && go build -ldflags "-s -w" -o digghper.exe`

## Architecture

### Backend Layering
```
cmd/main.go → initialize/* → internal/route/ → internal/controller/ → internal/service/ → internal/dao/
```
- **`initialize/`** — Bootstrap: config loading (viper from `configs/config.yaml`), PostgreSQL/MySQL (GORM), Redis, Fiber app with CORS, then route binding
- **`global/`** — Package-level singletons: `CONFIG`, `DataBase` (*gorm.DB), `Redis` (*redis.Client), `WebApp` (*fiber.App), `Log`/`SugarLog` (zap)
- **`internal/controller/`** — Fiber handlers, parse request, call service, respond via `newRespondIMP`/`Respond()`
- **`internal/service/`** — Business logic, coordinates DAOs and Redis cache; constructor via `NewXxxService()`
- **`internal/dao/`** — GORM model definitions + AutoMigrate in `BindDao()`; file names are lowercase
- **`internal/service/errMsg/`** — Centralized error codes and Chinese error messages
- **`pkg/middleware/auth/`** — JWT generation (`GenerateAdminToken` etc.) and Bearer token middleware + `RequireTokenType`
- **`pkg/logger/`** — Zap + lumberjack (console + file rotation), writes into `global.Log`/`global.SugarLog`

### Key Routes
- `/api/admin/login` (POST, no auth) — Admin login
- `/api/admin/*` (with JWT middleware) — Protected admin APIs

### Configuration
- Backend: `configs/config.yaml` (DB, Redis, server port, logger)
- Backend config struct in `global/config.go`

## Code Style (aligned with WatchAuth)

1. File names: lowercase (`admin.go`, not `Admin.go`)
2. Service split by concern: `admin.go` holds type + constructor, `admin_login.go` holds login logic
3. Controller responses: always `re := newRespondIMP(c)` then `re.withCode(code).Respond(data)`
4. Route binding: exported `WithAdminRoute`, package-private `adminRouters` struct
5. Errors: untyped const codes in `errMsg`, messages via `GetErrMsg`
6. Init side effects use `panic(err)` for fatal bootstrap failures
7. Logging: only via `global.Log` after `initialize.LoadConfigs()`
8. Import groups: stdlib, blank line, third-party / local modules (gofmt)