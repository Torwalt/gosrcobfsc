package obfuscate

import (
	"go/ast"
)

// NamedSymbols represents all the Nodes that comprise a package.
type NamedSymbols struct {
	Package *ast.Package
	Files   []*ast.File
}

type Visitor struct {
	ns NamedSymbols
}

func NewVisitor() *Visitor {
	return &Visitor{
		ns: NamedSymbols{
			Package: nil,
			Files:   []*ast.File{},
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
	}

	return v
}
