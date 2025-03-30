package generator_template

import (
	_ "embed"
	"microlog/tables/generator"
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

func (data *FuncTestObj) Generator(dirPath string, table *generator.InfoTableObj) error {
	return writeFileFromTemplate(filepath.Join(dirPath, "func_test.go"), FuncTestFile, data)
}
