package file_go

import (
	"fmt"
	"math/rand"
	"microlog/tables/generator"
	"microlog/tables/generator/file_go/generator_template"
	"path/filepath"
)

// // // // // // // // // //

func init() {
	generatorArr = append(generatorArr, generateFuncTest)
}

func generateFuncTest(dirPath string, table *generator.InfoTableObj) error {
	column := table.Columns[rand.Intn(len(table.Columns))].Name

	data := generator_template.FuncTestObj{
		PackageName:   filepath.Base(dirPath),
		GoName:        fmt.Sprintf("Name%s", goNamespace(column)),
		TableName:     table.Name,
		ColumnName:    column,
		ColumnNameSQL: "`" + table.Name + "." + column + "`",
	}

	// //

	return writeFileFromTemplate(filepath.Join(dirPath, "func_test.go"), generator_template.FuncTestFile, data)
}
