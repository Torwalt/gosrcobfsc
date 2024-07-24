package obfuscate

import (
	"go/ast"

	"github.com/Torwalt/gosrcobfsc/internal/repo"
)

/*

Hard part: module package must be renamed, too, but external packages cannot be renamed.
Super hard part: What if code has custom private packages imported??

*/

/*
TODO:
- pkg and directory renaming
- lookup of packages and their paths to change imports
*/

func Obfuscate(rpo repo.Repository) (repo.Repository, error) {
	// I.e., we need to walk the rpo and mutate the NamedSymbols.
	for _, pkgs := range rpo {
		for _, pkg := range pkgs.PkgMap {
			v := NewVisitor()
			ast.Walk(v, pkg)
			rename(v.ns)
		}
	}

	return rpo, nil
}
