package file_go

import (
	"bytes"
	"fmt"
	"math/rand"
	"microlog/tables/generator"
	"microlog/tables/generator/file_go/generator_template"
	"path/filepath"
	"text/template"
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

	t, err := template.New(DirPrefix + table.Name).Parse(generator_template.FuncTestFile)
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	err = t.Execute(&buf, data)
	if err != nil {
		return err
	}

	return writeGoFile(filepath.Join(dirPath, "func_test.go"), buf.Bytes())
}
