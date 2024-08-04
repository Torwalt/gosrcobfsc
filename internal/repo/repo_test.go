package repo_test

import (
	"path/filepath"
	"testing"

	"github.com/Torwalt/gosrcobfsc/internal/args"
	"github.com/Torwalt/gosrcobfsc/internal/repo"
	"github.com/Torwalt/gosrcobfsc/internal/repo/gitignore"
	"github.com/stretchr/testify/require"
)

var (
	thisRepoFullPath = "/home/ada/repos/gosrcobfsc"
	moduleName       = "github.com/Torwalt/gosrcobfsc"
	sink             = filepath.Join(thisRepoFullPath, "tests")
)

func TestNewRepository(t *testing.T) {
	sink := t.TempDir()
	args, err := args.NewArgs(&moduleName, &thisRepoFullPath, &sink)
	require.NoError(t, err)

	gi, err := gitignore.NewFromFilePath(thisRepoFullPath)
	require.NoError(t, err)

	dirs, err := repo.CollectDirs(args.Source, repo.FilterFuncWithGitIgnore(gi, args.Source))
	require.NoError(t, err)

	rpo, err := repo.NewRepository(dirs, args.Source)
	require.NoError(t, err)
	require.NotEmpty(t, rpo)
}
