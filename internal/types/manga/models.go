package manga

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

// StringList represents a list of strings serialized as JSON in DB.
type StringList []string

// Value implements driver.Valuer.
func (s StringList) Value() (driver.Value, error) {
	if s == nil {
		return "[]", nil
	}
	return json.Marshal(s)
}

// Scan implements sql.Scanner.
func (s *StringList) Scan(value interface{}) error {
	if value == nil {
		*s = StringList{}
		return nil
	}
	var bytes []byte
	switch v := value.(type) {
	case []byte:
		bytes = v
	case string:
		bytes = []byte(v)
	default:
		return fmt.Errorf("failed to scan StringList: unsupported type %T", value)
	}
	if len(bytes) == 0 {
		*s = StringList{}
		return nil
	}
	return json.Unmarshal(bytes, s)
}

// SpoilerScope defines boundaries for spoiler-filtered queries.
type SpoilerScope struct {
	SeriesID   string `json:"series_id"`
	MaxChapter int    `json:"max_chapter"` // 0 means unconstrained
}

// IsSpoiled returns true if chapterSeq exceeds the allowed boundary.
func (s SpoilerScope) IsSpoiled(chapterSeq int) bool {
	if s.MaxChapter <= 0 {
		return false
	}
	return chapterSeq > s.MaxChapter
}

// MangaSeries represents a manga title / work.
type MangaSeries struct {
	ID            string    `json:"id" gorm:"type:varchar(64);primaryKey"`
	TenantID      uint64    `json:"tenant_id" gorm:"index"`
	Title         string    `json:"title" gorm:"type:varchar(255);not null"`
	RomajiTitle   string    `json:"romaji_title" gorm:"type:varchar(255);default:''"`
	EnglishTitle  string    `json:"english_title" gorm:"type:varchar(255);default:''"`
	Author        string    `json:"author" gorm:"type:varchar(255);default:''"`
	Status        string    `json:"status" gorm:"type:varchar(64);default:'ongoing'"`
	TotalChapters int       `json:"total_chapters" gorm:"default:0"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (MangaSeries) TableName() string {
	return "manga_series"
}

// MangaArc represents a story arc inside a manga series.
type MangaArc struct {
	ID              string    `json:"id" gorm:"type:varchar(64);primaryKey"`
	TenantID        uint64    `json:"tenant_id" gorm:"index"`
	SeriesID        string    `json:"series_id" gorm:"type:varchar(64);index"`
	Name            string    `json:"name" gorm:"type:varchar(255);not null"`
	ArcSeq          int       `json:"arc_seq" gorm:"default:0"`
	StartChapterSeq int       `json:"start_chapter_seq" gorm:"default:0"`
	EndChapterSeq   int       `json:"end_chapter_seq" gorm:"default:0"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func (MangaArc) TableName() string {
	return "manga_arcs"
}

// MangaChapter represents a single publication chapter.
type MangaChapter struct {
	ID            string     `json:"id" gorm:"type:varchar(64);primaryKey"`
	TenantID      uint64     `json:"tenant_id" gorm:"index"`
	SeriesID      string     `json:"series_id" gorm:"type:varchar(64);uniqueIndex:uq_manga_chapters_seq,priority:2"`
	ArcID         string     `json:"arc_id" gorm:"type:varchar(64);default:''"`
	ChapterNumber string     `json:"chapter_number" gorm:"type:varchar(32);not null"`
	ChapterSeq    int        `json:"chapter_seq" gorm:"uniqueIndex:uq_manga_chapters_seq,priority:3;not null"`
	Title         string     `json:"title" gorm:"type:varchar(255);default:''"`
	ReleaseDate   *time.Time `json:"release_date,omitempty"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
}

func (MangaChapter) TableName() string {
	return "manga_chapters"
}

// MangaCharacter represents an entity / person in the story.
type MangaCharacter struct {
	ID                        string     `json:"id" gorm:"type:varchar(64);primaryKey"`
	TenantID                  uint64     `json:"tenant_id" gorm:"index"`
	SeriesID                  string     `json:"series_id" gorm:"type:varchar(64);index"`
	CanonicalName             string     `json:"canonical_name" gorm:"type:varchar(255);not null;index"`
	Aliases                   StringList `json:"aliases" gorm:"type:text"`
	Status                    string     `json:"status" gorm:"type:varchar(64);default:'alive'"`
	FirstAppearanceChapterSeq int        `json:"first_appearance_chapter_seq" gorm:"default:0;index"`
	CreatedAt                 time.Time  `json:"created_at"`
	UpdatedAt                 time.Time  `json:"updated_at"`
}

func (MangaCharacter) TableName() string {
	return "manga_characters"
}

// MangaEvent represents a plot event or encounter with temporal anchoring.
type MangaEvent struct {
	ID           string     `json:"id" gorm:"type:varchar(64);primaryKey"`
	TenantID     uint64     `json:"tenant_id" gorm:"index"`
	SeriesID     string     `json:"series_id" gorm:"type:varchar(64);index"`
	EventType    string     `json:"event_type" gorm:"type:varchar(64);not null"`
	ChapterSeq   int        `json:"chapter_seq" gorm:"default:0;index"`
	Participants StringList `json:"participants" gorm:"type:text"`
	Location     string     `json:"location" gorm:"type:varchar(255);default:''"`
	Outcome      string     `json:"outcome" gorm:"type:text;default:''"`
	Chronology   int        `json:"chronology" gorm:"default:0;index"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}

func (MangaEvent) TableName() string {
	return "manga_events"
}

// MangaRelation represents a directional or mutual relationship between characters.
type MangaRelation struct {
	ID                string    `json:"id" gorm:"type:varchar(64);primaryKey"`
	TenantID          uint64    `json:"tenant_id" gorm:"index"`
	SeriesID          string    `json:"series_id" gorm:"type:varchar(64);index"`
	SourceCharacterID string    `json:"source_character_id" gorm:"type:varchar(64);index"`
	TargetCharacterID string    `json:"target_character_id" gorm:"type:varchar(64);index"`
	RelationType      string    `json:"relation_type" gorm:"type:varchar(64);not null"`
	StartChapterSeq   int       `json:"start_chapter_seq" gorm:"default:0;index"`
	EndChapterSeq     *int      `json:"end_chapter_seq,omitempty"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

func (MangaRelation) TableName() string {
	return "manga_relations"
}

// MangaEvidence represents a cited claim grounded in a specific source and chapter.
type MangaEvidence struct {
	ID           string    `json:"id" gorm:"type:varchar(64);primaryKey"`
	TenantID     uint64    `json:"tenant_id" gorm:"index"`
	SeriesID     string    `json:"series_id" gorm:"type:varchar(64);index"`
	Claim        string    `json:"claim" gorm:"type:text;not null"`
	EvidenceText string    `json:"evidence_text" gorm:"type:text;not null"`
	ChapterSeq   int       `json:"chapter_seq" gorm:"default:0;index"`
	Page         int       `json:"page" gorm:"default:0"`
	SourceType   string    `json:"source_type" gorm:"type:varchar(64);default:'primary_manga'"`
	CanonTier    CanonTier `json:"canon_tier" gorm:"default:1;index"`
	Confidence   float32   `json:"confidence" gorm:"default:1.0"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (MangaEvidence) TableName() string {
	return "manga_evidence"
}
