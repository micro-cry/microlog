package generator_template

import _ "embed"

// // // // // // // // // //

//go:embed sqlite_other.tmpl
var SQLiteOtherFile string

type SQLiteOtherObj struct {
	PackageName   string
	SQLiteObjName string
}
