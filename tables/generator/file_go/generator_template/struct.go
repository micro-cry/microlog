package generator_template

import _ "embed"

// // // // // // // // // //

//go:embed struct.tmpl
var StructFile string

type StructObj struct {
	PackageName    string
	GoTableName    string
	ColumnNameType string

	ImportArr []string

	ObjArr      []string
	TableObjArr []string
}
