package generator_template

import (
	_ "embed"
	"microlog/tables/generator"
)

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

// //

func (data *ValuesObj) Generator(dirPath string, table *generator.InfoTableObj) error {}
