- super hard problem:

```go
type ObfuscatedPackage struct {
	Package         *repo.Package
	ObfuscatedPath  renamer.ObfuscatedPath
    Expr *regexp.Regexp
}

func (op *ObfuscatedPackage) SomeMethod() {
    op.Expr.MatchString("asd")
}
```

- how to hash op, Expr but not MatchString?
  - look into `Info     *types.Info` and WIP example parsing in internal/repo/repo.go

- remove all comments
