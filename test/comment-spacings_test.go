package test

import (
	"testing"

	"github.com/deepsourcelabs/revive/lint"
	"github.com/deepsourcelabs/revive/rule"
)

func TestCommentSpacings(t *testing.T) {

	testRule(t, "comment-spacings", &rule.CommentSpacingsRule{}, &lint.RuleConfig{
		Arguments: []interface{}{"myOwnDirective"}},
	)
}
