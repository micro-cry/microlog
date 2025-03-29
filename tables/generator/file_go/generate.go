package file_go

import (
	"microlog/tables/generator"
)

// // // // // // // // // //

const (
	DirPrefix = "table_"
)

func Generate(tables []generator.InfoTableObj, pathToDir string) error {
	err := clearOldDir(pathToDir)
	if err != nil {
		return err
	}

	for _, item := range tables {
		newPath, err := createDir(pathToDir, item.Name)
		if err != nil {
			return err
		}

		//

		if e := generateStruct(newPath, &item); e != nil {
			return e
		}

		if e := generateValues(newPath, &item); e != nil {
			return e
		}

		if e := generateFunc(newPath, &item); e != nil {
			return e
		}
	}

	return nil
}
