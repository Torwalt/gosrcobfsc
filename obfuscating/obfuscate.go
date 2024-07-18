package obfuscating

import (
	"bytes"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
)

/*
Goal: Collect all "user named" tokens, rename them and output same structure.

https://eli.thegreenplace.net/2021/rewriting-go-source-code-with-ast-tooling/
*/

/*

Hard part: module package must be renamed, too, but external packages cannot be renamed.
Super hard part: What if code has custom private packages imported??

*/

type Args struct {
	// The module name, e.g. for this repo github.com/Torwalt/gosrcobfsc. We
	// need this to identify import paths and pkg names that must be changed.
	// Import paths of external pkgs must not be touched.
	ModuleName string
	// Full file path to the to be obfuscated repo.
	Source string
	// Directory where the new repo should be created.
	Sink string
}

/*
TODO:
- pkg and directory renaming
- lookup of packages and their paths to change imports
- args controlling
- writing resulting repo
- walking a whole repo not just file
*/

func Obfuscate(content string) (string, error) {
	fset := token.NewFileSet()

	f, err := parser.ParseFile(fset, "", content, 0)
	if err != nil {
		return "", err
	}

	v := NewVisitor(fset)
	ast.Walk(v, f)

	b := bytes.NewBufferString("")
	if err := format.Node(b, fset, f); err != nil {
		return "", err
	}

	return b.String(), nil
}

type Package struct {
	name     string
	fullPath string
}
