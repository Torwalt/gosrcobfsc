package repo

import (
	"errors"
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"path/filepath"
	"strings"
	"sync"

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
	Info     *types.Info
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
			Info: &types.Info{
				Uses: make(map[*ast.Ident]types.Object, 1000),
			},
		})
	}

	gmd, err := gomod.NewGoModFromPath(repoPath)
	if err != nil {
		return Repository{}, errors.New(fmt.Sprintf("could not parse gomod file in given path: %v", err))
	}

	wg := sync.WaitGroup{}
	wg.Add(len(pkgs))
	for _, p := range pkgs {
		go func() {
			defer wg.Done()

			config := &types.Config{
				Importer: importer.ForCompiler(p.Fset, "source", nil),
			}
			for _, sp := range p.PkgMap {
				files := make([]*ast.File, 0, len(sp.Files))
				for _, f := range sp.Files {
					files = append(files, f)
				}

				pkgPath, _ := strings.CutPrefix(p.FullPath, repoPath)
				path := filepath.Join(gmd.ModuleName, pkgPath)

				if _, err := config.Check(path, p.Fset, files, p.Info); err != nil {
					panic(err)
				}
			}
		}()
	}
	wg.Wait()

	return Repository{
		Packages: pkgs,
		Path:     repoPath,
		Gomod:    gmd,
	}, nil
}
