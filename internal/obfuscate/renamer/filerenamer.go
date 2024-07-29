package renamer

import (
	"fmt"
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
		fr.renameDecl(decl)
	}
}

func (fr *FileRenamer) renameSpecs(in []ast.Spec) {
	for _, s := range in {
		switch t := s.(type) {
		case *ast.TypeSpec:
			fr.renameTypeSpec(t)
		case *ast.ImportSpec:
		case *ast.ValueSpec:
			fr.renameValueSpec(t)
		default:
			fmt.Printf("Found unhandled in renameSpecs: %v\n", t)
		}
	}
}

func (fr *FileRenamer) renameValueSpec(in *ast.ValueSpec) {
	fr.renameIdentList(in.Names)
	fr.renameTypeExpr(in.Type)
	fr.renameExprList(in.Values)
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

	fr.renameFuncType(in.Type)
	fr.renameBlockStmt(in.Body)
}

func (fr *FileRenamer) renameBlockStmt(in *ast.BlockStmt) {
	fr.renameStmtList(in.List)
}

func (fr *FileRenamer) renameStmt(in ast.Stmt) {
	if in == nil {
		return
	}

	switch t := in.(type) {
	case *ast.AssignStmt:
		fr.renameAssignStatement(t)
	case *ast.IfStmt:
		fr.renameIfStmt(t)
	case *ast.ExprStmt:
		fr.renameExprStmt(t)
	case *ast.ReturnStmt:
		fr.renameReturnStmt(t)
	case *ast.RangeStmt:
		fr.renameRangeStmt(t)
	case *ast.DeclStmt:
		fr.renameDeclStmt(t)
	case *ast.SwitchStmt:
		fr.renameSwitchStmt(t)
	case *ast.CaseClause:
		fr.renameCaseClause(t)
	case *ast.TypeSwitchStmt:
		fr.renameTypeSwitchStmt(t)
	case *ast.BranchStmt:
		fr.renameBranchStmt(t)
	default:
		fmt.Printf("Found unhandled in renameStmt: %v\n", t)
	}
}

func (fr *FileRenamer) renameBranchStmt(in *ast.BranchStmt) {
	fr.renameIdent(in.Label)
}

func (fr *FileRenamer) renameTypeSwitchStmt(in *ast.TypeSwitchStmt) {
	fr.renameStmt(in.Init)
	fr.renameStmt(in.Assign)
	fr.renameBlockStmt(in.Body)
}

func (fr *FileRenamer) renameStmtList(in []ast.Stmt) {
	for _, stmt := range in {
		fr.renameStmt(stmt)
	}
}

func (fr *FileRenamer) renameCaseClause(in *ast.CaseClause) {
	fr.renameStmtList(in.Body)
	fr.renameExprList(in.List)
}

func (fr *FileRenamer) renameSwitchStmt(in *ast.SwitchStmt) {
	fr.renameBlockStmt(in.Body)
	fr.renameStmt(in.Init)
	fr.renameExpr(in.Tag)
}

func (fr *FileRenamer) renameDeclStmt(in *ast.DeclStmt) {
	fr.renameDecl(in.Decl)
}

func (fr *FileRenamer) renameDecl(in ast.Decl) {
	switch t := in.(type) {
	case *ast.FuncDecl:
		fr.renameFunc(t)
	case *ast.GenDecl:
		fr.renameSpecs(t.Specs)
	default:
		fmt.Printf("Found unhandled in renameDecl: %v\n", t)
	}
}

func (fr *FileRenamer) renameRangeStmt(in *ast.RangeStmt) {
	fr.renameBlockStmt(in.Body)
	fr.renameExpr(in.X)
	fr.renameExpr(in.Key)
	fr.renameExpr(in.Value)
}

func (fr *FileRenamer) renameExprList(in []ast.Expr) {
	for _, r := range in {
		fr.renameExpr(r)
	}
}

func (fr *FileRenamer) renameReturnStmt(in *ast.ReturnStmt) {
	fr.renameExprList(in.Results)
}

func (fr *FileRenamer) renameExprStmt(in *ast.ExprStmt) {
	fr.renameExpr(in.X)
}

func (fr *FileRenamer) renameIfStmt(in *ast.IfStmt) {
	fr.renameBlockStmt(in.Body)
	fr.renameStmt(in.Init)
	fr.renameExpr(in.Cond)
	fr.renameStmt(in.Else)
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

	fr.renameExpr(in.Fun)
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
	case *ast.BasicLit:
	case *ast.StarExpr:
	case *ast.UnaryExpr:
	case *ast.CompositeLit:
		fr.renameCompositeLit(t)
	case *ast.KeyValueExpr:
		fr.renameKeyValueExpr(t)
	case *ast.ArrayType:
		fr.renameArrayType(t)
	case *ast.FuncLit:
		fr.renameFuncLit(t)
	case *ast.SliceExpr:
		fr.renameSliceExpr(t)
	case *ast.IndexExpr:
		fr.renameIndexExpr(t)
	case *ast.BinaryExpr:
		fr.renameBinaryExpr(t)
	case *ast.TypeAssertExpr:
		fr.renameTypeAssertExpr(t)
	case *ast.MapType:
		fr.renameMapType(t)
	default:
		fmt.Printf("Found unhandled in renameExpr: %v\n", t)
	}
}

func (fr *FileRenamer) renameMapType(in *ast.MapType) {
	fr.renameTypeExpr(in.Key)
	fr.renameTypeExpr(in.Value)
}

func (fr *FileRenamer) renameTypeAssertExpr(in *ast.TypeAssertExpr) {
	fr.renameExpr(in.X)
	fr.renameTypeExpr(in.Type)
}

func (fr *FileRenamer) renameBinaryExpr(in *ast.BinaryExpr) {
	fr.renameTypeExpr(in.X)
	fr.renameTypeExpr(in.Y)
}

func (fr *FileRenamer) renameIndexExpr(in *ast.IndexExpr) {
	fr.renameExpr(in.X)
	fr.renameExpr(in.Index)
}

func (fr *FileRenamer) renameSliceExpr(in *ast.SliceExpr) {
	fr.renameExpr(in.X)
	fr.renameExpr(in.Low)
	fr.renameExpr(in.High)
	fr.renameExpr(in.Max)
}

func (fr *FileRenamer) renameFuncType(in *ast.FuncType) {
	fr.renameFieldList(in.TypeParams)
	fr.renameFieldList(in.Params)
	fr.renameFieldList(in.Results)
}

func (fr *FileRenamer) renameFuncLit(in *ast.FuncLit) {
	fr.renameBlockStmt(in.Body)
	fr.renameFuncType(in.Type)
}

func (fr *FileRenamer) renameArrayType(in *ast.ArrayType) {
	fr.renameExpr(in.Len)
	fr.renameExpr(in.Elt)
}

func (fr *FileRenamer) renameKeyValueExpr(in *ast.KeyValueExpr) {
	fr.renameExpr(in.Key)
	fr.renameExpr(in.Value)
}

func (fr *FileRenamer) renameCompositeLit(in *ast.CompositeLit) {
	fr.renameExprList(in.Elts)
	fr.renameTypeExpr(in.Type)
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

func (fr *FileRenamer) renameIdentList(in []*ast.Ident) {
	for _, i := range in {
		fr.renameIdent(i)
	}
}

func (fr *FileRenamer) renameField(in *ast.Field) {
	fr.renameIdentList(in.Names)
	fr.renameTypeExpr(in.Type)
}
