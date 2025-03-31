package generator_template

import (
	_ "embed"
	"fmt"
	"math/rand"
	"microlog/tables"
	"path/filepath"
)

// // // // // // // // // //

//go:embed functest.tmpl
var FuncTestFile string

type FuncTestObj struct {
	PackageName   string
	GoName        string
	TableName     string
	ColumnName    string
	ColumnNameSQL string
}

// //

func (data *FuncTestObj) Generator(dirPath string, table *tables.InfoTableObj) error {
	column := table.Columns[rand.Intn(len(table.Columns))].Name

	data.PackageName = filepath.Base(dirPath)
	data.GoName = fmt.Sprintf("%s%s", ColumnNamePrefix, goNamespace(column))
	data.TableName = table.Name
	data.ColumnName = column
	data.ColumnNameSQL = "`" + table.Name + "." + column + "`"

	return writeFileFromTemplate(filepath.Join(dirPath, "func_test.go"), FuncTestFile, data)
}
