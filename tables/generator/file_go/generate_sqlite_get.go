package file_go

import (
	"microlog/tables/generator"
	"microlog/tables/generator/file_go/generator_template"
)

// // // // // // // // // //

func init() {
	generatorArr = append(generatorArr, generateSQLiteGet)
}

func generateSQLiteGet(dirPath string, table *generator.InfoTableObj) error {
	data := new(generator_template.SQLiteGetObj)

	return data.Generator(dirPath, table)
}
