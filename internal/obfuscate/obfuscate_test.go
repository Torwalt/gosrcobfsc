package obfuscate_test

import (
	"path/filepath"
	"testing"

	"github.com/Torwalt/gosrcobfsc/internal/obfuscate"
	"github.com/Torwalt/gosrcobfsc/internal/repo"
	"github.com/Torwalt/gosrcobfsc/internal/repo/gitignore"
	"github.com/stretchr/testify/require"
)

var (
	moduleName       = "github.com/Torwalt/gosrcobfsc"
	sink             = filepath.Join(thisRepoFullPath, "tests")
	thisRepoFullPath = "/home/ada/repos/gosrcobfsc/"
)

func TestObfuscate(t *testing.T) {
	gi, err := gitignore.NewFromFilePath(thisRepoFullPath)
	require.NoError(t, err)

	dirs, err := repo.CollectDirs(thisRepoFullPath, repo.FilterFuncWithGitIgnore(gi, thisRepoFullPath))
	require.NoError(t, err)

	rpo, err := repo.NewRepository(dirs)
	require.NoError(t, err)

	_, err = obfuscate.Obfuscate(rpo, moduleName)
	require.NoError(t, err)
}
