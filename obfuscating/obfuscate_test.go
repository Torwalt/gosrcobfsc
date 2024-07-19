package obfuscating_test

import (
	"path/filepath"
	"testing"

	"github.com/Torwalt/gosrcobfsc/obfuscating"
	"github.com/stretchr/testify/require"
)

var (
	moduleName = "github.com/Torwalt/gosrcobfsc"
	sink       = filepath.Join(thisRepoFullPath, "tests")
)

func TestObfuscate(t *testing.T) {
	args, err := obfuscating.NewArgs(&moduleName, &thisRepoFullPath, &sink)
	require.NoError(t, err)

	err = obfuscating.Obfuscate(args)
	require.NoError(t, err)
}
