package obfuscate

import (
	"go/ast"

	"github.com/Torwalt/gosrcobfsc/internal/obfuscate/renamer"
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

func Obfuscate(rpo repo.Repository, moduleName string) (repo.Repository, error) {
	for _, pkgs := range rpo {
		for _, pkg := range pkgs.PkgMap {
			rename(pkg, moduleName)
		}
	}

	return rpo, nil
}

func rename(pkg *ast.Package, moduleName string) {
	v := NewVisitor()
	ast.Walk(v, pkg)

	for _, file := range v.ns.Files {
		ic := renamer.NewImportChecker(file, moduleName)
		fr := renamer.NewFileRenamer(pkg, file, ic)
		fr.Rename()
	}
}
