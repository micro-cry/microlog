package file_go

import (
	"microlog/tables/generator"
	"microlog/tables/generator/file_go/generator_template"
)

// // // // // // // // // //

var generatorArr = []GeneratorInterface{
	new(generator_template.FuncObj),
	new(generator_template.StructObj),
	new(generator_template.ValuesObj),

	new(generator_template.FuncTestObj),

	new(generator_template.SQLiteObj),
	new(generator_template.SQLiteGetObj),
	new(generator_template.SQLiteOtherObj),
	new(generator_template.SQLiteTableObj),
}

//

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

		for _, obj := range generatorArr {
			if e := obj.Generator(newPath, &item); e != nil {
				return e
			}
		}
	}

	return nil
}
