package generator_go

import (
	"os"
	"path/filepath"
	"strings"
	"unicode"
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

// // //

func NameValGo(s string) string {
	if len(s) == 0 {
		return s
	}

	runes := []rune(s)
	first := unicode.ToUpper(runes[0])
	rest := strings.ToLower(string(runes[1:]))
	return string(first) + rest
}
