package renamer

import (
	"go/ast"

	"github.com/Torwalt/gosrcobfsc/internal/hasher"
)

const (
	mainFunc  = "main"
	errorType = "error"
)

/* TODO:
- create directories and files with hashed names
- hash module pkgs, but dont hash other pkgs (e.g. std lib, external pkgs)
*/

type FileRenamer struct {
	pkg  *ast.Package
	file *ast.File
	ic   *ImportChecker
}

func NewFileRenamer(pkg *ast.Package, file *ast.File, ic *ImportChecker) *FileRenamer {
	return &FileRenamer{
		pkg:  pkg,
		file: file,
		ic:   ic,
	}
}

func (fr *FileRenamer) Rename() {
	for _, decl := range fr.file.Decls {
		switch t := decl.(type) {
		case *ast.FuncDecl:
			fr.renameFunc(t)
		case *ast.GenDecl:
			fr.renameSpecs(t.Specs)
		default:
			_ = ""
			print("")
		}
	}
}

func (fr *FileRenamer) renameSpecs(in []ast.Spec) {
	for _, s := range in {
		switch t := s.(type) {
		case *ast.TypeSpec:
			fr.renameTypeSpec(t)
		default:
			_ = ""
			print("")
		}
	}
}

func (fr *FileRenamer) renameTypeSpec(in *ast.TypeSpec) {
	fr.renameIdent(in.Name)
	fr.renameTypeExpr(in.Type)
}

func (fr *FileRenamer) renameStructType(in *ast.StructType) {
	fr.renameFieldList(in.Fields)
}

func (fr *FileRenamer) renameFunc(in *ast.FuncDecl) {
	if in == nil {
		return
	}

	if in.Name.Name != mainFunc {
		fr.renameIdent(in.Name)
	}

	fr.renameFieldList(in.Type.TypeParams)
	fr.renameFieldList(in.Type.Params)
	fr.renameFieldList(in.Type.Results)

	for _, stmt := range in.Body.List {
		switch t := stmt.(type) {
		case *ast.AssignStmt:
			fr.renameAssignStatement(t)
		case *ast.IfStmt:
			fr.renameIfStmt(t)
		default:
			_ = ""
			print("")
		}
	}
}

func (fr *FileRenamer) renameIfStmt(in *ast.IfStmt) {
}

func (fr *FileRenamer) renameAssignStatement(in *ast.AssignStmt) {
	if in == nil {
		return
	}

	for _, l := range in.Lhs {
		fr.renameExpr(l)
	}

	for _, r := range in.Rhs {
		fr.renameExpr(r)
	}
}

func (fr *FileRenamer) renameCallExpr(in *ast.CallExpr) {
	if in == nil {
		return
	}

	switch t := in.Fun.(type) {
	case *ast.SelectorExpr:
		fr.renameSelectorExpr(t)
	}
	for _, a := range in.Args {
		fr.renameExpr(a)
	}

}

func (fr *FileRenamer) renameSelectorExpr(in *ast.SelectorExpr) {
	if in == nil {
		return
	}

	// We must not rename external symbols from external packages.
	if fr.isExternal(in) {
		return
	}

	fr.renameIdent(in.Sel)
	fr.renameExpr(in.X)
}

func (fr *FileRenamer) isExternal(in *ast.SelectorExpr) bool {
	switch t := in.X.(type) {
	case *ast.Ident:
		return fr.ic.IsExternalImport(t.Name)
	}

	return false
}

func (fr *FileRenamer) renameExpr(in ast.Expr) {
	if in == nil {
		return
	}

	switch t := in.(type) {
	case *ast.Ident:
		fr.renameIdent(t)
	case *ast.CallExpr:
		fr.renameCallExpr(t)
	case *ast.SelectorExpr:
		fr.renameSelectorExpr(t)
	case *ast.StructType:
		fr.renameStructType(t)
	}
}

// renameTypeExpr handles special Expr cases where we need to be sure the type
// is not a primitive type, e.g. int, error, string, etc.
func (fr *FileRenamer) renameTypeExpr(in ast.Expr) {
	switch t := in.(type) {
	case *ast.Ident:
		if isNonrenameableType(t.Name) {
			return
		}
	}

	fr.renameExpr(in)
}

func (fr *FileRenamer) renameIdent(in *ast.Ident) {
	if in == nil {
		return
	}

	in.Name = hasher.Hash(in.Name)
}

func (fr *FileRenamer) renameFieldList(in *ast.FieldList) {
	if in == nil {
		return
	}

	for _, f := range in.List {
		fr.renameField(f)
	}
}

func (fr *FileRenamer) renameField(in *ast.Field) {
	for _, n := range in.Names {
		fr.renameIdent(n)
	}

	fr.renameTypeExpr(in.Type)
}
