package repo_test

import (
	"path/filepath"
	"testing"

	"github.com/Torwalt/gosrcobfsc/internal/args"
	"github.com/Torwalt/gosrcobfsc/internal/repo"
	"github.com/stretchr/testify/require"
)

var (
	thisRepoFullPath = "/home/ada/repos/gosrcobfsc/"
	moduleName       = "github.com/Torwalt/gosrcobfsc"
	sink             = filepath.Join(thisRepoFullPath, "tests")
)

func TestWriteObfuscated(t *testing.T) {
	sink := t.TempDir()
	args, err := args.NewArgs(&moduleName, &thisRepoFullPath, &sink)
	require.NoError(t, err)

	dirs, err := repo.CollectDirs(args.Source)
	require.NoError(t, err)

	rpo, err := repo.NewRepository(dirs)
	require.NoError(t, err)

	err = repo.WriteObfuscated(rpo, args)
	require.NoError(t, err)
}
