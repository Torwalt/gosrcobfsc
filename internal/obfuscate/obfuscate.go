package obfuscate

import (
	"go/ast"

	"github.com/Torwalt/gosrcobfsc/internal/hasher"
	"github.com/Torwalt/gosrcobfsc/internal/obfuscate/renamer"
	"github.com/Torwalt/gosrcobfsc/internal/paths"
	"github.com/Torwalt/gosrcobfsc/internal/repo"
	"github.com/Torwalt/gosrcobfsc/internal/repo/gomod"
)

type ObfuscatedRepository struct {
	Packages        []ObfuscatedPackage
	ObfuscatedGomod gomod.GoMod
}

type Filename = string

type ObfuscatedPackage struct {
	Package         *repo.Package
	ObfuscatedPath  renamer.ObfuscatedPath
	ObfuscatedFiles map[Filename]renamer.ObfuscatedPath
}

func Obfuscate(rpo repo.Repository) (ObfuscatedRepository, error) {
	out := ObfuscatedRepository{}
	ops := make([]ObfuscatedPackage, 0)

	for _, pkgs := range rpo.Packages {
		for _, pkg := range pkgs.PkgMap {
			rename(pkg, rpo.Gomod.ModuleName)
		}
		ops = append(ops, NewObfuscatedPackage(pkgs, rpo.Path))
	}

	out.Packages = ops
	out.ObfuscatedGomod = gomod.GoMod{
		Version:    rpo.Gomod.Version,
		ModuleName: hasher.Hash(rpo.Gomod.ModuleName),
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
		fr := renamer.NewFileRenamer(file, ic)
		fr.Rename()
	}
}
