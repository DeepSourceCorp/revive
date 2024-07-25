package test

import (
	"testing"

	"github.com/DeepSourceCorp/revive/lint"
	"github.com/DeepSourceCorp/revive/rule"
)

func TestCyclomatic(t *testing.T) {
	testRule(t, "cyclomatic", &rule.CyclomaticRule{}, &lint.RuleConfig{
		Arguments: []any{int64(1)},
	})
	testRule(t, "cyclomatic-2", &rule.CyclomaticRule{}, &lint.RuleConfig{
		Arguments: []any{int64(3)},
	})
}
