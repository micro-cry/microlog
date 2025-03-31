package file_go

import (
	"microlog/tables"
	generator_template2 "microlog/tables/file_go/generator_template"
)

// // // // // // // // // //

var generatorArr = []GeneratorInterface{
	new(generator_template2.FuncObj),
	new(generator_template2.StructObj),
	new(generator_template2.ValuesObj),

	new(generator_template2.FuncTestObj),

	new(generator_template2.SQLiteObj),
	new(generator_template2.SQLiteGetObj),
	new(generator_template2.SQLiteOtherObj),
	new(generator_template2.SQLiteTableObj),
}

//

func Generate(tables []tables.InfoTableObj, rootDirName, pathToDir string) error {
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

		for _, obj := range generatorArr {
			if e := obj.Generator(newPath, &item); e != nil {
				return e
			}
		}
	}

	return nil
}
