package generators

import (
	"fmt"
	"microlog"
	"microlog/generator_go"
)

// // // // // // // // // //

var generatorArr = []func(string) generator_go.GeneratorInterface{
	NewStruct,
	NewValues,
	NewFunc,

	NewFuncTest,

	NewSQLite,
	NewSQLiteGet,
	NewSQLiteOther,
	NewSQLiteTable,
}

func Generate(tables []microlog.InfoTableObj, rootDirName, pathToDir string) error {
	err := generator_go.DirRemoveAll(pathToDir)
	if err != nil {
		return fmt.Errorf("removing path to dir [%s]: %s", pathToDir, err.Error())
	}

	for _, item := range tables {
		newPath, err := generator_go.DirCreate(pathToDir, item.Name)
		if err != nil {
			return fmt.Errorf("create dir [%s/%s]: %s", pathToDir, item.Name, err.Error())
		}

		//

		for _, f := range generatorArr {
			if e := f(rootDirName).Generator(newPath, &item); e != nil {
				return e
			}
		}
	}

	return nil
}
