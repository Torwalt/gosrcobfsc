package renamer_test

import (
	"fmt"
	"go/parser"
	"go/token"
	"testing"

	"github.com/Torwalt/gosrcobfsc/internal/obfuscate/renamer"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var moduleName = "github.com/Torwalt/gosrcobfsc"

var file = `
package main

import (
	"flag"
	"log"
    "go/format"

	"github.com/Torwalt/gosrcobfsc/internal/args"
	"github.com/Torwalt/gosrcobfsc/internal/obfuscate"
	"github.com/Torwalt/gosrcobfsc/internal/repo"
)

func main() {
	moduleNameFlag := flag.String("moduleName", "", "The name of the module (top of go.mod).")
	sourceFlag := flag.String("source", "", "The full path of the source repository.")
	sinkFlag := flag.String("sink", "", "The full path where to write obfuscated directory.")
	flag.Parse()

	args, err := args.NewArgs(moduleNameFlag, sourceFlag, sinkFlag)
	if err != nil {
		flag.PrintDefaults()
		log.Fatalf("%v", err)
	}

	if err := run(args); err != nil {
		log.Fatalf("%v", err)
	}
}

func run(a args.Args) error {
	dirs, err := repo.CollectDirs(a.Source)
	if err != nil {
		return err
	}

	rpo, err := repo.NewRepository(dirs)
	if err != nil {
		return err
	}

	rpo, err = obfuscate.Obfuscate(rpo, a.ModuleName)
	if err != nil {
		return err
	}

	err = repo.WriteObfuscated(rpo, a)
	if err != nil {
		return err
	}

	log.Printf("Successfully obfuscated %v and wrote result into %v", a.Source, a.Sink)

	return nil
}
`

func TestNewImportChecker(t *testing.T) {
	content := file

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", content, 0)
	require.NoError(t, err)

	ic := renamer.NewImportChecker(f, moduleName)
	require.NotEmpty(t, ic)

	isExternalImportTests := []struct {
		path       string
		isExternal bool
	}{
		{
			path:       "flag",
			isExternal: true,
		},
		{
			path:       "log",
			isExternal: true,
		},
		{
			path:       "go/format",
			isExternal: true,
		},
		{
			path:       "github.com/Torwalt/gosrcobfsc/internal/args",
			isExternal: false,
		},
		{
			path:       "github.com/Torwalt/gosrcobfsc/internal/obfuscate",
			isExternal: false,
		},
		{
			path:       "github.com/Torwalt/gosrcobfsc/internal/repo",
			isExternal: false,
		},
	}

	for _, tt := range isExternalImportTests {
		t.Run(fmt.Sprintf("Test with path: %v", tt.path), func(t *testing.T) {
			isExternal := ic.IsExternalImport(tt.path)
			assert.Equal(t, tt.isExternal, isExternal)
		})
	}

	t.Run("Test HashImport", func(t *testing.T) {
		p := "github.com/Torwalt/gosrcobfsc/internal/repo"
		res := ic.HashImport(p)
		exp := "ddbggbhbcdbbaccffcbdbcijgabaacdfaiefdbcdagfbeaeeieccfjecahgaeiibXXX/dbedccbdadacfhbgaiefeaiecaccgicdffcaecgjhgddfecfeffcicjgfgbcafefXXX/ahbcaccchhfehafidhaadefjheiedfjejfiaabffdhgfcecfdjbaecbcchbcgdfeXXX"
		assert.Equal(t, exp, res)
	})

	t.Run("Test IsExternalPackage", func(t *testing.T) {
		isExternalImportTests := []struct {
			path       string
			isExternal bool
		}{
			{
				path:       "flag",
				isExternal: true,
			},
			{
				path:       "log",
				isExternal: true,
			},
			{
				path:       "format",
				isExternal: true,
			},
			{
				path:       "args",
				isExternal: false,
			},
		}
		for _, tt := range isExternalImportTests {
			t.Run(fmt.Sprintf("Test with path: %v", tt.path), func(t *testing.T) {
				isExternal := ic.IsExternalPackage(tt.path)
				assert.Equal(t, tt.isExternal, isExternal)
			})
		}
	})
}

func TestAddEscapedQuotes(t *testing.T) {
	act := renamer.AddEscapedQuotes("ahbcaccchhfehafidhaadefjheiedfjejfiaabffdhgfcecfdjbaecbcchbcgdfeXXX")
	exp := "\"ahbcaccchhfehafidhaadefjheiedfjejfiaabffdhgfcecfdjbaecbcchbcgdfeXXX\""
	assert.Equal(t, exp, act)
}
