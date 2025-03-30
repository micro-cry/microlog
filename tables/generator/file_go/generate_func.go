package file_go

import (
	"microlog/tables/generator"
	"microlog/tables/generator/file_go/generator_template"
	"path/filepath"
)

// // // // // // // // // //

func init() {
	generatorArr = append(generatorArr, generateFunc)
}

func generateFunc(dirPath string, table *generator.InfoTableObj) error {
	data := generator_template.FuncObj{
		PackageName:    filepath.Base(dirPath),
		TableConstName: "Table",
		MapName:        "NameToTypeMap",

		ColumnTypeName:   TypeColumnName,
		DataObjName:      nameObj(table.Name),
		DataTableObjName: nameTableObj(table.Name),
	}

	return data.Generator(dirPath, table)
}
