package file_go

import (
	"microlog/tables/generator"
	"microlog/tables/generator/file_go/generator_template"
)

// // // // // // // // // //

func init() {
	generatorArr = append(generatorArr, generateStruct)
}

func generateStruct(dirPath string, table *generator.InfoTableObj) error {
	data := new(generator_template.StructObj)

	return data.Generator(dirPath, table)
}
