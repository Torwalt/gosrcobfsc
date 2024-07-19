package obfuscating

import (
	"go/ast"
	"go/parser"
	"go/token"
)

type Repository []Package

type Package struct {
	fset     *token.FileSet
	name     string
	fullPath string
	pkgMap   map[string]*ast.Package
}

func NewRepository(dirs Dirs) (Repository, error) {
	repo := Repository{}
	for _, dir := range dirs {
		fset := token.NewFileSet()
		pkgMap, err := parser.ParseDir(fset, dir, nil, 0)
		if err != nil {
			return repo, err
		}
		repo = append(repo, Package{
			fset:     fset,
			name:     "",
			fullPath: dir,
			pkgMap:   pkgMap,
		})
	}

	return repo, nil
}
