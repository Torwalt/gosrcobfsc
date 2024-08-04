package repo

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"

	"github.com/Torwalt/gosrcobfsc/internal/repo/gomod"
)

type Repository struct {
	Packages []*Package
	Path     string
	Gomod    gomod.GoMod
}

type Package struct {
	Fset     *token.FileSet
	FullPath string
	PkgMap   map[string]*ast.Package
}

func NewRepository(dirs Dirs, repoPath string) (Repository, error) {
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

	gmd, err := gomod.NewGoModFromPath(repoPath)
	if err != nil {
		return Repository{}, errors.New(fmt.Sprintf("could not parse gomod file in given path: %v", err))
	}

	return Repository{
		Packages: pkgs,
		Path:     repoPath,
		Gomod:    gmd,
	}, nil
}
