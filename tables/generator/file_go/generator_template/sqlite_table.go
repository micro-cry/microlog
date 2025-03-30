package generator_template

import _ "embed"

// // // // // // // // // //

//go:embed sqlite_table.tmpl
var SQLiteTableFile string

type SQLiteTableObj struct {
	PackageName   string
	SQLiteObjName string
}
