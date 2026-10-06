# Manga Wiki — Refactoring & Pruning Plan

**Document:** Refactoring and Scope Alignment Plan  
**Target:** Transform WeKnora v0.8.0 codebase into Manga Wiki Knowledge Engine  
**Reference:** `docs/manga-wiki/PRD.md`  

---

## 1. Executive Summary

WeKnora provides a mature foundation for RAG, Auto-Wiki, Knowledge Graph, and ReAct agents. However, it contains substantial enterprise bloat (IM bots, SaaS datasources, code sandboxes, office parsers, multi-tenant RBAC) outside the scope of Manga Wiki.

This plan details a 4-phase refactor to prune unneeded systems, decouple core dependencies, introduce the Manga Domain model, and align the UI.

---

## 2. Scope Matrix

### 2.1 Keep & Inherit
* **Wiki Engine** (`internal/application/service/wiki_*.go`): Markdown page generation, slug handling, revision history, diffs, dedup, and linting.
* **Hybrid Search / RAG** (`internal/application/service/knowledgebase_search*.go`, `chunk.go`): Vector search (pgvector, sqlite-vec), BM25 lexical search, reranking, chunk provenance.
* **Knowledge Graph** (`internal/infrastructure/neo4j`, `internal/application/service/graph.go`): Graph persistence and visualization hooks.
* **Agent Core** (`internal/agent/`): ReAct reasoning loop, question decomposition, citation tracking.
* **Frontend Core** (`frontend/src/`): Wiki reader/editor, graph explorer, chat interface with citation popovers.
* **Lite Mode** (`make build-lite`): SQLite + local storage engine for rapid local development.

### 2.2 Prune & Remove
* **Enterprise IM Integrations** (`internal/im/`, `miniprogram/`): Feishu, Lark, WeCom, WeChat, DingTalk, Telegram, Slack, Yunzhijia, QQBot.
* **Enterprise SaaS Datasources** (`internal/datasource/`): Feishu Drive/Wiki, GitLab, Notion, Yuque, Tencent IMA, RSS sync.
* **Heavy Office Parsers** (`third_party/anydoc-go/`, `docreader/`): Word, Excel, PowerPoint, Rust CGo bindings. (Retain text, markdown, JSON, HTML, PDF/EPUB parsers).
* **Code Execution Sandboxes** (`internal/sandbox/`, `tenant_skill_*`): Docker/E2B/Cube container runtimes and shell execution.
* **Desktop App** (`cmd/desktop/`, Wails): macOS/Windows desktop packaging.
* **Enterprise RBAC Bloat** (`tenant_invitation`, `tenant_member`, complex 4-tier matrix): Streamline to standard User/Admin authentication.

### 2.3 Build & Add
* **Manga Domain Entities** (`internal/types/manga/`, `internal/models/manga/`): Series, Arc, Chapter, Character, Alias, Event, Ability, Organization, Location.
* **Spoiler-Aware Boundary** (`internal/application/service/spoiler.go`): Retrieval filter based on `max_chapter` or `arc_id`.
* **Canon & Authority Model** (`internal/application/service/canon.go`): Source hierarchy (Manga Primary > Databook > Adaptation > Theory) and contradiction detection.
* **Evidence Attribution**: Grounding to Chapter, Page, and Panel/Quote instead of arbitrary doc chunk IDs.

---

## 3. Implementation Phases

```
┌─────────────────────────┐
│ Phase 1: Edge Pruning   │ ──► Cut isolated packages (IM, desktop, anydoc)
└────────────┬────────────┘
             ▼
┌─────────────────────────┐
│ Phase 2: Core Decouple  │ ──► Remove sandbox & SaaS sync from container & router
└────────────┬────────────┘
             ▼
┌─────────────────────────┐
│ Phase 3: Manga Domain   │ ──► Add Manga entities, spoiler boundary, canon engine
└────────────┬────────────┘
             ▼
┌─────────────────────────┐
│ Phase 4: Frontend Align │ ──► Strip enterprise settings, adapt Wiki/Graph/Chat
└─────────────────────────┘
```

---

### Phase 1: Edge Pruning (Isolated Modules)

**Goal:** Remove non-critical directories that have minimal coupling with the main server core.

1. **Delete non-web targets**:
   * Delete `cmd/desktop/`
   * Delete `miniprogram/`
   * Delete `third_party/anydoc-go/`
   * Delete `Formula/` (Homebrew formula)
2. **Delete IM packages**:
   * Delete `internal/im/`
   * Remove IM routes and handlers from `internal/router/` and `internal/handler/`
3. **Clean build configurations**:
   * Clean `Makefile`: remove `desktop`, `anydoc-lib`, `build-anydoc`, `miniprogram` targets.
   * Clean `docker-compose.yml` and `docker-compose.dev.yml`: remove anydoc/IM specific services.
4. **Acceptance Criteria**:
   * `go build ./cmd/server` succeeds without build tags or errors.

---

### Phase 2: Core Decoupling (Backend Services)

**Goal:** Disentangle enterprise SaaS sync and code sandboxes from the core DI container.

1. **Remove Sandboxes & Skills**:
   * Remove `internal/sandbox/`
   * Remove `tenant_skill_*.go` in `internal/application/service/`
   * Remove sandbox config and execution hooks from `internal/agent/`
2. **Remove SaaS Datasources**:
   * Remove `internal/datasource/` (keep generic file upload / directory import)
   * Remove datasource sync cron jobs and task queues from `internal/container/`
3. **Simplify Dependency Injection**:
   * Refactor `internal/container/` to instantiate only Storage, Database, LLM, Vector Store, Graph Store, and Wiki/RAG services.
   * Strip unused tables from new migrations or deprecate in GORM models.
4. **Acceptance Criteria**:
   * `go test ./internal/...` passes.
   * Server boots and serves REST API via `make run` without panics or missing dependencies.

---

### Phase 3: Manga Domain Core Implementation

**Goal:** Implement the foundational schema and logic required by `docs/manga-wiki/PRD.md`.

1. **Database Schema & Migrations**:
   * Tables: `manga_series`, `manga_arcs`, `manga_chapters`, `manga_characters`, `manga_events`, `manga_relations`, `manga_evidence`.
   * Add migrations under `migrations/`.
2. **Domain Models & DTOs**:
   * Implement structs in `internal/types/manga/` and GORM models in `internal/models/manga/`.
3. **Spoiler-Aware Retrieval**:
   * Implement spoiler filter middleware/service: query predicates inject `chapter_seq <= target_chapter_seq`.
4. **Canon & Source Authority**:
   * Tag sources with canon tiers (`primary_manga`, `official_databook`, `author_statement`, `derivative_anime`, `fan_analysis`).
5. **Manga Agent Tools**:
   * Implement ReAct agent tools:
     * `lookup_character(name_or_alias)`
     * `get_timeline(series_id, character_id, max_chapter)`
     * `get_character_relationships(character_id)`
     * `search_manga_evidence(query, max_chapter)`
6. **Acceptance Criteria**:
   * Unit tests for spoiler boundaries pass.
   * End-to-end test querying character relations and chapter-bounded timeline succeeds.

---

### Phase 4: Frontend Pruning & Alignment

**Goal:** Make the web UI reflect Manga Wiki rather than an enterprise knowledge portal.

1. **Prune Enterprise Views**:
   * Remove IM integration settings, sandbox runtime monitors, and SaaS datasource configs.
   * Remove unused routes in `frontend/src/router/`.
2. **Customize Core Views**:
   * **Wiki View**: Add Manga metadata cards (Character infobox: Debut chapter, Arc appearances, Affiliations, Status).
   * **Graph View**: Adapt relationship types (`Ally`, `Enemy`, `Master/Student`, `Family`) and event connections.
   * **Chat View**: Add chapter/spoiler slider filter to conversation header; format citations as `[Ch. {num}, p. {page}]`.
3. **Clean Dependencies**:
   * Prune unused npm dependencies in `frontend/package.json`.
4. **Acceptance Criteria**:
   * `npm run build` passes.
   * `npm run type-check` passes with zero errors.

---

## 4. Risk Mitigation

1. **Breaking Existing RAG / Wiki Functions**:
   * Keep `knowledgebase.go`, `wiki_page.go`, and `chunk.go` untouched during Phase 1 & 2. Only remove foreign callers.
2. **Database Migration Conflicts**:
   * Keep original migration sequence intact; add new manga tables as new versioned migration files rather than rewriting history.
3. **Rollback Strategy**:
   * Commit after each phase. Each phase must leave the repository in a cleanly compiling and runnable state.
