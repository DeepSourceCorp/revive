package test

import (
	"testing"

	"github.com/DeepSourceCorp/revive/lint"
	"github.com/DeepSourceCorp/revive/rule"
)

func TestErrorStringsWithCustomFunctions(t *testing.T) {
	args := []interface{}{"pkgErrors.Wrap"}
	testRule(t, "error-strings-with-custom-functions", &rule.ErrorStringsRule{}, &lint.RuleConfig{
		Arguments: args,
	})
}
