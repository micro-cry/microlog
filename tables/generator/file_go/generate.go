package file_go

import (
	"microlog/tables/generator"
)

// // // // // // // // // //

var generatorArr = make([]func(string, *generator.InfoTableObj) error, 0)

//

func Generate(tables []generator.InfoTableObj, pathToDir string) error {
	err := clearOldDir(pathToDir)
	if err != nil {
		return err
	}

	pathsMap := make(map[string]*generator.InfoTableObj)

	for _, item := range tables {
		newPath, err := createDir(pathToDir, item.Name)
		if err != nil {
			return err
		}

		//

		for _, f := range generatorArr {
			if e := f(newPath, &item); e != nil {
				return e
			}
		}

		pathsMap[newPath] = &item
	}

	return nil
}
