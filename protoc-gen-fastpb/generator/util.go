package generator

import (
	"path/filepath"
)

func adjustPath(path string) (ret string) {
	cur, _ := filepath.Abs(".")
	if filepath.IsAbs(path) {
		path, _ = filepath.Rel(cur, path)
	}
	return path
}
