package generator_template

import (
	_ "embed"
	"microlog/tables/generator"
)

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

// //

func (data *StructObj) Generator(dirPath string, table *generator.InfoTableObj) error {}
