package agent

import (
	"github.com/Tencent/WeKnora/internal/agent/skills"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestToolGuidanceUsesActualCapabilities(t *testing.T) {
	metadata := []*skills.SkillMetadata{{Name: "demo", Description: "demo skill"}}
	text := formatSkillsMetadata(metadata, true)
	require.Contains(t, text, "read_file")
	require.NotContains(t, text, "execute_skill_script")
	require.NotContains(t, text, "MANDATORY")
	require.Empty(t, formatToolGuidance(nil))
}
