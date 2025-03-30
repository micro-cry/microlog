package file_go

import (
	"microlog/tables/generator"
	"microlog/tables/generator/file_go/generator_template"
)

// // // // // // // // // //

func init() {
	generatorArr = append(generatorArr, generateSQLiteOther)
}

func generateSQLiteOther(dirPath string, table *generator.InfoTableObj) error {
	data := new(generator_template.SQLiteOtherObj)

	return data.Generator(dirPath, table)
}
