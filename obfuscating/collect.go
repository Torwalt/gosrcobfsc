package obfuscating

import (
	"io/fs"
	"path/filepath"
	"strings"
)

type Dirs = []string

func CollectDirs(source string) (Dirs, error) {
	dirs := Dirs{}

	walkerFunc := func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() {
			return nil
		}

		if !keep(path) {
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

func keep(path string) bool {
	if strings.Contains(path, "/.git") {
		return false
	}

	return true
}
