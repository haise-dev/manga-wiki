-- Manga Domain Core Tables (SQLite)
-- Implements entities from Manga Wiki PRD:
-- Series, Arcs, Chapters, Characters, Events, Relations, Evidence

CREATE TABLE IF NOT EXISTS manga_series (
    id                VARCHAR(64) PRIMARY KEY,
    tenant_id         INTEGER NOT NULL,
    title             VARCHAR(255) NOT NULL,
    romaji_title      VARCHAR(255) NOT NULL DEFAULT '',
    english_title     VARCHAR(255) NOT NULL DEFAULT '',
    author            VARCHAR(255) NOT NULL DEFAULT '',
    status            VARCHAR(64) NOT NULL DEFAULT 'ongoing',
    total_chapters    INTEGER NOT NULL DEFAULT 0,
    created_at        DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at        DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_manga_series_tenant_title
    ON manga_series (tenant_id, title);

CREATE TABLE IF NOT EXISTS manga_arcs (
    id                VARCHAR(64) PRIMARY KEY,
    tenant_id         INTEGER NOT NULL,
    series_id         VARCHAR(64) NOT NULL,
    name              VARCHAR(255) NOT NULL,
    arc_seq           INTEGER NOT NULL DEFAULT 0,
    start_chapter_seq INTEGER NOT NULL DEFAULT 0,
    end_chapter_seq   INTEGER NOT NULL DEFAULT 0,
    created_at        DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at        DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_manga_arcs_series_seq
    ON manga_arcs (tenant_id, series_id, arc_seq);

CREATE TABLE IF NOT EXISTS manga_chapters (
    id                VARCHAR(64) PRIMARY KEY,
    tenant_id         INTEGER NOT NULL,
    series_id         VARCHAR(64) NOT NULL,
    arc_id            VARCHAR(64) NOT NULL DEFAULT '',
    chapter_number    VARCHAR(32) NOT NULL,
    chapter_seq       INTEGER NOT NULL DEFAULT 0,
    title             VARCHAR(255) NOT NULL DEFAULT '',
    release_date      DATETIME,
    created_at        DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at        DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE UNIQUE INDEX IF NOT EXISTS uq_manga_chapters_seq
    ON manga_chapters (tenant_id, series_id, chapter_seq);

CREATE TABLE IF NOT EXISTS manga_characters (
    id                          VARCHAR(64) PRIMARY KEY,
    tenant_id                   INTEGER NOT NULL,
    series_id                   VARCHAR(64) NOT NULL,
    canonical_name              VARCHAR(255) NOT NULL,
    aliases                     TEXT NOT NULL DEFAULT '[]',
    status                      VARCHAR(64) NOT NULL DEFAULT 'alive',
    first_appearance_chapter_seq INTEGER NOT NULL DEFAULT 0,
    created_at                  DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at                  DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_manga_characters_name
    ON manga_characters (tenant_id, series_id, canonical_name);

CREATE INDEX IF NOT EXISTS idx_manga_characters_first_app
    ON manga_characters (tenant_id, series_id, first_appearance_chapter_seq);

CREATE TABLE IF NOT EXISTS manga_events (
    id                VARCHAR(64) PRIMARY KEY,
    tenant_id         INTEGER NOT NULL,
    series_id         VARCHAR(64) NOT NULL,
    event_type        VARCHAR(64) NOT NULL,
    chapter_seq       INTEGER NOT NULL DEFAULT 0,
    participants      TEXT NOT NULL DEFAULT '[]',
    location          VARCHAR(255) NOT NULL DEFAULT '',
    outcome           TEXT NOT NULL DEFAULT '',
    chronology        INTEGER NOT NULL DEFAULT 0,
    created_at        DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at        DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_manga_events_chapter
    ON manga_events (tenant_id, series_id, chapter_seq);

CREATE INDEX IF NOT EXISTS idx_manga_events_chronology
    ON manga_events (tenant_id, series_id, chronology);

CREATE TABLE IF NOT EXISTS manga_relations (
    id                  VARCHAR(64) PRIMARY KEY,
    tenant_id           INTEGER NOT NULL,
    series_id           VARCHAR(64) NOT NULL,
    source_character_id VARCHAR(64) NOT NULL,
    target_character_id VARCHAR(64) NOT NULL,
    relation_type       VARCHAR(64) NOT NULL,
    start_chapter_seq   INTEGER NOT NULL DEFAULT 0,
    end_chapter_seq     INTEGER,
    created_at          DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at          DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_manga_relations_pair
    ON manga_relations (tenant_id, series_id, source_character_id, target_character_id);

CREATE INDEX IF NOT EXISTS idx_manga_relations_chapters
    ON manga_relations (tenant_id, series_id, start_chapter_seq, end_chapter_seq);

CREATE TABLE IF NOT EXISTS manga_evidence (
    id                VARCHAR(64) PRIMARY KEY,
    tenant_id         INTEGER NOT NULL,
    series_id         VARCHAR(64) NOT NULL,
    claim             TEXT NOT NULL,
    evidence_text     TEXT NOT NULL,
    chapter_seq       INTEGER NOT NULL DEFAULT 0,
    page              INTEGER NOT NULL DEFAULT 0,
    source_type       VARCHAR(64) NOT NULL DEFAULT 'primary_manga',
    canon_tier        INTEGER NOT NULL DEFAULT 1,
    confidence        REAL NOT NULL DEFAULT 1.0,
    created_at        DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at        DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_manga_evidence_chapter
    ON manga_evidence (tenant_id, series_id, chapter_seq);

CREATE INDEX IF NOT EXISTS idx_manga_evidence_tier
    ON manga_evidence (tenant_id, series_id, canon_tier);
