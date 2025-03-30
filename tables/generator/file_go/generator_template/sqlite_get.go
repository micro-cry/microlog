package generator_template

import _ "embed"

// // // // // // // // // //

//go:embed sqlite_get.tmpl
var SQLiteGetFile string

type SQLiteGetObj struct {
	PackageName   string
	SQLiteObjName string
}
