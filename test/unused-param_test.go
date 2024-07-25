package test

import (
	"testing"

	"github.com/DeepSourceCorp/revive/lint"
	"github.com/DeepSourceCorp/revive/rule"
)

func TestUnusedParam(t *testing.T) {
	testRule(t, "unused-param", &rule.UnusedParamRule{})
	testRule(t, "unused-param-custom-regex", &rule.UnusedParamRule{}, &lint.RuleConfig{Arguments: []any{
		map[string]any{"allowRegex": "^xxx"},
	}})
}

func BenchmarkUnusedParam(b *testing.B) {
	var t *testing.T
	for i := 0; i <= b.N; i++ {
		testRule(t, "unused-param", &rule.UnusedParamRule{})
	}
}
