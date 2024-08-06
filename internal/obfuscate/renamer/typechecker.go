package renamer

import (
	"go/ast"
	"go/types"
)

type TypeChecker struct {
	info *types.Info
}

func NewTypeChecker(i *types.Info) *TypeChecker {
	return &TypeChecker{
		info: i,
	}
}

func (tc *TypeChecker) Package(in *ast.Ident) *types.Package {
	obj, ok := tc.info.Uses[in]
	if !ok {
		return nil
	}

	return obj.Pkg()
}
