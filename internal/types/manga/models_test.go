package manga

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCanonTierComparison(t *testing.T) {
	assert.True(t, CanonTierPrimaryManga.IsMoreAuthoritative(CanonTierOfficialDatabook))
	assert.True(t, CanonTierOfficialDatabook.IsMoreAuthoritative(CanonTierDerivativeAnime))
	assert.False(t, CanonTierFanAnalysis.IsMoreAuthoritative(CanonTierPrimaryManga))
	assert.Equal(t, "primary_manga", CanonTierPrimaryManga.String())
	assert.Equal(t, CanonTierAuthorStatement, ParseCanonTier("author_statement"))
	assert.Equal(t, CanonTierDerivativeAnime, ParseCanonTier("anime"))
}

func TestStringListScanValue(t *testing.T) {
	list := StringList{"Luffy", "Straw Hat", "Lucy"}
	val, err := list.Value()
	require.NoError(t, err)

	var scanned StringList
	err = scanned.Scan(val)
	require.NoError(t, err)
	assert.Equal(t, list, scanned)

	// Test nil scan
	var nilScanned StringList
	err = nilScanned.Scan(nil)
	require.NoError(t, err)
	assert.Empty(t, nilScanned)
}

func TestSpoilerScope(t *testing.T) {
	scope := SpoilerScope{SeriesID: "one-piece", MaxChapter: 100}
	assert.False(t, scope.IsSpoiled(50))
	assert.False(t, scope.IsSpoiled(100))
	assert.True(t, scope.IsSpoiled(101))

	unconstrained := SpoilerScope{SeriesID: "one-piece", MaxChapter: 0}
	assert.False(t, unconstrained.IsSpoiled(1050))
}
