package repository

import (
	"context"
	"testing"
	"time"

	"github.com/Tencent/WeKnora/internal/types/manga"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupMangaTestDB(t *testing.T) *gorm.DB {
	t.Helper()
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	require.NoError(t, err)
	require.NoError(t, db.AutoMigrate(
		&manga.MangaSeries{},
		&manga.MangaArc{},
		&manga.MangaChapter{},
		&manga.MangaCharacter{},
		&manga.MangaEvent{},
		&manga.MangaRelation{},
		&manga.MangaEvidence{},
	))
	return db
}

func TestMangaRepository_SeriesAndChapters(t *testing.T) {
	db := setupMangaTestDB(t)
	repo := NewMangaRepository(db)
	ctx := context.Background()

	series := &manga.MangaSeries{
		ID:           "one-piece",
		TenantID:     1,
		Title:        "One Piece",
		RomajiTitle:  "Wan Pīsu",
		EnglishTitle: "One Piece",
		Author:       "Eiichiro Oda",
		Status:       "ongoing",
	}
	err := repo.CreateSeries(ctx, series)
	require.NoError(t, err)

	fetched, err := repo.GetSeries(ctx, 1, "one-piece")
	require.NoError(t, err)
	assert.Equal(t, "One Piece", fetched.Title)
	assert.Equal(t, "Eiichiro Oda", fetched.Author)

	chapter := &manga.MangaChapter{
		ID:            "ch-1",
		TenantID:      1,
		SeriesID:      "one-piece",
		ChapterNumber: "1",
		ChapterSeq:    1,
		Title:         "Romance Dawn",
	}
	err = repo.CreateChapter(ctx, chapter)
	require.NoError(t, err)

	fetchedCh, err := repo.GetChapterBySeq(ctx, 1, "one-piece", 1)
	require.NoError(t, err)
	assert.Equal(t, "Romance Dawn", fetchedCh.Title)
	assert.Equal(t, 1, fetchedCh.ChapterSeq)
}

func TestMangaRepository_CharacterSpoilerBoundary(t *testing.T) {
	db := setupMangaTestDB(t)
	repo := NewMangaRepository(db)
	ctx := context.Background()

	luffy := &manga.MangaCharacter{
		ID:                        "char-luffy",
		TenantID:                  1,
		SeriesID:                  "one-piece",
		CanonicalName:             "Monkey D. Luffy",
		Aliases:                   manga.StringList{"Luffy", "Straw Hat", "Lucy"},
		FirstAppearanceChapterSeq: 1,
	}
	yamato := &manga.MangaCharacter{
		ID:                        "char-yamato",
		TenantID:                  1,
		SeriesID:                  "one-piece",
		CanonicalName:             "Yamato",
		Aliases:                   manga.StringList{"Oni Princess"},
		FirstAppearanceChapterSeq: 971,
	}
	require.NoError(t, repo.CreateCharacter(ctx, luffy))
	require.NoError(t, repo.CreateCharacter(ctx, yamato))

	// Search Luffy by alias "Lucy" with max_chapter=100 -> found
	foundLuffy, err := repo.GetCharacterByName(ctx, 1, "one-piece", "Lucy", 100)
	require.NoError(t, err)
	assert.Equal(t, "Monkey D. Luffy", foundLuffy.CanonicalName)

	// Search Yamato with max_chapter=500 -> should be redacted (not found)
	_, err = repo.GetCharacterByName(ctx, 1, "one-piece", "Yamato", 500)
	assert.ErrorIs(t, err, ErrNotFound)

	// Search Yamato with max_chapter=1000 -> found
	foundYamato, err := repo.GetCharacterByName(ctx, 1, "one-piece", "Yamato", 1000)
	require.NoError(t, err)
	assert.Equal(t, "Yamato", foundYamato.CanonicalName)

	// ListCharacters with max_chapter=100 -> only Luffy
	chars, err := repo.ListCharacters(ctx, 1, "one-piece", 100, 10, 0)
	require.NoError(t, err)
	assert.Len(t, chars, 1)
	assert.Equal(t, "Monkey D. Luffy", chars[0].CanonicalName)
}

func TestMangaRepository_TimelineEvents(t *testing.T) {
	db := setupMangaTestDB(t)
	repo := NewMangaRepository(db)
	ctx := context.Background()

	events := []*manga.MangaEvent{
		{
			ID:           "ev-1",
			TenantID:     1,
			SeriesID:     "one-piece",
			EventType:    "Departure",
			ChapterSeq:   1,
			Participants: manga.StringList{"char-luffy", "char-shanks"},
			Outcome:      "Luffy sets sail and receives straw hat",
			Chronology:   1,
		},
		{
			ID:           "ev-2",
			TenantID:     1,
			SeriesID:     "one-piece",
			EventType:    "Execution",
			ChapterSeq:   0, // Prologue/flashback
			Participants: manga.StringList{"char-roger"},
			Outcome:      "Roger initiates the Great Pirate Era",
			Chronology:   0,
		},
		{
			ID:           "ev-3",
			TenantID:     1,
			SeriesID:     "one-piece",
			EventType:    "War",
			ChapterSeq:   550,
			Participants: manga.StringList{"char-luffy", "char-ace", "char-whitebeard"},
			Outcome:      "Marineford War begins",
			Chronology:   10,
		},
	}
	for _, ev := range events {
		require.NoError(t, repo.CreateEvent(ctx, ev))
	}

	// Timeline up to chapter 100 for Luffy
	tl, err := repo.GetTimeline(ctx, 1, "one-piece", "char-luffy", 100, 10)
	require.NoError(t, err)
	require.Len(t, tl, 1)
	assert.Equal(t, "Departure", tl[0].EventType)

	// Timeline up to chapter 600 for Luffy -> includes War
	tl2, err := repo.GetTimeline(ctx, 1, "one-piece", "char-luffy", 600, 10)
	require.NoError(t, err)
	require.Len(t, tl2, 2)
}

func TestMangaRepository_RelationsIntervals(t *testing.T) {
	db := setupMangaTestDB(t)
	repo := NewMangaRepository(db)
	ctx := context.Background()

	end574 := 574
	rel1 := &manga.MangaRelation{
		ID:                "rel-1",
		TenantID:          1,
		SeriesID:          "one-piece",
		SourceCharacterID: "char-luffy",
		TargetCharacterID: "char-ace",
		RelationType:      "Brother",
		StartChapterSeq:   158,
		EndChapterSeq:     &end574, // Ace dies in 574
	}
	require.NoError(t, repo.CreateRelation(ctx, rel1))

	// At chapter 300: relationship is active
	activeRels, err := repo.GetCharacterRelations(ctx, 1, "one-piece", "char-luffy", 300)
	require.NoError(t, err)
	assert.Len(t, activeRels, 1)

	// Before chapter 158 (e.g. 100): Ace not introduced/not brothers yet
	earlyRels, err := repo.GetCharacterRelations(ctx, 1, "one-piece", "char-luffy", 100)
	require.NoError(t, err)
	assert.Empty(t, earlyRels)

	// Unconstrained (max_chapter=0)
	allRels, err := repo.GetCharacterRelations(ctx, 1, "one-piece", "char-luffy", 0)
	require.NoError(t, err)
	assert.Len(t, allRels, 1)
}

func TestMangaRepository_EvidenceSearchAndCanonHierarchy(t *testing.T) {
	db := setupMangaTestDB(t)
	repo := NewMangaRepository(db)
	ctx := context.Background()

	ev1 := &manga.MangaEvidence{
		ID:           "ev-fan",
		TenantID:     1,
		SeriesID:     "one-piece",
		Claim:        "Luffy devil fruit is Joyboy",
		EvidenceText: "Fan forum speculation from 2015",
		ChapterSeq:   700,
		CanonTier:    manga.CanonTierFanAnalysis,
		Confidence:   0.4,
		CreatedAt:    time.Now(),
	}
	ev2 := &manga.MangaEvidence{
		ID:           "ev-databook",
		TenantID:     1,
		SeriesID:     "one-piece",
		Claim:        "Luffy rubber body properties",
		EvidenceText: "Vivre Card Databook Vol 1",
		ChapterSeq:   100,
		CanonTier:    manga.CanonTierOfficialDatabook,
		Confidence:   0.95,
		CreatedAt:    time.Now(),
	}
	ev3 := &manga.MangaEvidence{
		ID:           "ev-manga",
		TenantID:     1,
		SeriesID:     "one-piece",
		Claim:        "Luffy fruit true name is Nika",
		EvidenceText: "Gorosei revelation in chapter 1044",
		ChapterSeq:   1044,
		CanonTier:    manga.CanonTierPrimaryManga,
		Confidence:   1.0,
		CreatedAt:    time.Now(),
	}

	require.NoError(t, repo.CreateEvidence(ctx, ev1))
	require.NoError(t, repo.CreateEvidence(ctx, ev2))
	require.NoError(t, repo.CreateEvidence(ctx, ev3))

	// Search with maxTier = Databook (Tier 2) and maxChapter = 500
	// ev1 (Fan tier 5) -> excluded by tier
	// ev3 (ch 1044) -> excluded by chapter
	// ev2 (ch 100, tier 2) -> included
	results, err := repo.SearchEvidence(ctx, 1, "one-piece", "Luffy", manga.CanonTierOfficialDatabook, 500, 10)
	require.NoError(t, err)
	require.Len(t, results, 1)
	assert.Equal(t, "ev-databook", results[0].ID)

	// Search up to chapter 1100 with all tiers -> Tier 1 (manga) ranked before Tier 2 before Tier 5
	allResults, err := repo.SearchEvidence(ctx, 1, "one-piece", "Luffy", manga.CanonTierFanAnalysis, 1100, 10)
	require.NoError(t, err)
	require.Len(t, allResults, 3)
	assert.Equal(t, manga.CanonTierPrimaryManga, allResults[0].CanonTier)
	assert.Equal(t, manga.CanonTierOfficialDatabook, allResults[1].CanonTier)
	assert.Equal(t, manga.CanonTierFanAnalysis, allResults[2].CanonTier)
}
