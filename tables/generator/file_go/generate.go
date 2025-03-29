package file_go

import (
	"microlog/tables/generator"
	"os"
	"path/filepath"
)

// // // // // // // // // //

func Generate(tables []generator.InfoTableObj, pathToDir string) error {
	items, err := os.ReadDir(pathToDir)
	if err != nil {
		return err
	}

	for _, item := range items {
		if item.IsDir() {
			fullPath := filepath.Join(pathToDir, item.Name())
			err = os.RemoveAll(fullPath)
			if err != nil {
				return err
			}
		}
	}

	// //

	for _, item := range tables {
		newPath := filepath.Join(pathToDir, item.Name)
		err = os.Mkdir(newPath, 0755)
		if err != nil {
			return err
		}

		//

	}

	return nil
}
