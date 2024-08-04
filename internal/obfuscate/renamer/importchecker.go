package renamer

import (
	"fmt"
	"go/ast"
	"path/filepath"
	"strings"

	"github.com/Torwalt/gosrcobfsc/internal/hasher"
	"github.com/Torwalt/gosrcobfsc/internal/paths"
)

type (
	importPath  = string
	packageName = string
)

type imprt struct {
	fullPath    importPath
	packageName packageName
}

type ImportChecker struct {
	externalImports  map[importPath]struct{}
	externalPkgs     map[packageName]struct{}
	moduleName       string
	hashedModuleName string
}

func NewImportChecker(file *ast.File, moduleName string) *ImportChecker {
	exImp := make(map[importPath]struct{})
	exPkgs := make(map[packageName]struct{})

	for _, f := range file.Imports {
		path := cleanImportString(f.Path.Value)
		if strings.Contains(path, moduleName) {
			continue
		}

		pkg := filepath.Base(path)

		exPkgs[pkg] = struct{}{}
		exImp[path] = struct{}{}
	}

	return &ImportChecker{
		externalImports:  exImp,
		externalPkgs:     exPkgs,
		moduleName:       moduleName,
		hashedModuleName: hasher.Hash(moduleName),
	}
}

func (ic *ImportChecker) IsExternalPackage(in string) bool {
	_, isExternal := ic.externalPkgs[in]

	return isExternal
}

func (ic *ImportChecker) IsExternalImport(in string) bool {
	in = cleanImportString(in)
	_, isExternal := ic.externalImports[in]

	return isExternal
}

func (ic *ImportChecker) HashImport(in string) string {
	in = cleanImportString(in)
	fp, _ := strings.CutPrefix(in, ic.moduleName)

	pathParts := paths.SplitAndFilter(fp, paths.FilterEmpty)

	hashedParts := []string{}
	hashedParts = append(hashedParts, ic.hashedModuleName)
	for _, part := range pathParts {
		hashedParts = append(hashedParts, hasher.Hash(part))
	}

	return filepath.Join(hashedParts...)
}

func AddEscapedQuotes(in string) string {
	return fmt.Sprintf("\"%v\"", in)
}

func cleanImportString(in string) string {
	in = strings.TrimSpace(in)
	in, _ = strings.CutPrefix(in, "\"")
	in, _ = strings.CutSuffix(in, "\"")

	return in
}
