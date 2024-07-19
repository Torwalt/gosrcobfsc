package obfuscating

import "go/ast"

type NamedSymbols struct {
	Funcs    []*ast.FuncDecl
	Fields   []*ast.Field
	Comments []*ast.Comment
	Vals     []*ast.ValueSpec
	Imports  []*ast.ImportSpec
	Types    []*ast.TypeSpec
}
