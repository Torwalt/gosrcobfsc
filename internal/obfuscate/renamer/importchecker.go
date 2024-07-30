package renamer

import (
	"fmt"
	"go/ast"
	"path/filepath"
	"strings"

	"github.com/Torwalt/gosrcobfsc/internal/hasher"
)

type importPath = string

type ImportChecker struct {
	externalImports  map[importPath]struct{}
	internalImports  map[importPath]struct{}
	moduleName       string
	hashedModuleName string
}

func NewImportChecker(file *ast.File, moduleName string) *ImportChecker {
	exImp := make(map[string]struct{})
	inImp := make(map[string]struct{})

	for _, f := range file.Imports {
		path := cleanImportString(f.Path.Value)
		if strings.Contains(path, moduleName) {
			inImp[path] = struct{}{}
			continue
		}

		exImp[path] = struct{}{}
	}

	return &ImportChecker{
		externalImports:  exImp,
		internalImports:  inImp,
		moduleName:       moduleName,
		hashedModuleName: hasher.Hash(moduleName),
	}
}

func (ic *ImportChecker) IsExternalImport(in string) bool {
	_, isExternal := ic.externalImports[in]

	return isExternal
}

func (ic *ImportChecker) HashImport(in string) string {
	fp, _ := strings.CutPrefix(in, ic.moduleName)
	pathParts := strings.Split(fp, string(filepath.Separator))

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
