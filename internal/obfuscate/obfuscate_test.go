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

	rpo, err := repo.NewRepository(dirs, thisRepoFullPath)
	require.NoError(t, err)

	op, err := obfuscate.Obfuscate(rpo)
	require.NoError(t, err)
	require.NotEmpty(t, op)
}
