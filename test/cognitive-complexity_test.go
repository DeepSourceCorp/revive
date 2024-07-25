package test

import (
	"testing"

	"github.com/DeepSourceCorp/revive/lint"
	"github.com/DeepSourceCorp/revive/rule"
)

func TestCognitiveComplexity(t *testing.T) {
	testRule(t, "cognitive-complexity", &rule.CognitiveComplexityRule{}, &lint.RuleConfig{
		Arguments: []any{int64(0)},
	})
}
