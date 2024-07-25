package test

import (
	"testing"

	"github.com/DeepSourceCorp/revive/lint"
	"github.com/DeepSourceCorp/revive/rule"
)

// TestStructTag tests struct-tag rule
func TestStructTag(t *testing.T) {
	testRule(t, "struct-tag", &rule.StructTagRule{})
}

func TestStructTagWithUserOptions(t *testing.T) {
	testRule(t, "struct-tag-useroptions", &rule.StructTagRule{}, &lint.RuleConfig{
		Arguments: []any{"json,inline,outline", "bson,gnu"},
	})
}
