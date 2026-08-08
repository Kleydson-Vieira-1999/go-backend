# Project Structure Summary

This document provides a high-level overview of the structure of the Restaurant Orders project.

---

## 📂 Directory Layout

```
.
├── .antigravityrules        # Rules & guidelines for AI agent operations
├── .antigravity-rules.md     # Markdown version of AI agent guidelines
├── flake.nix                # Nix Flakes environment setup and aliases
├── flake.lock               # Pin lock file for Nix inputs
├── backend/                 # Golang Modular Backend
├── frontend/                # Next.js 16 Frontend
└── infra/                   # Infrastructure configuration (Database, Docker, Diagrams)
```

---

## ⚙️ Nix Development Environment (`flake.nix`)
The project utilizes **Nix Flakes** to standardize tools across development systems. When entering the shell via `nix develop`, the environment sets up:
- **Go / gopls / delve** (backend development & debugging)
- **Node.js** (frontend development)
- **Environment variables** (GOPATH, GOCACHE, global npm prefix)
- **Aliases**:
  - `testar-dev` / `stop-dev` (Run/stop dev infrastructure in Podman/Docker)
  - `testar-prod` / `stop-prod` (Run/stop production-like containers)

---

## 🖥️ Backend Module Structure (`backend/`)
The backend is structured modularly. Each domain is isolated within [backend/modules](file:///home/doctor/develop/resturant-orders/backend/modules).

Every module except `core` is completely flattened under its own Go package (e.g. `package menu`, `package store`):
- **Unified Package**: Route routing (`router.go`), request controllers/handlers (`create.go`, `select.go`, etc.), database GORM models (`model.go`), and migrations registry (`setup.go`) reside directly in the module's root folder under the same package.
- **No Sub-Packages**: There are no subdirectories (like `controllers`, `models`, `routes`) within the domain modules.

Modules:
- [core/](file:///home/doctor/develop/resturant-orders/backend/modules/core): Contains shared infrastructure (DB connection, global routing, standard response formats, SSE broker) and is NOT flattened.
  - `middleware/`: Common middlewares (e.g., JWT Auth and Token Builder) to break circular imports between modules.
- [auth/](file:///home/doctor/develop/resturant-orders/backend/modules/auth): Google OAuth 2.0 implementation.
- [store/](file:///home/doctor/develop/resturant-orders/backend/modules/store): Management of restaurants/establishments.
- [user/](file:///home/doctor/develop/resturant-orders/backend/modules/user): User accounts and profiles.
- [kitchen/](file:///home/doctor/develop/resturant-orders/backend/modules/kitchen): Kitchen SSE event streaming.
- [menu/](file:///home/doctor/develop/resturant-orders/backend/modules/menu): CRUD for restaurant menus.
- [product/](file:///home/doctor/develop/resturant-orders/backend/modules/product): CRUD for dishes/products.
- [order/](file:///home/doctor/develop/resturant-orders/backend/modules/order): Creating and managing orders, broadcasting updates.
- [table/](file:///home/doctor/develop/resturant-orders/backend/modules/table): Restaurant tables & table sessions.
- [waiter/](file:///home/doctor/develop/resturant-orders/backend/modules/waiter): Waiter authentication access codes.

---

## 🎨 Frontend Architecture (`frontend/`)
The frontend is a **Next.js 16** application with **TypeScript** and **Tailwind CSS v4**:
- **`app/`**: Next.js App Router folders.
  - `(loged)/`: Authenticated sections of the app.
  - `(work)/`: Kitchen/waiter workflow interfaces.
  - `callback/`: OAuth callback page.
- **`components/`**: Reusable UI components.
- **`lib/`**: Next.js utility functions.
- **`services/`**: API connectors and SSE listeners.
- **`types/`**: TypeScript type definitions.
- **Global State**: Managed using **Redux Toolkit** and persisted with **Redux Persist**.

---

## 📦 Infrastructure & Database Configuration (`infra/`)
- [diagram.mmd](file:///home/doctor/develop/resturant-orders/infra/diagram.mmd): Entity-Relationship diagram in Mermaid format.
- [enums.sql](file:///home/doctor/develop/resturant-orders/infra/enums.sql): Declares PostgreSQL custom types (`order_status`, `role_type`) required before running backend GORM migrations.
- [database.sql](file:///home/doctor/develop/resturant-orders/infra/database.sql) & [test.sql](file:///home/doctor/develop/resturant-orders/infra/test.sql): DB backup/restoration schemas.
- `docker-compose.yml` & `docker-compose.prod.yml`: Docker-compose orchestration manifests configuration for PostgreSQL, PgAdmin, and optional stack tools (Elasticsearch, Kibana, RabbitMQ, Filebeat).
