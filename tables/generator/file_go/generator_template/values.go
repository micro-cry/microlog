package generator_template

import _ "embed"

// // // // // // // // // //

//go:embed values.tmpl
var ValuesFile string

type ValuesObj struct {
	PackageName    string
	TableName      string
	TableConstName string
	MapName        string

	ConstArr []string
	MapArr   []string
}
