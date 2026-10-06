package tools

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/Tencent/WeKnora/internal/types/interfaces"
	"github.com/Tencent/WeKnora/internal/types/manga"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type fakeMangaRepo struct {
	interfaces.MangaRepository
	characters []*manga.MangaCharacter
	events     []*manga.MangaEvent
	relations  []*manga.MangaRelation
	evidence   []*manga.MangaEvidence
}

func (f *fakeMangaRepo) GetCharacterByName(ctx context.Context, tenantID uint64, seriesID string, name string, maxChapter int) (*manga.MangaCharacter, error) {
	for _, c := range f.characters {
		if c.CanonicalName == name {
			if maxChapter > 0 && c.FirstAppearanceChapterSeq > maxChapter {
				return nil, errors.New("spoiler")
			}
			return c, nil
		}
		for _, alias := range c.Aliases {
			if alias == name {
				if maxChapter > 0 && c.FirstAppearanceChapterSeq > maxChapter {
					return nil, errors.New("spoiler")
				}
				return c, nil
			}
		}
	}
	return nil, errors.New("not found")
}

func (f *fakeMangaRepo) GetTimeline(ctx context.Context, tenantID uint64, seriesID string, characterID string, maxChapter int, limit int) ([]*manga.MangaEvent, error) {
	var matched []*manga.MangaEvent
	for _, e := range f.events {
		if maxChapter > 0 && e.ChapterSeq > maxChapter {
			continue
		}
		matched = append(matched, e)
	}
	return matched, nil
}

func (f *fakeMangaRepo) GetCharacterRelations(ctx context.Context, tenantID uint64, seriesID string, characterID string, maxChapter int) ([]*manga.MangaRelation, error) {
	var matched []*manga.MangaRelation
	for _, r := range f.relations {
		if maxChapter > 0 {
			if r.StartChapterSeq > maxChapter {
				continue
			}
			if r.EndChapterSeq != nil && *r.EndChapterSeq < maxChapter {
				continue
			}
		}
		matched = append(matched, r)
	}
	return matched, nil
}

func (f *fakeMangaRepo) SearchEvidence(ctx context.Context, tenantID uint64, seriesID string, query string, maxTier manga.CanonTier, maxChapter int, limit int) ([]*manga.MangaEvidence, error) {
	var matched []*manga.MangaEvidence
	for _, ev := range f.evidence {
		if maxTier > 0 && ev.CanonTier > maxTier {
			continue
		}
		if maxChapter > 0 && ev.ChapterSeq > maxChapter {
			continue
		}
		matched = append(matched, ev)
	}
	return matched, nil
}

func TestLookupCharacterTool_Execute(t *testing.T) {
	repo := &fakeMangaRepo{
		characters: []*manga.MangaCharacter{
			{
				ID:                        "char-luffy",
				CanonicalName:             "Monkey D. Luffy",
				Aliases:                   manga.StringList{"Luffy", "Straw Hat"},
				Status:                    "alive",
				FirstAppearanceChapterSeq: 1,
			},
			{
				ID:                        "char-yamato",
				CanonicalName:             "Yamato",
				Aliases:                   manga.StringList{"Oni Princess"},
				Status:                    "alive",
				FirstAppearanceChapterSeq: 971,
			},
		},
	}
	tool := NewLookupCharacterTool(repo)
	ctx := context.Background()

	// Found by alias
	res, err := tool.Execute(ctx, json.RawMessage(`{"series_id":"one-piece","name":"Straw Hat"}`))
	require.NoError(t, err)
	assert.True(t, res.Success)
	assert.Contains(t, res.Output, "Monkey D. Luffy")

	// Spoiler protected
	res, err = tool.Execute(ctx, json.RawMessage(`{"series_id":"one-piece","name":"Yamato","max_chapter":500}`))
	require.NoError(t, err)
	assert.True(t, res.Success)
	assert.Contains(t, res.Output, "spoiler protected")

	// Revealed when max_chapter >= 971
	res, err = tool.Execute(ctx, json.RawMessage(`{"series_id":"one-piece","name":"Yamato","max_chapter":1000}`))
	require.NoError(t, err)
	assert.True(t, res.Success)
	assert.Contains(t, res.Output, "Yamato")
}

func TestGetTimelineTool_Execute(t *testing.T) {
	repo := &fakeMangaRepo{
		events: []*manga.MangaEvent{
			{EventType: "Departure", ChapterSeq: 1, Outcome: "Luffy sets sail"},
			{EventType: "Raid", ChapterSeq: 980, Outcome: "Onigashima Raid begins"},
		},
	}
	tool := NewGetTimelineTool(repo)
	ctx := context.Background()

	// With max_chapter=100 -> only chapter 1 event
	res, err := tool.Execute(ctx, json.RawMessage(`{"series_id":"one-piece","max_chapter":100}`))
	require.NoError(t, err)
	assert.True(t, res.Success)
	assert.Contains(t, res.Output, "Departure")
	assert.NotContains(t, res.Output, "Onigashima")
}

func TestGetCharacterRelationshipsTool_Execute(t *testing.T) {
	endCh := 574
	repo := &fakeMangaRepo{
		relations: []*manga.MangaRelation{
			{
				SourceCharacterID: "char-luffy",
				TargetCharacterID: "char-ace",
				RelationType:      "Brother",
				StartChapterSeq:   158,
				EndChapterSeq:     &endCh,
			},
		},
	}
	tool := NewGetCharacterRelationshipsTool(repo)
	ctx := context.Background()

	// Active at chapter 300
	res, err := tool.Execute(ctx, json.RawMessage(`{"series_id":"one-piece","character_id":"char-luffy","max_chapter":300}`))
	require.NoError(t, err)
	assert.True(t, res.Success)
	assert.Contains(t, res.Output, "Brother")

	// Not active before chapter 158
	res, err = tool.Execute(ctx, json.RawMessage(`{"series_id":"one-piece","character_id":"char-luffy","max_chapter":100}`))
	require.NoError(t, err)
	assert.True(t, res.Success)
	assert.Contains(t, res.Output, "No active relationships found")
}

func TestSearchMangaEvidenceTool_Execute(t *testing.T) {
	repo := &fakeMangaRepo{
		evidence: []*manga.MangaEvidence{
			{
				Claim:        "Luffy fruit is Nika",
				EvidenceText: "Gorosei dialog",
				ChapterSeq:   1044,
				CanonTier:    manga.CanonTierPrimaryManga,
				Confidence:   1.0,
			},
			{
				Claim:        "Luffy fruit is rubber",
				EvidenceText: "Vivre Card databook",
				ChapterSeq:   100,
				CanonTier:    manga.CanonTierOfficialDatabook,
				Confidence:   0.9,
			},
			{
				Claim:        "Fan speculation",
				EvidenceText: "Reddit theory",
				ChapterSeq:   500,
				CanonTier:    manga.CanonTierFanAnalysis,
				Confidence:   0.3,
			},
		},
	}
	tool := NewSearchMangaEvidenceTool(repo)
	ctx := context.Background()

	// Query with default tier (Databook=2) and max_chapter=500
	// Excludes tier 5 (fan analysis) and chapter 1044 (spoiler)
	res, err := tool.Execute(ctx, json.RawMessage(`{"series_id":"one-piece","query":"fruit","max_chapter":500}`))
	require.NoError(t, err)
	assert.True(t, res.Success)
	assert.Contains(t, res.Output, "Vivre Card")
	assert.NotContains(t, res.Output, "Nika")
	assert.NotContains(t, res.Output, "Reddit")
}
