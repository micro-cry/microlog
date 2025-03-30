package file_go

import (
	"microlog/tables/generator"
	"microlog/tables/generator/file_go/generator_template"
	"path/filepath"
)

// // // // // // // // // //

func init() {
	generatorArr = append(generatorArr, generateSQLiteOther)
}

func generateSQLiteOther(dirPath string, table *generator.InfoTableObj) error {
	data := generator_template.SQLiteOtherObj{
		PackageName:   filepath.Base(dirPath),
		SQLiteObjName: "SQLiteObj",
	}

	// //

	return data.Generator(dirPath, table)
}
