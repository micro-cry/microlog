package generators

import (
	"microlog"
	"microlog/generator_go"
)

// // // // // // // // // //

var generatorArr = []generator_go.GeneratorInterface{
	new(FuncObj),
	new(StructObj),
	new(ValuesObj),

	new(FuncTestObj),

	new(SQLiteObj),
	new(SQLiteGetObj),
	new(SQLiteOtherObj),
	new(SQLiteTableObj),
}

//

func Generate(tables []microlog.InfoTableObj, rootDirName, pathToDir string) error {
	err := generator_go.DirRemoveAll(pathToDir)
	if err != nil {
		return err
	}

	for _, item := range tables {
		newPath, err := generator_go.DirCreate(pathToDir, item.Name)
		if err != nil {
			return err
		}

		//

		for _, obj := range generatorArr {
			if e := obj.Generator(newPath, &item); e != nil {
				return e
			}
		}
	}

	return nil
}
