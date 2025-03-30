package generator_template

import _ "embed"

// // // // // // // // // //

//go:embed struct.template
var StructFile string

type StructObj struct {
	PackageName    string
	ColumnNameType string
	ImportArr      []string
	GoTableName    string
	ObjArr         []string
	TableObjArr    []string
}
