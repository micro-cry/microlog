package generator_go

import (
	"os"
	"path/filepath"
)

// // // // // // // // // //

func DirRemoveAll(pathToDir string) error {
	items, err := os.ReadDir(pathToDir)
	if err != nil {
		return err
	}

	for _, item := range items {
		if item.IsDir() && len(item.Name()) > len(DirPrefix) {
			if item.Name()[:len(DirPrefix)] == DirPrefix {
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

func DirCreate(pathToDir, dirName string) (string, error) {
	newPath := filepath.Join(pathToDir, DirPrefix+dirName)
	err := os.Mkdir(newPath, 0755)
	if err != nil {
		return "", err
	}

	return newPath, nil
}
