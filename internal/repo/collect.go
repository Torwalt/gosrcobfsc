package repo

import (
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/Torwalt/gosrcobfsc/internal/repo/gitignore"
)

type Dirs = []string

type FilterFunc func(path string) bool

func CollectDirs(source string, ff FilterFunc) (Dirs, error) {
	dirs := Dirs{}

	walkerFunc := func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() {
			return nil
		}

		if ff(path) {
			return nil
		}

		dirs = append(dirs, path)

		return nil
	}

	if err := filepath.WalkDir(source, walkerFunc); err != nil {
		return Dirs{}, err
	}

	return dirs, nil
}

func NoFilter(_ string) bool {
	return false
}

func FilterFuncWithGitIgnore(gi gitignore.GitIgnore, rootPath string) FilterFunc {
	return func(path string) bool {
		trimmed := cutRootPath(path, rootPath)
		return gi.PathExcluded(trimmed)
	}
}

func cutRootPath(path, root string) string {
	if path == root {
		return path
	}

	out, _ := strings.CutPrefix(path, root)

	return out
}
