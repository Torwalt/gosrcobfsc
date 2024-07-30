package gitignore

import "regexp"

type (
	literal = string
	ruleMap = map[literal]literalOrExpr
)

type literalOrExpr struct {
	lit  literal
	expr *regexp.Regexp
}

func (loe literalOrExpr) matchOrCompare(in string) bool {
	if loe.expr == nil {
		return in == loe.lit
	}

	return loe.expr.MatchString(in)
}
