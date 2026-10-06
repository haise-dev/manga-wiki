package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/Tencent/WeKnora/internal/types"
	"github.com/Tencent/WeKnora/internal/types/interfaces"
	"github.com/Tencent/WeKnora/internal/types/manga"
	"github.com/Tencent/WeKnora/internal/utils"
)

func resolveTenantID(ctx context.Context) uint64 {
	if tid, ok := types.TenantIDFromContext(ctx); ok && tid > 0 {
		return tid
	}
	return 1
}

// ---------------------------------------------------------------------------
// 1. LookupCharacterTool
// ---------------------------------------------------------------------------

type LookupCharacterInput struct {
	SeriesID   string `json:"series_id" jsonschema:"ID of the manga series (e.g. 'one-piece')"`
	Name       string `json:"name" jsonschema:"Character name, alias, title, or nickname"`
	MaxChapter int    `json:"max_chapter,omitempty" jsonschema:"Optional: maximum chapter sequence for spoiler boundary (e.g. 500)"`
}

var lookupCharacterToolDef = BaseTool{
	name: ToolMangaLookupCharacter,
	description: `Lookup a manga character by name, alias, or title with spoiler-aware boundary.
If the character has not yet appeared or been introduced by max_chapter, returns a spoiler protection notice without revealing future identity or affiliations.`,
	schema: utils.GenerateSchema[LookupCharacterInput](),
}

type LookupCharacterTool struct {
	BaseTool
	repo interfaces.MangaRepository
}

func NewLookupCharacterTool(repo interfaces.MangaRepository) *LookupCharacterTool {
	return &LookupCharacterTool{
		BaseTool: lookupCharacterToolDef,
		repo:     repo,
	}
}

func (t *LookupCharacterTool) Execute(ctx context.Context, args json.RawMessage) (*types.ToolResult, error) {
	var input LookupCharacterInput
	if err := json.Unmarshal(args, &input); err != nil {
		return &types.ToolResult{Success: false, Error: "invalid arguments: " + err.Error()}, nil
	}
	if strings.TrimSpace(input.SeriesID) == "" {
		return &types.ToolResult{Success: false, Error: "series_id is required"}, nil
	}
	if strings.TrimSpace(input.Name) == "" {
		return &types.ToolResult{Success: false, Error: "character name is required"}, nil
	}

	tenantID := resolveTenantID(ctx)
	char, err := t.repo.GetCharacterByName(ctx, tenantID, input.SeriesID, input.Name, input.MaxChapter)
	if err != nil {
		if input.MaxChapter > 0 {
			return &types.ToolResult{
				Success: true,
				Output:  fmt.Sprintf("Character %q not found or has not appeared yet up to Chapter %d (spoiler protected).", input.Name, input.MaxChapter),
			}, nil
		}
		return &types.ToolResult{
			Success: true,
			Output:  fmt.Sprintf("Character %q not found in series %q.", input.Name, input.SeriesID),
		}, nil
	}

	out := fmt.Sprintf("### Character: %s\n- **ID**: %s\n- **Status**: %s\n- **First Appearance**: Chapter %d\n- **Aliases**: %s",
		char.CanonicalName, char.ID, char.Status, char.FirstAppearanceChapterSeq, strings.Join(char.Aliases, ", "))

	return &types.ToolResult{
		Success: true,
		Output:  out,
		Data: map[string]interface{}{
			"id":               char.ID,
			"canonical_name":   char.CanonicalName,
			"aliases":          char.Aliases,
			"status":           char.Status,
			"first_appearance": char.FirstAppearanceChapterSeq,
		},
	}, nil
}

// ---------------------------------------------------------------------------
// 2. GetTimelineTool
// ---------------------------------------------------------------------------

type GetTimelineInput struct {
	SeriesID    string `json:"series_id" jsonschema:"ID of the manga series (e.g. 'one-piece')"`
	CharacterID string `json:"character_id,omitempty" jsonschema:"Optional: filter timeline events involving this character ID"`
	MaxChapter  int    `json:"max_chapter,omitempty" jsonschema:"Optional: maximum chapter sequence for spoiler boundary"`
	Limit       int    `json:"limit,omitempty" jsonschema:"Maximum number of events to return (default 20, max 50)"`
}

var getTimelineToolDef = BaseTool{
	name: ToolMangaGetTimeline,
	description: `Retrieve chronological events in a manga series or character life up to an optional chapter spoiler boundary.
Respects story chronology and publication sequence.`,
	schema: utils.GenerateSchema[GetTimelineInput](),
}

type GetTimelineTool struct {
	BaseTool
	repo interfaces.MangaRepository
}

func NewGetTimelineTool(repo interfaces.MangaRepository) *GetTimelineTool {
	return &GetTimelineTool{
		BaseTool: getTimelineToolDef,
		repo:     repo,
	}
}

func (t *GetTimelineTool) Execute(ctx context.Context, args json.RawMessage) (*types.ToolResult, error) {
	var input GetTimelineInput
	if err := json.Unmarshal(args, &input); err != nil {
		return &types.ToolResult{Success: false, Error: "invalid arguments: " + err.Error()}, nil
	}
	if strings.TrimSpace(input.SeriesID) == "" {
		return &types.ToolResult{Success: false, Error: "series_id is required"}, nil
	}

	tenantID := resolveTenantID(ctx)
	events, err := t.repo.GetTimeline(ctx, tenantID, input.SeriesID, input.CharacterID, input.MaxChapter, input.Limit)
	if err != nil {
		return &types.ToolResult{Success: false, Error: "failed to fetch timeline: " + err.Error()}, nil
	}

	if len(events) == 0 {
		return &types.ToolResult{
			Success: true,
			Output:  "No events found matching the criteria.",
		}, nil
	}

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("### Timeline (Found %d events):\n", len(events)))
	for idx, ev := range events {
		sb.WriteString(fmt.Sprintf("%d. **[Ch. %d] %s**: %s (Location: %s, Participants: %s)\n",
			idx+1, ev.ChapterSeq, ev.EventType, ev.Outcome, ev.Location, strings.Join(ev.Participants, ", ")))
	}

	return &types.ToolResult{
		Success: true,
		Output:  sb.String(),
		Data: map[string]interface{}{
			"count": len(events),
		},
	}, nil
}

// ---------------------------------------------------------------------------
// 3. GetCharacterRelationshipsTool
// ---------------------------------------------------------------------------

type GetCharacterRelationshipsInput struct {
	SeriesID    string `json:"series_id" jsonschema:"ID of the manga series (e.g. 'one-piece')"`
	CharacterID string `json:"character_id" jsonschema:"ID of the character to inspect (e.g. 'char-luffy')"`
	MaxChapter  int    `json:"max_chapter,omitempty" jsonschema:"Optional: maximum chapter sequence for spoiler boundary"`
}

var getCharacterRelationshipsToolDef = BaseTool{
	name: ToolMangaGetRelationships,
	description: `Retrieve typed relationships between characters that are active/known up to a specific chapter point.
Avoids future relationship changes, betrayals, alliances, or deaths that occur after max_chapter.`,
	schema: utils.GenerateSchema[GetCharacterRelationshipsInput](),
}

type GetCharacterRelationshipsTool struct {
	BaseTool
	repo interfaces.MangaRepository
}

func NewGetCharacterRelationshipsTool(repo interfaces.MangaRepository) *GetCharacterRelationshipsTool {
	return &GetCharacterRelationshipsTool{
		BaseTool: getCharacterRelationshipsToolDef,
		repo:     repo,
	}
}

func (t *GetCharacterRelationshipsTool) Execute(ctx context.Context, args json.RawMessage) (*types.ToolResult, error) {
	var input GetCharacterRelationshipsInput
	if err := json.Unmarshal(args, &input); err != nil {
		return &types.ToolResult{Success: false, Error: "invalid arguments: " + err.Error()}, nil
	}
	if strings.TrimSpace(input.SeriesID) == "" {
		return &types.ToolResult{Success: false, Error: "series_id is required"}, nil
	}
	if strings.TrimSpace(input.CharacterID) == "" {
		return &types.ToolResult{Success: false, Error: "character_id is required"}, nil
	}

	tenantID := resolveTenantID(ctx)
	relations, err := t.repo.GetCharacterRelations(ctx, tenantID, input.SeriesID, input.CharacterID, input.MaxChapter)
	if err != nil {
		return &types.ToolResult{Success: false, Error: "failed to fetch relations: " + err.Error()}, nil
	}

	if len(relations) == 0 {
		return &types.ToolResult{
			Success: true,
			Output:  fmt.Sprintf("No active relationships found for character %s up to Chapter %d.", input.CharacterID, input.MaxChapter),
		}, nil
	}

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("### Relationships for %s (%d found):\n", input.CharacterID, len(relations)))
	for _, r := range relations {
		partner := r.TargetCharacterID
		if partner == input.CharacterID {
			partner = r.SourceCharacterID
		}
		endInfo := "ongoing"
		if r.EndChapterSeq != nil {
			endInfo = fmt.Sprintf("ended at Ch. %d", *r.EndChapterSeq)
		}
		sb.WriteString(fmt.Sprintf("- **%s** with `%s` (from Ch. %d, %s)\n",
			r.RelationType, partner, r.StartChapterSeq, endInfo))
	}

	return &types.ToolResult{
		Success: true,
		Output:  sb.String(),
		Data: map[string]interface{}{
			"count": len(relations),
		},
	}, nil
}

// ---------------------------------------------------------------------------
// 4. SearchMangaEvidenceTool
// ---------------------------------------------------------------------------

type SearchMangaEvidenceInput struct {
	SeriesID     string `json:"series_id" jsonschema:"ID of the manga series (e.g. 'one-piece')"`
	Query        string `json:"query" jsonschema:"Search query regarding lore, claim, or event"`
	MaxCanonTier int    `json:"max_canon_tier,omitempty" jsonschema:"Maximum canon tier: 1 (Primary Manga), 2 (Databook), 3 (Author Statement), 4 (Anime/Movie), 5 (Fan Analysis). Default is 2."`
	MaxChapter   int    `json:"max_chapter,omitempty" jsonschema:"Optional: maximum chapter sequence for spoiler boundary"`
	Limit        int    `json:"limit,omitempty" jsonschema:"Maximum items to retrieve (default 10)"`
}

var searchMangaEvidenceToolDef = BaseTool{
	name: ToolMangaSearchEvidence,
	description: `Search evidence-grounded claims about manga lore with canon tier authority ranking and spoiler protection.
Ranks primary canon higher than secondary databooks, anime adaptations, or fan theories.`,
	schema: utils.GenerateSchema[SearchMangaEvidenceInput](),
}

type SearchMangaEvidenceTool struct {
	BaseTool
	repo interfaces.MangaRepository
}

func NewSearchMangaEvidenceTool(repo interfaces.MangaRepository) *SearchMangaEvidenceTool {
	return &SearchMangaEvidenceTool{
		BaseTool: searchMangaEvidenceToolDef,
		repo:     repo,
	}
}

func (t *SearchMangaEvidenceTool) Execute(ctx context.Context, args json.RawMessage) (*types.ToolResult, error) {
	var input SearchMangaEvidenceInput
	if err := json.Unmarshal(args, &input); err != nil {
		return &types.ToolResult{Success: false, Error: "invalid arguments: " + err.Error()}, nil
	}
	if strings.TrimSpace(input.SeriesID) == "" {
		return &types.ToolResult{Success: false, Error: "series_id is required"}, nil
	}

	maxTier := manga.CanonTier(input.MaxCanonTier)
	if maxTier <= 0 {
		maxTier = manga.CanonTierOfficialDatabook
	}

	tenantID := resolveTenantID(ctx)
	evidenceList, err := t.repo.SearchEvidence(ctx, tenantID, input.SeriesID, input.Query, maxTier, input.MaxChapter, input.Limit)
	if err != nil {
		return &types.ToolResult{Success: false, Error: "failed to search evidence: " + err.Error()}, nil
	}

	if len(evidenceList) == 0 {
		return &types.ToolResult{
			Success: true,
			Output:  "No verified evidence found matching the query within the specified canon tier and chapter boundaries.",
		}, nil
	}

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("### Verified Evidence (%d items):\n", len(evidenceList)))
	for idx, ev := range evidenceList {
		sb.WriteString(fmt.Sprintf("%d. **[%s - Tier %d]** (Ch. %d, Confidence: %.1f)\n   - **Claim**: %s\n   - **Evidence**: %s\n",
			idx+1, ev.CanonTier.String(), ev.CanonTier, ev.ChapterSeq, ev.Confidence, ev.Claim, ev.EvidenceText))
	}

	return &types.ToolResult{
		Success: true,
		Output:  sb.String(),
		Data: map[string]interface{}{
			"count": len(evidenceList),
		},
	}, nil
}
