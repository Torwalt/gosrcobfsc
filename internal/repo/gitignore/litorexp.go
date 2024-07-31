package gitignore

import (
	"regexp"
	"strings"
)

type (
	literal = string
	ruleMap = map[literal]literalOrExpr
)

type literalOrExpr struct {
	lit  literal
	expr *regexp.Regexp
}

func (loe literalOrExpr) matchOrCompare(in string) bool {
	ok := loe.actualMatchOrCompare(in)
	if ok {
		return true
	}

	// Kinda hacky but meh.
	return strings.HasPrefix(in, loe.lit)
}

func (loe literalOrExpr) actualMatchOrCompare(in string) bool {
	if in == loe.lit {
		return true
	}

	if loe.expr == nil {
		return false
	}

	return loe.expr.MatchString(in)
}
