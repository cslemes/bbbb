package blog

import (
	"os"
	"path/filepath"
)

func getMarkdownFiles(dir string) ([]string, error) {
	var files []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".md" {
			files = append(files, filepath.Base(path))
		}
		return nil
	})
	return files, err
}
