package obfuscating

import (
	"go/ast"
	"go/token"
)

type NamedSymbols struct {
	Funcs    []*ast.FuncDecl
	Fields   []*ast.Field
	Comments []*ast.Comment
	Vals     []*ast.ValueSpec
}

type Visitor struct {
	fset *token.FileSet

	ns NamedSymbols
}

func NewVisitor(fs *token.FileSet) *Visitor {
	return &Visitor{
		fset: fs,
		ns: NamedSymbols{
			Funcs:    []*ast.FuncDecl{},
			Fields:   []*ast.Field{},
			Comments: []*ast.Comment{},
			Vals:     []*ast.ValueSpec{},
		},
	}
}

func (v *Visitor) NamedSymbols() NamedSymbols {
	return v.ns
}

func (v *Visitor) Visit(n ast.Node) ast.Visitor {
	if n == nil {
		return nil
	}

	switch t := n.(type) {
	case *ast.FuncDecl:
		v.ns.Funcs = append(v.ns.Funcs, t)
	case *ast.Field:
		v.ns.Fields = append(v.ns.Fields, t)
	case *ast.Comment:
		v.ns.Comments = append(v.ns.Comments, t)
	case *ast.ValueSpec:
		v.ns.Vals = append(v.ns.Vals, t)
	default:
		// For debugging right now.
	}

	return v
}
