package repo_test

import (
	"testing"

	"github.com/Torwalt/gosrcobfsc/internal/repo"
	"github.com/stretchr/testify/require"
)

func TestCollectDirs(t *testing.T) {
	d, err := repo.CollectDirs(thisRepoFullPath)
	require.NoError(t, err)
	require.NotEmpty(t, d)
	require.NotContains(t, d, "/home/ada/repos/gosrcobfsc/.git")
}
