package rule

import (
	"fmt"
	"go/ast"
	"go/token"

	"github.com/DeepSourceCorp/revive/lint"
)

func findBlockStmtRef(block *ast.BlockStmt, obj *ast.Object) bool {
	var found bool

	fn := func(node ast.Node) bool {
		id, ok := node.(*ast.Ident)
		if !ok {
			return true
		}

		if obj != id.Obj {
			return true
		}
		found = true
		return false
	}

	// Only recurse the block's tree and nothing else. Try to find
	// idenitifers that refer to the same obj.
	ast.Inspect(block, fn)

	return found
}

// SuperfluousElseRule lints given else constructs.
type SuperfluousElseRule struct{}

// Apply applies the rule to given file.
func (*SuperfluousElseRule) Apply(file *lint.File, _ lint.Arguments) []lint.Failure {
	var failures []lint.Failure
	onFailure := func(failure lint.Failure) {
		failures = append(failures, failure)
	}

	branchingFunctions := map[string]map[string]bool{
		"os": {"Exit": true},
		"log": {
			"Fatal":   true,
			"Fatalf":  true,
			"Fatalln": true,
			"Panic":   true,
			"Panicf":  true,
			"Panicln": true,
		},
	}

	w := lintSuperfluousElse{make(map[*ast.IfStmt]bool), onFailure, branchingFunctions}
	ast.Walk(w, file.AST)
	return failures
}

// Name returns the rule name.
func (*SuperfluousElseRule) Name() string {
	return "superfluous-else"
}

type lintSuperfluousElse struct {
	ignore             map[*ast.IfStmt]bool
	onFailure          func(lint.Failure)
	branchingFunctions map[string]map[string]bool
}

func (w lintSuperfluousElse) Visit(node ast.Node) ast.Visitor {
	ifStmt, ok := node.(*ast.IfStmt)
	if !ok || ifStmt.Else == nil {
		return w
	}
	if w.ignore[ifStmt] {
		if elseif, ok := ifStmt.Else.(*ast.IfStmt); ok {
			w.ignore[elseif] = true
		}
		return w
	}
	if elseif, ok := ifStmt.Else.(*ast.IfStmt); ok {
		w.ignore[elseif] = true
		return w
	}

	elseBlock, ok := ifStmt.Else.(*ast.BlockStmt)
	if !ok {
		// only care about elses without conditions
		return w
	}
	if len(ifStmt.Body.List) == 0 {
		return w
	}
	shortDecl := false // does the if statement have a ":=" initialization statement?
	if ifStmt.Init != nil {
		if as, ok := ifStmt.Init.(*ast.AssignStmt); ok && as.Tok == token.DEFINE {
			shortDecl = true
		}
	}
	extra := ""
	if shortDecl {
		extra = " (move short variable declaration to its own line if necessary)"
	}

	assignStmt, ok := ifStmt.Init.(*ast.AssignStmt)
	if ok && assignStmt.Tok == token.DEFINE { // only ":=" assignment are considered
		for _, lhs := range assignStmt.Lhs {
			ident, ok := lhs.(*ast.Ident)
			if !ok {
				continue
			}
			obj := ident.Obj
			if obj == nil {
				continue
			}
			if findBlockStmtRef(elseBlock, obj) {
				return w
			}
		}
	}

	lastStmt := ifStmt.Body.List[len(ifStmt.Body.List)-1]
	switch stmt := lastStmt.(type) {
	case *ast.BranchStmt:
		tok := stmt.Tok.String()
		if tok != "fallthrough" {
			w.onFailure(newFailure(ifStmt.Else, "if block ends with a "+tok+" statement, so drop this else and outdent its block"+extra))
		}
	case *ast.ExprStmt:
		if ce, ok := stmt.X.(*ast.CallExpr); ok { // it's a function call
			if fc, ok := ce.Fun.(*ast.SelectorExpr); ok {
				if id, ok := fc.X.(*ast.Ident); ok {
					fn := fc.Sel.Name
					pkg := id.Name
					if w.branchingFunctions[pkg][fn] { // it's a call to a branching function
						w.onFailure(
							newFailure(ifStmt.Else, fmt.Sprintf("if block ends with call to %s.%s function, so drop this else and outdent its block%s", pkg, fn, extra)))
					}
				}
			}
		}
	}

	return w
}

func newFailure(node ast.Node, msg string) lint.Failure {
	return lint.Failure{
		Confidence: 1,
		Node:       node,
		Category:   "indent",
		Failure:    msg,
	}
}
