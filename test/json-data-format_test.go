package test

import (
	"testing"

	"github.com/DeepSourceCorp/revive/lint"
	"github.com/DeepSourceCorp/revive/rule"
)

func TestJsonDataFormat(t *testing.T) {
	testRule(t, "json-data-format-atomic", &rule.AtomicRule{})

}
func TestJsonDataFormatVarNaming(t *testing.T) {
	testRule(t, "json-data-format-var-naming", &rule.VarNamingRule{}, &lint.RuleConfig{})

}
