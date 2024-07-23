package obfuscate_test

import (
	"path/filepath"
	"testing"

	"github.com/Torwalt/gosrcobfsc/internal/args"
	"github.com/Torwalt/gosrcobfsc/internal/obfuscate"
	"github.com/stretchr/testify/require"
)

var (
	moduleName       = "github.com/Torwalt/gosrcobfsc"
	sink             = filepath.Join(thisRepoFullPath, "tests")
	thisRepoFullPath = "/home/ada/repos/gosrcobfsc/"
)

func TestObfuscate(t *testing.T) {
	sink := t.TempDir()
	args, err := args.NewArgs(&moduleName, &thisRepoFullPath, &sink)
	require.NoError(t, err)

	_, err = obfuscate.Obfuscate(args)
	require.NoError(t, err)
}
