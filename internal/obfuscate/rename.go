package obfuscate

import (
	"go/ast"

	"github.com/Torwalt/gosrcobfsc/internal/hasher"
)

const (
	mainFunc  = "main"
	errorType = "error"
)

/* TODO:
- create directories and files with hashed names
- hash module pkgs, but dont hash other pkgs (e.g. std lib, external pkgs)
*/

func rename(ns NamedSymbols) {
	for _, f := range ns.Funcs {

		renameFunc(f)

		for _, stmt := range f.Body.List {
			switch t := stmt.(type) {
			case *ast.AssignStmt:
				renameAssignStatement(t)
			}
		}
	}
}

func renameAssignStatement(in *ast.AssignStmt) {
	if in == nil {
		return
	}

	for _, l := range in.Lhs {
		renameExpr(l)
	}

	for _, r := range in.Rhs {
		renameExpr(r)
	}
}

func renameCallExpr(in *ast.CallExpr) {
	if in == nil {
		return
	}

	switch t := in.Fun.(type) {
	case *ast.SelectorExpr:
		renameSelectorExpr(t)
	}
	for _, a := range in.Args {
		renameExpr(a)
	}

}

func renameSelectorExpr(in *ast.SelectorExpr) {
	if in == nil {
		return
	}

	renameIdent(in.Sel)
	renameExpr(in.X)
}

func renameExpr(in ast.Expr) {
	if in == nil {
		return
	}

	switch t := in.(type) {
	case *ast.Ident:
		renameIdent(t)
	case *ast.CallExpr:
		renameCallExpr(t)
	case *ast.SelectorExpr:
		renameSelectorExpr(t)
	}
}

func renameIdent(in *ast.Ident) {
	if in == nil {
		return
	}

	in.Name = hasher.Hash(in.Name)
}

func renameFunc(in *ast.FuncDecl) {
	if in == nil {
		return
	}

	if in.Name.Name == mainFunc {
		return
	}
	renameIdent(in.Name)
	renameFieldList(in.Type.TypeParams)
	renameFieldList(in.Type.Params)
	renameFieldList(in.Type.Results)

}

func renameFieldList(in *ast.FieldList) {
	if in == nil {
		return
	}

	for _, f := range in.List {
		for _, n := range f.Names {
			renameIdent(n)
		}

		// We must not rename error return type.
		if !isErrorType(f.Type) {
			renameExpr(f.Type)
		}
	}
}

func isErrorType(in ast.Expr) bool {
	if in == nil {
		return false
	}

	switch t := in.(type) {
	case *ast.Ident:
		if t.Name == errorType {
			return true
		}
	}

	return false
}
