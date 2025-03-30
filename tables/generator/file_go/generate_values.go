package file_go

import (
	"microlog/tables/generator"
	"microlog/tables/generator/file_go/generator_template"
	"path/filepath"
)

// // // // // // // // // //

func init() {
	generatorArr = append(generatorArr, generateValues)
}

func generateValues(dirPath string, table *generator.InfoTableObj) error {
	data := generator_template.ValuesObj{
		PackageName:    filepath.Base(dirPath),
		TableName:      table.Name,
		TableConstName: "Table",
		MapName:        "NameToTypeMap",
	}

	// //

	return data.Generator(dirPath, table)
}
