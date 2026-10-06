package repository

import (
	"context"
	"errors"
	"strings"

	"github.com/Tencent/WeKnora/internal/types/interfaces"
	"github.com/Tencent/WeKnora/internal/types/manga"
	"gorm.io/gorm"
)

// ErrNotFound indicates a requested manga entity was not found.
var ErrNotFound = errors.New("manga entity not found")

type mangaRepository struct {
	db *gorm.DB
}

// NewMangaRepository creates a new instance of MangaRepository.
func NewMangaRepository(db *gorm.DB) interfaces.MangaRepository {
	return &mangaRepository{db: db}
}

func (r *mangaRepository) CreateSeries(ctx context.Context, s *manga.MangaSeries) error {
	return r.db.WithContext(ctx).Create(s).Error
}

func (r *mangaRepository) GetSeries(ctx context.Context, tenantID uint64, seriesID string) (*manga.MangaSeries, error) {
	var s manga.MangaSeries
	err := r.db.WithContext(ctx).
		Where("tenant_id = ? AND id = ?", tenantID, seriesID).
		First(&s).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrNotFound
	}
	return &s, err
}

func (r *mangaRepository) CreateChapter(ctx context.Context, ch *manga.MangaChapter) error {
	return r.db.WithContext(ctx).Create(ch).Error
}

func (r *mangaRepository) GetChapterBySeq(ctx context.Context, tenantID uint64, seriesID string, seq int) (*manga.MangaChapter, error) {
	var ch manga.MangaChapter
	err := r.db.WithContext(ctx).
		Where("tenant_id = ? AND series_id = ? AND chapter_seq = ?", tenantID, seriesID, seq).
		First(&ch).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrNotFound
	}
	return &ch, err
}

func (r *mangaRepository) CreateCharacter(ctx context.Context, c *manga.MangaCharacter) error {
	return r.db.WithContext(ctx).Create(c).Error
}

func (r *mangaRepository) GetCharacterByName(ctx context.Context, tenantID uint64, seriesID string, name string, maxChapter int) (*manga.MangaCharacter, error) {
	cleanName := strings.TrimSpace(name)
	if cleanName == "" {
		return nil, ErrNotFound
	}

	query := r.db.WithContext(ctx).
		Where("tenant_id = ? AND series_id = ?", tenantID, seriesID).
		Where("(LOWER(canonical_name) = LOWER(?) OR LOWER(aliases) LIKE LOWER(?))", cleanName, "%\""+cleanName+"\"%")

	if maxChapter > 0 {
		query = query.Where("first_appearance_chapter_seq <= ?", maxChapter)
	}

	var c manga.MangaCharacter
	err := query.First(&c).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrNotFound
	}
	return &c, err
}

func (r *mangaRepository) ListCharacters(ctx context.Context, tenantID uint64, seriesID string, maxChapter int, limit, offset int) ([]*manga.MangaCharacter, error) {
	if limit <= 0 || limit > 100 {
		limit = 20
	}
	query := r.db.WithContext(ctx).
		Where("tenant_id = ? AND series_id = ?", tenantID, seriesID)

	if maxChapter > 0 {
		query = query.Where("first_appearance_chapter_seq <= ?", maxChapter)
	}

	var characters []*manga.MangaCharacter
	err := query.Order("first_appearance_chapter_seq ASC, canonical_name ASC").
		Limit(limit).
		Offset(offset).
		Find(&characters).Error
	return characters, err
}

func (r *mangaRepository) CreateEvent(ctx context.Context, e *manga.MangaEvent) error {
	return r.db.WithContext(ctx).Create(e).Error
}

func (r *mangaRepository) GetTimeline(ctx context.Context, tenantID uint64, seriesID string, characterID string, maxChapter int, limit int) ([]*manga.MangaEvent, error) {
	if limit <= 0 || limit > 100 {
		limit = 50
	}
	query := r.db.WithContext(ctx).
		Where("tenant_id = ? AND series_id = ?", tenantID, seriesID)

	if characterID != "" {
		query = query.Where("participants LIKE ?", "%\""+characterID+"\"%")
	}
	if maxChapter > 0 {
		query = query.Where("chapter_seq <= ?", maxChapter)
	}

	var events []*manga.MangaEvent
	err := query.Order("chapter_seq ASC, chronology ASC").
		Limit(limit).
		Find(&events).Error
	return events, err
}

func (r *mangaRepository) CreateRelation(ctx context.Context, rel *manga.MangaRelation) error {
	return r.db.WithContext(ctx).Create(rel).Error
}

func (r *mangaRepository) GetCharacterRelations(ctx context.Context, tenantID uint64, seriesID string, characterID string, maxChapter int) ([]*manga.MangaRelation, error) {
	query := r.db.WithContext(ctx).
		Where("tenant_id = ? AND series_id = ?", tenantID, seriesID).
		Where("(source_character_id = ? OR target_character_id = ?)", characterID, characterID)

	if maxChapter > 0 {
		query = query.Where("start_chapter_seq <= ?", maxChapter).
			Where("(end_chapter_seq IS NULL OR end_chapter_seq >= ?)", maxChapter)
	}

	var relations []*manga.MangaRelation
	err := query.Order("start_chapter_seq ASC").Find(&relations).Error
	return relations, err
}

func (r *mangaRepository) CreateEvidence(ctx context.Context, ev *manga.MangaEvidence) error {
	return r.db.WithContext(ctx).Create(ev).Error
}

func (r *mangaRepository) SearchEvidence(ctx context.Context, tenantID uint64, seriesID string, query string, maxTier manga.CanonTier, maxChapter int, limit int) ([]*manga.MangaEvidence, error) {
	if limit <= 0 || limit > 100 {
		limit = 20
	}
	q := r.db.WithContext(ctx).
		Where("tenant_id = ? AND series_id = ?", tenantID, seriesID)

	if strings.TrimSpace(query) != "" {
		pattern := "%" + strings.TrimSpace(query) + "%"
		q = q.Where("(claim LIKE ? OR evidence_text LIKE ?)", pattern, pattern)
	}
	if maxTier > 0 {
		q = q.Where("canon_tier <= ?", int(maxTier))
	}
	if maxChapter > 0 {
		q = q.Where("chapter_seq <= ?", maxChapter)
	}

	var evidenceList []*manga.MangaEvidence
	err := q.Order("canon_tier ASC, confidence DESC, chapter_seq DESC").
		Limit(limit).
		Find(&evidenceList).Error
	return evidenceList, err
}
