package manga

import "strings"

// CanonTier represents the authoritative ranking of a manga evidence or source.
// Lower numeric value indicates higher authority (Tier 1 is most authoritative).
type CanonTier int

const (
	// CanonTierPrimaryManga represents primary canonical manga chapters written/drawn by author.
	CanonTierPrimaryManga CanonTier = 1
	// CanonTierOfficialDatabook represents official databooks, guidebooks, and Vivre Cards.
	CanonTierOfficialDatabook CanonTier = 2
	// CanonTierAuthorStatement represents author interviews, SBS, and direct creator comments.
	CanonTierAuthorStatement CanonTier = 3
	// CanonTierDerivativeAnime represents anime adaptations, movies, filler, and spin-offs.
	CanonTierDerivativeAnime CanonTier = 4
	// CanonTierFanAnalysis represents fan theories, community synthesis, and third-party commentary.
	CanonTierFanAnalysis CanonTier = 5
)

// String returns human-readable label for the CanonTier.
func (t CanonTier) String() string {
	switch t {
	case CanonTierPrimaryManga:
		return "primary_manga"
	case CanonTierOfficialDatabook:
		return "official_databook"
	case CanonTierAuthorStatement:
		return "author_statement"
	case CanonTierDerivativeAnime:
		return "derivative_anime"
	case CanonTierFanAnalysis:
		return "fan_analysis"
	default:
		return "unknown"
	}
}

// ParseCanonTier parses a string into CanonTier.
func ParseCanonTier(s string) CanonTier {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "primary", "primary_manga", "manga", "canon":
		return CanonTierPrimaryManga
	case "databook", "official_databook", "guidebook":
		return CanonTierOfficialDatabook
	case "author", "author_statement", "sbs", "interview":
		return CanonTierAuthorStatement
	case "anime", "derivative_anime", "movie", "spin_off":
		return CanonTierDerivativeAnime
	case "theory", "fan_analysis", "community", "analysis":
		return CanonTierFanAnalysis
	default:
		return CanonTierPrimaryManga
	}
}

// IsMoreAuthoritative returns true if t has higher authority than other.
func (t CanonTier) IsMoreAuthoritative(other CanonTier) bool {
	return t < other
}

// EvidenceConfidence represents the evidential grounding state.
type EvidenceConfidence string

const (
	ConfidenceConfirmed       EvidenceConfidence = "confirmed"
	ConfidenceStronglyImplied EvidenceConfidence = "strongly_implied"
	ConfidencePossible        EvidenceConfidence = "possible"
	ConfidenceSpeculative     EvidenceConfidence = "speculative"
	ConfidenceContradicted    EvidenceConfidence = "contradicted"
)
