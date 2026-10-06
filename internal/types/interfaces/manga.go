package interfaces

import (
	"context"

	"github.com/Tencent/WeKnora/internal/types/manga"
)

// MangaRepository defines data access methods for manga entities with spoiler boundaries and canon ranking.
type MangaRepository interface {
	// Series operations
	CreateSeries(ctx context.Context, s *manga.MangaSeries) error
	GetSeries(ctx context.Context, tenantID uint64, seriesID string) (*manga.MangaSeries, error)

	// Chapter operations
	CreateChapter(ctx context.Context, ch *manga.MangaChapter) error
	GetChapterBySeq(ctx context.Context, tenantID uint64, seriesID string, seq int) (*manga.MangaChapter, error)

	// Character operations (with spoiler boundary)
	CreateCharacter(ctx context.Context, c *manga.MangaCharacter) error
	GetCharacterByName(ctx context.Context, tenantID uint64, seriesID string, name string, maxChapter int) (*manga.MangaCharacter, error)
	ListCharacters(ctx context.Context, tenantID uint64, seriesID string, maxChapter int, limit, offset int) ([]*manga.MangaCharacter, error)

	// Timeline / Event operations
	CreateEvent(ctx context.Context, e *manga.MangaEvent) error
	GetTimeline(ctx context.Context, tenantID uint64, seriesID string, characterID string, maxChapter int, limit int) ([]*manga.MangaEvent, error)

	// Relationship operations
	CreateRelation(ctx context.Context, r *manga.MangaRelation) error
	GetCharacterRelations(ctx context.Context, tenantID uint64, seriesID string, characterID string, maxChapter int) ([]*manga.MangaRelation, error)

	// Evidence & Canon retrieval
	CreateEvidence(ctx context.Context, ev *manga.MangaEvidence) error
	SearchEvidence(ctx context.Context, tenantID uint64, seriesID string, query string, maxTier manga.CanonTier, maxChapter int, limit int) ([]*manga.MangaEvidence, error)
}
