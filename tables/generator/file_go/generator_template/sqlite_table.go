package generator_template

import (
	_ "embed"
	"microlog/tables/generator"
	"path/filepath"
)

// // // // // // // // // //

//go:embed sqlite_table.tmpl
var SQLiteTableFile string

type SQLiteTableObj struct {
	PackageName   string
	SQLiteObjName string
}

// //

func (data *SQLiteTableObj) Generator(dirPath string, table *generator.InfoTableObj) error {
	return writeFileFromTemplate(filepath.Join(dirPath, "sqlite_table.go"), SQLiteTableFile, data)
}
