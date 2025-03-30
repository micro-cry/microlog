package generator_template

import _ "embed"

// // // // // // // // // //

//go:embed sqlite.tmpl
var SQLiteFile string

type SQLiteObj struct {
	PackageName   string
	SQLiteObjName string
}
