package repo

import (
	"go/ast"
	"go/parser"
	"go/token"
)

type Repository struct {
	Packages   []*Package
	Path       string
	ModuleName string
}

type Package struct {
	Fset     *token.FileSet
	FullPath string
	PkgMap   map[string]*ast.Package
}

func NewRepository(dirs Dirs, repoPath, moduleName string) (Repository, error) {
	pkgs := make([]*Package, 0, len(dirs))
	for _, dir := range dirs {
		fset := token.NewFileSet()
		pkgMap, err := parser.ParseDir(fset, dir, nil, 0)
		if err != nil {
			return Repository{}, err
		}
		pkgs = append(pkgs, &Package{
			Fset:     fset,
			FullPath: dir,
			PkgMap:   pkgMap,
		})
	}

	return Repository{
		Packages:   pkgs,
		Path:       repoPath,
		ModuleName: moduleName,
	}, nil
}
