package obfuscate_test

import (
	"path/filepath"
	"testing"

	"github.com/Torwalt/gosrcobfsc/internal/obfuscate"
	"github.com/Torwalt/gosrcobfsc/internal/repo"
	"github.com/stretchr/testify/require"
)

var (
	moduleName       = "github.com/Torwalt/gosrcobfsc"
	sink             = filepath.Join(thisRepoFullPath, "tests")
	thisRepoFullPath = "/home/ada/repos/gosrcobfsc/"
)

func TestObfuscate(t *testing.T) {
	dirs, err := repo.CollectDirs(thisRepoFullPath)
	require.NoError(t, err)

	rpo, err := repo.NewRepository(dirs)
	require.NoError(t, err)

	_, err = obfuscate.Obfuscate(rpo)
	require.NoError(t, err)
}
