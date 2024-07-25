package test

import (
	"testing"

	"github.com/DeepSourceCorp/revive/lint"
	"github.com/DeepSourceCorp/revive/rule"
)

func TestDotImports(t *testing.T) {
	args := []any{map[string]any{
		"allowedPackages": []any{"errors", "context", "github.com/BurntSushi/toml"},
	}}

	testRule(t, "import-dot", &rule.DotImportsRule{}, &lint.RuleConfig{
		Arguments: args,
	})
}
