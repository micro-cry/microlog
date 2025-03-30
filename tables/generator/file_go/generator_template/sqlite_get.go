package generator_template

import (
	_ "embed"
	"microlog/tables/generator"
	"path/filepath"
)

// // // // // // // // // //

//go:embed sqlite_get.tmpl
var SQLiteGetFile string

type SQLiteGetObj struct {
	PackageName   string
	SQLiteObjName string
}

// //

func (data *SQLiteGetObj) Generator(dirPath string, table *generator.InfoTableObj) error {
	return writeFileFromTemplate(filepath.Join(dirPath, "sqlite_get.go"), SQLiteGetFile, data)
}
