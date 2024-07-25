package test

import (
	"testing"

	"github.com/DeepSourceCorp/revive/lint"
	"github.com/DeepSourceCorp/revive/rule"
)

func TestCommentSpacings(t *testing.T) {
	testRule(t, "comment-spacings", &rule.CommentSpacingsRule{}, &lint.RuleConfig{
		Arguments: []any{"myOwnDirective:", "+optional"}},
	)
}
