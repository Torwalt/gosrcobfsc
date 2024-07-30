package gitignore_test

import (
	"fmt"
	"testing"

	"github.com/Torwalt/gosrcobfsc/internal/repo/gitignore"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var thisRepoFullPath = "/home/ada/repos/gosrcobfsc/"

func TestNewFromFilePath(t *testing.T) {
	gi, err := gitignore.NewFromFilePath(thisRepoFullPath)
	require.NoError(t, err)
	require.NotEmpty(t, gi)
}

func TestNewGitIgnore(t *testing.T) {
	gitignoreContent := `
tests/*
gosrcobfsc
    `

	tsts := []struct {
		path       string
		isExcluded bool
	}{
		// {
		// 	path:       ".git",
		// 	isExcluded: true,
		// },
		// {
		// 	path:       "git",
		// 	isExcluded: false,
		// },
		// {
		// 	path:       "gyatt",
		// 	isExcluded: false,
		// },
		// {
		// 	path:       "tests/cmd/main.go",
		// 	isExcluded: true,
		// },
		{
			path:       "/home/ada/repos/gosrcobfsc/internal/repo/gitignore",
			isExcluded: false,
		},
	}

	for _, tt := range tsts {
		t.Run(fmt.Sprintf("Test with path: %v", tt.path), func(t *testing.T) {
			gi := gitignore.NewGitIgnore(gitignoreContent)
			assert.Equal(t, tt.isExcluded, gi.PathExcluded(tt.path))
		})
	}
}
