package repo_test

import (
	"path/filepath"
	"testing"

	"github.com/Torwalt/gosrcobfsc/internal/repo"
	"github.com/Torwalt/gosrcobfsc/internal/repo/gitignore"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCollectDirs(t *testing.T) {
	allDirs, err := repo.CollectDirs(thisRepoFullPath, repo.NoFilter)
	require.NoError(t, err)
	require.NotEmpty(t, allDirs)
	require.Contains(t, allDirs, "/home/ada/repos/gosrcobfsc/internal/repo/gitignore")

	gi, err := gitignore.NewFromFilePath(thisRepoFullPath)
	require.NoError(t, err)

	filtered, err := repo.CollectDirs(thisRepoFullPath, repo.FilterFuncWithGitIgnore(gi, thisRepoFullPath))
	require.NoError(t, err)

	assert.NotContains(t, filtered, filepath.Join(thisRepoFullPath, ".git"))
	assert.Contains(t, filtered, "/home/ada/repos/gosrcobfsc/internal/repo/gitignore")
	assert.NotContains(t, filtered, filepath.Join(thisRepoFullPath, "tests"))
	assert.NotContains(t, filtered, filepath.Join(thisRepoFullPath, "gosrcobfsc"))
}
