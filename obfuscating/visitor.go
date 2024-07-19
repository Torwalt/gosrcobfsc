package obfuscating

import (
	"go/ast"
	"go/token"
)

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
			Imports:  []*ast.ImportSpec{},
			Types:    []*ast.TypeSpec{},
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
	case *ast.ImportSpec:
		v.ns.Imports = append(v.ns.Imports, t)
	case *ast.TypeSpec:
		v.ns.Types = append(v.ns.Types, t)
	default:
		print("asd")
		// For debugging right now.
	}

	return v
}
