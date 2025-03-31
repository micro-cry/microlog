package generator_go

import (
	"os"
	"path/filepath"
	"strings"
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
	err := os.MkdirAll(newPath, 0755)
	if err != nil {
		return "", err
	}

	return newPath, nil
}

func PathImport(rootDirName, dirPath string) string {
	bufPath := strings.Split(dirPath, "/")

	for pos, path := range bufPath {
		if path == rootDirName {
			if pos < len(bufPath)-1 || bufPath[pos+1] == rootDirName {
				bufPath[pos+1] = ""
			}
			break
		}
		bufPath[pos] = ""
	}
	bufPath[len(bufPath)-1] = ""

	var importPath []string
	for _, s := range bufPath {
		if s != "" {
			importPath = append(importPath, s)
		}
	}

	return strings.Join(importPath, "/")
}
