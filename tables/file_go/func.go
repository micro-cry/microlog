package file_go

import (
	"microlog/tables/file_go/generator_template"
	"os"
	"path/filepath"
)

// // // // // // // // // //

func clearOldDir(pathToDir string) error {
	items, err := os.ReadDir(pathToDir)
	if err != nil {
		return err
	}

	for _, item := range items {
		if item.IsDir() && len(item.Name()) > len(generator_template.DirPrefix) {
			if item.Name()[:len(generator_template.DirPrefix)] == generator_template.DirPrefix {
				fullPath := filepath.Join(pathToDir, item.Name())
				err = os.RemoveAll(fullPath)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func createDir(pathToDir, dirName string) (string, error) {
	newPath := filepath.Join(pathToDir, generator_template.DirPrefix+dirName)
	err := os.Mkdir(newPath, 0755)
	if err != nil {
		return "", err
	}

	return newPath, nil
}
