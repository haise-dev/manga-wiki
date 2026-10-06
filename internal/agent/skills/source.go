package skills

// SkillSource is where one set of skills comes from. Two implementations
// exist: host skill directories used in tests (*Loader) and the projection of
// skills an administrator installed into a workspace sandbox config's snapshot
// image (*TenantSkillSource).
//
// The five methods are the Progressive Disclosure levels the agent asks for:
// metadata for the system prompt, the SKILL.md body, individual resource
// files, the file listing, and the directory a script runs from.
type SkillSource interface {
	DiscoverSkills() ([]*SkillMetadata, error)
	LoadSkillInstructions(name string) (*Skill, error)
	LoadSkillFile(name, relativePath string) (*SkillFile, error)
	ListSkillFiles(name string) ([]string, error)
	GetSkillBasePath(name string) (string, error)
}

var _ SkillSource = (*Loader)(nil)

