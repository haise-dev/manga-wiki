package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/Tencent/WeKnora/internal/agent/skills"
	"github.com/Tencent/WeKnora/internal/types"
	"github.com/Tencent/WeKnora/internal/utils"
)

// ReadFileTool exposes read operations over skill-scoped resources.
type ReadFileTool struct {
	BaseTool
	skills *skills.Manager
}

type ReadFileInput struct {
	Path     string `json:"path" jsonschema:"File address: skill://<name>/<relative-file>. Read a skill's SKILL.md before applying it."`
	Offset   int    `json:"offset,omitempty" jsonschema:"1-based line number; omit for the first page."`
	Limit    int    `json:"limit,omitempty" jsonschema:"Maximum lines to return; defaults to 2000."`
	MaxBytes int64  `json:"max_bytes,omitempty" jsonschema:"Maximum returned text bytes."`
}

func NewReadFileTool(manager *skills.Manager) *ReadFileTool {
	t := &ReadFileTool{
		BaseTool: BaseTool{name: ToolReadFile, schema: utils.GenerateSchema[ReadFileInput]()},
		skills:   manager,
	}
	t.description = "Skill resources: skill://<name>/SKILL.md loads skill instructions and file list; skill://<name>/<relative-file> reads a bundled resource."
	return t
}

func (t *ReadFileTool) WithSkills(manager *skills.Manager, _ bool) *ReadFileTool {
	t.skills = manager
	return t
}

func (t *ReadFileTool) Execute(ctx context.Context, args json.RawMessage) (*types.ToolResult, error) {
	var input ReadFileInput
	if err := json.Unmarshal(args, &input); err != nil {
		return &types.ToolResult{Success: false, Error: fmt.Sprintf("invalid read_file arguments: %v", err)}, nil
	}
	input.Path = strings.TrimSpace(input.Path)
	if input.Path == "" {
		return &types.ToolResult{Success: false, Error: "path is required; use skill://<name>/<relative-file>"}, nil
	}
	if strings.HasPrefix(input.Path, "skill://") {
		return t.readSkillResource(ctx, input), nil
	}
	return &types.ToolResult{Success: false, Error: "path is outside scope: file access is limited to listed skill:// resources"}, nil
}

func (t *ReadFileTool) readSkillResource(ctx context.Context, input ReadFileInput) *types.ToolResult {
	fail := func(err error) *types.ToolResult { return &types.ToolResult{Success: false, Error: err.Error()} }
	if t.skills == nil || !t.skills.IsEnabled() {
		return fail(fmt.Errorf("skills are not enabled for this reader"))
	}
	name, rel, ok := strings.Cut(strings.TrimPrefix(input.Path, "skill://"), "/")
	if !ok || name == "" || rel == "" {
		return fail(fmt.Errorf("use skill://<name>/SKILL.md or skill://<name>/<relative-file>"))
	}
	for _, segment := range strings.Split(rel, "/") {
		if segment == ".." || segment == "." || segment == "" {
			return fail(fmt.Errorf("skill resource must have a canonical relative file path without traversal"))
		}
	}
	if strings.ContainsAny(name+rel, "\\\x00") {
		return fail(fmt.Errorf("invalid skill resource path"))
	}
	listed := false
	for _, metadata := range t.skills.GetAllMetadata() {
		if metadata != nil && metadata.Name == name {
			listed = true
			break
		}
	}
	if !listed {
		return fail(fmt.Errorf("skill %q is not available to this agent", name))
	}

	var content string
	if rel == skills.SkillFileName {
		skill, err := t.skills.LoadSkill(ctx, name)
		if err != nil {
			return fail(err)
		}
		var b strings.Builder
		fmt.Fprintf(&b, "# %s\n\n%s\n\n", skill.Name, skill.Description)
		b.WriteString(skill.Instructions)
		files, err := t.skills.ListSkillFiles(ctx, name)
		if err == nil && len(files) > 0 {
			fmt.Fprintf(&b, "\n\n## Bundled files\nRead these relative paths under skill://%s/:\n", name)
			for _, f := range files {
				fmt.Fprintf(&b, "- %s\n", f)
			}
		}
		content = b.String()
	} else {
		var err error
		content, err = t.skills.ReadSkillFile(ctx, name, rel)
		if err != nil {
			return fail(err)
		}
	}

	return &types.ToolResult{
		Success: true,
		Output:  content,
		Data: map[string]interface{}{
			"skill_name": name,
			"file_path":  rel,
			"path":       input.Path,
		},
	}
}
