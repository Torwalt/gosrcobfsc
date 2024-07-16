package obfuscating

import (
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

/*
Goal: Collect all "user named" tokens, rename them and output same structure.

https://eli.thegreenplace.net/2021/rewriting-go-source-code-with-ast-tooling/
*/

func Obfuscate(content string) (string, error) {
	fset := token.NewFileSet()

	f, err := parser.ParseFile(fset, "", content, 0)
	if err != nil {
		return "", err
	}

	v := newVisitor(fset)
	ast.Walk(v, f)

	return strings.Join(v.funcs, ","), nil
}

type Visitor struct {
	fset  *token.FileSet
	funcs []string
}

func newVisitor(fs *token.FileSet) *Visitor {
	return &Visitor{
		fset: fs,
	}
}

func (v *Visitor) Visit(n ast.Node) ast.Visitor {
	if n == nil {
		return nil
	}

	// Lets start by finding functions and printing their name.

	switch t := n.(type) {
	case *ast.FuncDecl:
		v.funcs = append(v.funcs, t.Name.Name)
	}

	return v
}
