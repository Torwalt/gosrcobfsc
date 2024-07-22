package obfuscating

import (
	"go/ast"
)

// "bytes"
// "go/ast"
// "go/format"
// "go/parser"
// "go/token"

/*
Goal: Collect all "user named" tokens, rename them and output same structure.

https://eli.thegreenplace.net/2021/rewriting-go-source-code-with-ast-tooling/
*/

/*

Hard part: module package must be renamed, too, but external packages cannot be renamed.
Super hard part: What if code has custom private packages imported??

*/

/*
TODO:
- pkg and directory renaming
- lookup of packages and their paths to change imports
- args controlling
- writing resulting repo
- walking a whole repo not just file
*/

type Obfuscated struct {
	Pkg Package
	Ns  NamedSymbols
}

func Obfuscate(args Args) ([]Obfuscated, error) {
	dirs, err := CollectDirs(args.Source)
	if err != nil {
		return []Obfuscated{}, err
	}

	repo, err := NewRepository(dirs)
	if err != nil {
		return []Obfuscated{}, err
	}

	obf := []Obfuscated{}
	for _, pkgs := range repo {
		for _, pkg := range pkgs.pkgMap {
			v := NewVisitor()
			ast.Walk(v, pkg)
			obf = append(obf, Obfuscated{
				Pkg: pkgs,
				Ns:  v.NamedSymbols(),
			})
		}
	}

	// Obfuscate every user defined symbol in nsSet.
	// Create Identical dir structure in sink.
	// Write obfuscated fset.

	// f, err := parser.ParseFile(fset, "", content, 0)
	// if err != nil {
	// 	return "", err
	// }
	//
	// v := NewVisitor(fset)
	// ast.Walk(v, f)
	//
	// b := bytes.NewBufferString("")
	// if err := format.Node(b, fset, f); err != nil {
	// 	return "", err
	// }
	//
	// return b.String(), nil

	return obf, nil
}
