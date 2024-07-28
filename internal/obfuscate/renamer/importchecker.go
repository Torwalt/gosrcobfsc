package renamer

import (
	"go/ast"
	"strings"
)

type importPath = string

type ImportChecker struct {
	externalImports map[importPath]struct{}
	internalImports map[importPath]struct{}
	moduleName      string
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
		externalImports: exImp,
		internalImports: inImp,
		moduleName:      moduleName,
	}
}

func (ic *ImportChecker) IsExternalImport(in string) bool {
	_, isExternal := ic.externalImports[in]

	return isExternal
}

func cleanImportString(in string) string {
	in = strings.TrimSpace(in)
	in, _ = strings.CutPrefix(in, "\"")
	in, _ = strings.CutSuffix(in, "\"")

	return in
}
