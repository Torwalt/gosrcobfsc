package obfuscate

import (
	"go/ast"
)

// NamedSymbols represents all the Nodes that comprise a package.
type NamedSymbols struct {
	Package  *ast.Package
	Files    []*ast.File
	Funcs    []*ast.FuncDecl
	Fields   []*ast.Field
	Comments []*ast.Comment
	Vals     []*ast.ValueSpec
	Imports  []*ast.ImportSpec
	Types    []*ast.TypeSpec
}

type Visitor struct {
	ns NamedSymbols
}

func NewVisitor() *Visitor {
	return &Visitor{
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
	case *ast.Package:
		if v.ns.Package != nil {
			panic("There should be only one package!")
		}
		v.ns.Package = t
	case *ast.File:
		v.ns.Files = append(v.ns.Files, t)
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
		// For debugging right now.
	}

	return v
}
