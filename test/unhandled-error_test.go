package test

import (
	"testing"

	"github.com/DeepSourceCorp/revive/lint"
	"github.com/DeepSourceCorp/revive/rule"
)

func TestUnhandledError(t *testing.T) {
	testRule(t, "unhandled-error", &rule.UnhandledErrorRule{})
}

func TestUnhandledErrorWithIgnoreList(t *testing.T) {
	args := []any{
		`unhandledError1`,
		`fmt\.Print`,
		`os\.(Create|WriteFile|Chmod)`,
		`net\..*`,
		`bytes\.Buffer\.Write`,
		`fixtures\.unhandledErrorStruct2\.reterr`,
	}

	testRule(t, "unhandled-error-w-ignorelist", &rule.UnhandledErrorRule{}, &lint.RuleConfig{Arguments: args})
}
