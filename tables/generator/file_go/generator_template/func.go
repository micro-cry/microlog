package generator_template

import _ "embed"

// // // // // // // // // //

//go:embed func.tmpl
var FuncFile string

type FuncObj struct {
	PackageName    string
	TableConstName string
	MapName        string

	ColumnTypeName   string
	DataObjName      string
	DataTableObjName string
	ParentComment    string

	ChildrenArr []string
	ParentArr   []string
}
