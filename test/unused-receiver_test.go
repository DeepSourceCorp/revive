package test

import (
	"testing"

	"github.com/DeepSourceCorp/revive/lint"
	"github.com/DeepSourceCorp/revive/rule"
)

func TestUnusedReceiver(t *testing.T) {
	testRule(t, "unused-receiver", &rule.UnusedReceiverRule{})
	testRule(t, "unused-receiver-custom-regex", &rule.UnusedReceiverRule{}, &lint.RuleConfig{Arguments: []any{
		map[string]any{"allowRegex": "^xxx"},
	}})
}
