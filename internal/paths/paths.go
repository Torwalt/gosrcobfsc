package paths

import (
	"path/filepath"
	"strings"
)

type FilterFunc func(in string) bool

func FilterEmpty(in string) bool {
	return in == ""
}

func SplitAndFilter(path string, ff FilterFunc) []string {
	out := []string{}
	for _, part := range strings.Split(path, string(filepath.Separator)) {
		if ff(part) {
			continue
		}
		out = append(out, part)
	}

	return out
}

func NonRootPath(fullPath, root string) string {
	nonRootPath, _ := strings.CutPrefix(fullPath, root)
	return nonRootPath
}
