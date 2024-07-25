package test

import (
	"testing"

	"github.com/DeepSourceCorp/revive/rule"
)

func TestImportShadowing(t *testing.T) {
	testRule(t, "import-shadowing", &rule.ImportShadowingRule{})
}
