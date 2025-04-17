## 🧱 Architecture Overview

This project is a high-performance, modular messenger backend, currently implemented in Go using the [fasthttp](https://github.com/valyala/fasthttp) web framework. It is designed around a **microservice** architecture.

The `auth` module is fully functional and handles registration and secure credential management.

---

## 📁 Project Structure (Core Parts)

```
project/
├── cmd/
│   └── gateway/               # Entry point: initializes and starts HTTP server
│       ├── main.go
│       └── commonSetup.go     # Shared setup logic: DB connection, middleware
├── internal/
│   ├── auth/                  # Authentication service (sign-up, password hash, etc.)
│   │   ├── handler.go         # HTTP handlers (e.g., signup)
│   │   ├── model.go           # Data models and validation
│   │   ├── repository.go      # Storage layer (PostgreSQL and in-memory)
│   │   ├── routes.go          # Route registration for the service
│   │   └── service.go         # Business logic (user creation, password update)
│   └── ...                    # Other services (chat, user, etc.)
├── migrations/
│   └── auth/                  # SQL migrations for `auth` database
│       └── V001__init.sql     # Schema for auth.credentials table
├── configs/
│   └── config.yaml            # Configuration file (DB, server, etc.)
├── web/                       # Svelte frontend (experimental)
├── api/                       # API specs (OpenAPI planned)
└── ...
```

---

## 🔐 Auth Service

The `auth` service provides:

- **User registration**
- **Password hashing** via `bcrypt`
- **Email + phone validation** (with optional E.164 format)
- **PostgreSQL** and **in-memory** repository support (useful for testing)
- **Safe JSON serialization** (no password leakage)
- **Thread-safe repo access** using RWMutex

Planned features:
- Login endpoint
- JWT or session-based auth
- Password reset flow

---

## 🗃️ Migrations

SQL migrations for the `auth` module are located in `migrations/auth`.

Example schema (PostgreSQL):

```sql
CREATE SCHEMA IF NOT EXISTS auth;

CREATE TABLE IF NOT EXISTS auth.credentials (
    user_id UUID PRIMARY KEY,
    email TEXT UNIQUE NOT NULL,
    phone_number TEXT UNIQUE,
    passhash TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);
```

Migration tool example: [golang-migrate](https://github.com/golang-migrate/migrate)

```bash
migrate -path ./migrations/auth -database "postgres://..." up
```

---

## 🌐 Frontend (WIP)

The `web` folder contains a lightweight [Svelte](https://svelte.dev/) frontend:

- Communicates with the Go backend via HTTP (CORS enabled)
- Currently supports **sign-up** via `/api/signup`
- Plans for full auth UI and messaging interface

You’ll need to run the Svelte dev server on `localhost:5173` (default) to match CORS configuration.

---

## 🚀 Run Locally

1. **Set up PostgreSQL** and run the migrations.
2. **Configure** your DB in `configs/config.yaml`.
3. **Run the server**:

```bash
go run ./cmd/gateway
```

4. **Start the frontend** (optional):

```bash
cd web
npm install
npm run dev
```
