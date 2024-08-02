package obfuscate

import (
	"go/ast"

	"github.com/Torwalt/gosrcobfsc/internal/obfuscate/renamer"
	"github.com/Torwalt/gosrcobfsc/internal/paths"
	"github.com/Torwalt/gosrcobfsc/internal/repo"
)

type ObfuscatedRepository []ObfuscatedPackage

type Filename = string

type ObfuscatedPackage struct {
	Package         *repo.Package
	ObfuscatedPath  renamer.ObfuscatedPath
	ObfuscatedFiles map[Filename]renamer.ObfuscatedPath
}

func Obfuscate(rpo repo.Repository) (ObfuscatedRepository, error) {
	out := make(ObfuscatedRepository, 0)

	for _, pkgs := range rpo.Packages {
		for _, pkg := range pkgs.PkgMap {
			rename(pkg, rpo.ModuleName)
		}
		out = append(out, NewObfuscatedPackage(pkgs, rpo.Path))
	}

	return out, nil
}

func NewObfuscatedPackage(repoPkg *repo.Package, repoPath string) ObfuscatedPackage {
	op := renamer.RenamePackage(paths.NonRootPath(repoPkg.FullPath, repoPath))
	fileMap := make(map[Filename]renamer.ObfuscatedPath)
	for _, pkg := range repoPkg.PkgMap {
		for name := range pkg.Files {
			nrp := paths.NonRootPath(name, repoPath)
			r := renamer.RenamePackage(nrp)
			fileMap[nrp] = r
		}

	}

	return ObfuscatedPackage{
		Package:         repoPkg,
		ObfuscatedPath:  op,
		ObfuscatedFiles: fileMap,
	}
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
