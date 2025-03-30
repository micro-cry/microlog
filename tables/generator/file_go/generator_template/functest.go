package generator_template

import _ "embed"

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
