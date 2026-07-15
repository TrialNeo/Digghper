<p align="center">
  <img src="./ico.png" width="200" alt="Digghper"/>
</p>

<h1 align="center">Digghper</h1>

<p align="center">
  <strong>dig + gopher = digghper</strong><br/>
  向真木老哥致敬
</p>

<p align="center">
  轻量 Go 后端脚手架 · 分层清晰 · 与 WatchAuth 代码风格对齐
</p>

---

## 项目介绍

Digghper 是一套面向业务快速落地的 Go 后端基础框架，约定清晰的分层结构、统一的响应与错误码、JWT 鉴权与可替换组件，适合作为新服务的起点。

当前内置能力：

- 管理员登录（JWT Bearer）
- Fiber HTTP 服务 + CORS
- GORM 多驱动（PostgreSQL / MySQL）
- Redis 缓存接入
- Zap + lumberjack 结构化日志
- 统一响应体与错误码（`errMsg`）

更细的开发约定见 [AGENTS.md](./AGENTS.md)。

## 技术栈

| 组件 | 选型 |
|------|------|
| 语言 | Go 1.24+ |
| Web | Fiber v2 |
| ORM | GORM（postgres / mysql） |
| 缓存 | go-redis v8 |
| 配置 | Viper |
| 日志 | zap + lumberjack |
| 鉴权 | golang-jwt + bcrypt |

## 环境要求

- Go 1.24+
- PostgreSQL 或 MySQL
- Redis（可选，视业务是否使用）

## 快速开始

### 1. 配置

编辑 [`configs/config.yaml`](./configs/config.yaml)：

```yaml
web:
  port: 9090

database:
  driverName: postgres          # 或 mysql
  dataSourceName: diggpher
  host: localhost
  port: 5432
  user: postgres
  psw: "your-password"
  timeZone: Asia/Shanghai

redis:
  addr: localhost:6379
  password: ""

logger:
  level: info
  console: true
  dir: ./logs
```

### 2. 安装依赖并启动

```shell
go mod tidy
go run cmd/main.go
```

服务默认监听 `:9090`。

### 3. 构建

```shell
cd cmd
go build -ldflags "-s -w" -o digghper.exe
```

## 项目结构

```
Digghper/
├── cmd/                    # 入口 main.go
├── configs/                # 配置文件（无业务代码）
├── global/                 # 全局单例与配置结构
├── initialize/             # 启动引导：配置 / DB / Redis / Fiber
├── internal/               # 业务内部包（禁止被外部模块 import）
│   ├── controller/         # HTTP 处理层
│   ├── service/            # 业务逻辑
│   │   └── errMsg/         # 错误码与中文文案
│   ├── dao/                # 模型与数据访问
│   ├── route/              # 路由绑定
│   ├── job/                # 定时 / 后台任务
│   └── client/             # 外部客户端（如 RPC）
├── pkg/                    # 可复用工具
│   ├── crypto/             # 密码加密
│   ├── logger/             # 日志
│   ├── middleware/auth/    # JWT 中间件
│   └── utils/
├── AGENTS.md               # AI / 协作开发约定
└── readme.md
```

## 分层约定

```
cmd/main.go
  → initialize/*          # 配置、Redis、DB、Web
  → internal/route/       # 路由绑定
  → controller/           # 解析请求、组装响应
  → service/              # 业务逻辑
  → dao/                  # 持久化
```

| 目录 | 职责 |
|------|------|
| `cmd` | 进程入口，只做启动编排 |
| `configs` | 配置文件，不放代码 |
| `global` | `CONFIG` / `DataBase` / `Redis` / `WebApp` / `Log` |
| `initialize` | 引导失败直接 `panic` |
| `internal` | 仅包内互调，可依赖 `pkg` / 第三方 |
| `pkg` | 可复用、无业务状态的工具 |

## 关键 API

| 方法 | 路径 | 鉴权 | 说明 |
|------|------|------|------|
| POST | `/api/admin/login` | 无 | 管理员登录，返回 JWT |
| * | `/api/admin/*` | Bearer JWT | 需管理员 Token（`RequireTokenType`） |

登录示例：

```shell
curl -X POST http://localhost:9090/api/admin/login \
  -H "Content-Type: application/json" \
  -d "{\"username\":\"admin\",\"password\":\"your-password\"}"
```

后续请求携带：

```http
Authorization: Bearer <token>
```

## 代码风格（与 WatchAuth 对齐）

1. 文件名小写：`admin.go`，不要 `Admin.go`
2. Service 按职责拆分：`admin.go`（类型 + 构造），`admin_login.go`（登录逻辑）
3. Controller 统一响应：

   ```go
   re := newRespondIMP(c)
   re.withCode(code).Respond(data)
   ```

4. 路由：导出 `WithAdminRoute`，包内 `adminRouters` 结构体
5. 错误码：`errMsg` 无类型常量 + `GetErrMsg`
6. 初始化致命错误：`panic(err)`
7. 日志：仅通过 `global.Log` / `global.SugarLog`（在 `LoadConfigs` 之后）
8. import 分组：标准库 → 空行 → 第三方 / 本模块（gofmt）

完整约定见 [AGENTS.md](./AGENTS.md)。

## 规范说明

1. 每个目录宜有独立 `readme.md`，说明本包职责
2. 业务代码风格与 WatchAuth 对齐，避免引入另一套分层习惯

## 更换 Web 框架

若替换 Fiber，优先改这些位置：

| 路径 | 变量 / 符号 | 说明 |
|------|-------------|------|
| `global/const.go` | `FbConfig` | 框架相关配置常量 |
| `global/global.go` | `WebApp` | HTTP 应用实例 |
| `initialize/fiber.go` | `RunWebService` | 启动与中间件 |
| `internal/route/*` | `BindRoute` / `WithXxxRoute` | 路由注册 |
| `internal/controller/*` | handler 签名 | 请求上下文类型 |

## 可替换组件

| 路径 | 说明 |
|------|------|
| `pkg/crypto/encrypt.go` | 登录密码哈希 / 校验 |
| `pkg/logger/logger.go` | 日志实现（zap + 滚动文件） |
| `pkg/middleware/auth/` | JWT 签发与鉴权中间件 |
| `configs/config.yaml` | 数据库驱动、端口、日志级别 |

## 参考资料

- [Fiber 文档](https://docs.gofiber.io/)
- [GORM 中文文档](https://gorm.io/zh_CN/docs/index.html)
- [go-redis 指南](https://redis.ac.cn/docs/latest/develop/clients/go/)
- [Viper](https://github.com/spf13/viper)
- [zap](https://github.com/uber-go/zap)

## 贡献

欢迎提交 Issue 与 Pull Request。新增能力时请保持与现有分层、响应格式和错误码风格一致。