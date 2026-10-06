# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Overview
Manga Wiki is a domain-specific knowledge engine and RAG/Agent platform built on the WeKnora architecture. It provides entity-aware, event-aware, temporal, and spoiler-aware retrieval alongside structured wiki curation and graph reasoning (PRD in `docs/manga-wiki/PRD.md`).

## Architecture & Subsystems
- **Backend (`cmd/server/`, `internal/`)**: Go REST API service built with Gin. Handles session state, RAG pipelines, agent workflows, vector search, data source syncing, and MCP tools.
  - `internal/container`: Dependency injection and component lifecycle initialization.
  - `internal/agent`: ReAct reasoning loop, sandbox execution (Docker/E2B/Cube), tool dispatch.
  - `internal/application`: Core business services (knowledge base, document management, wiki engine, chat, memory).
  - `internal/infrastructure`: Database, vector store drivers (PostgreSQL/pgvector/ParadeDB, Milvus, Qdrant, OpenSearch), graph storage (Neo4j), LLM/embedding clients.
  - `internal/handler` & `internal/router`: Gin HTTP routing, middleware, SSE/WebSocket streams.
  - `internal/models` & `internal/types`: GORM database models, domain entities, and DTOs.
  - `internal/datasource`: External data connectors (Feishu, GitLab, Notion, Yuque, RSS).
- **Frontend (`frontend/`)**: Vue 3 SPA with TypeScript, Vite, Pinia, Vue Router, and TDesign (`tdesign-vue-next`). Markdown rendering via `marked`, `katex`, `mermaid`.
- **DocReader (`docreader/`)**: Python gRPC microservice using `uv` for document parsing and OCR (PDF, DOCX, PPTX, XLSX, HTML, ePUB).
- **AnyDoc (`third_party/anydoc-go/`)**: In-process Rust office parser linked via CGo build tag `anydoc`.
- **CLI (`cli/`)**: Standalone Go CLI tool for platform interactions.
- **Desktop (`cmd/desktop/`)**: Wails-based desktop application wrapper.

## Development & Build Commands

### Backend (Go)
- Build server: `make build` (or `go build -o WeKnora ./cmd/server`)
- Build with anydoc engine: `make build-anydoc` (requires `make anydoc-lib`)
- Build Lite single-binary (SQLite): `make build-lite`
- Run server: `make run` or `./WeKnora`
- Run all tests: `make test` (or `go test -v ./...`)
- Run tests in a package: `go test -v ./internal/application/...`
- Run single test: `go test -v ./internal/logger -run TestLogger`
- Run tests with race detector: `go test -race ./internal/...`
- Lint: `make lint` (or `golangci-lint run`)
- Format: `make fmt` (or `go fmt ./...`)
- Generate Swagger docs: `make docs` (requires `swag`)
- Database migrations: `make migrate-up` / `make migrate-down`

### Frontend (`frontend/`)
- Install dependencies: `npm install` (or `npm ci`)
- Development server: `npm run dev`
- Build production assets: `npm run build`
- Type checking: `npm run type-check`
- Run tests: `npm test`
- Run single test: `npx tsx --test src/i18n/localeKeyAudit.test.ts`
- Audit i18n: `npm run check-i18n`

### DocReader (`docreader/`)
- Install dependencies: `uv sync`
- Run server: `uv run python main.py` (or `make run`)
- Generate protobuf: `make proto`

### Infrastructure & Docker
- Start dev infrastructure (PostgreSQL, Redis, DocReader, Langfuse): `make dev-start`
  - Optional flags: `make dev-start DEV_ARGS="--qdrant --neo4j --minio --dex"`
- Stop dev infrastructure: `make dev-stop`
- Start all services with Docker Compose: `make start-all` (or `docker-compose up`)
