package file_go

import (
	"microlog/tables/generator"
	"microlog/tables/generator/file_go/generator_template"
)

// // // // // // // // // //

func init() {
	generatorArr = append(generatorArr, generateValues)
}

func generateValues(dirPath string, table *generator.InfoTableObj) error {
	data := new(generator_template.ValuesObj)

	return data.Generator(dirPath, table)
}
