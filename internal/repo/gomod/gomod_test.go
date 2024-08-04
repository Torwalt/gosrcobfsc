package gomod_test

import (
	"testing"

	"github.com/Torwalt/gosrcobfsc/internal/repo/gomod"
	"github.com/stretchr/testify/require"
)

var testGomod = `
module github.com/Torwalt/gosrcobfsc

go 1.22.4

require github.com/stretchr/testify v1.9.0

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
`

func TestNewGoMod(t *testing.T) {
	gmd, err := gomod.NewGoMod(testGomod)
	expGmd := gomod.GoMod{
		Version:    "1.22.4",
		ModuleName: "github.com/Torwalt/gosrcobfsc",
	}

	require.NoError(t, err)
	require.Equal(t, expGmd, gmd)
}
