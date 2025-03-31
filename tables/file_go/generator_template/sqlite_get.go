package generator_template

import (
	_ "embed"
	"microlog/tables"
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

func (data *SQLiteGetObj) Generator(dirPath string, table *tables.InfoTableObj) error {
	data.PackageName = filepath.Base(dirPath)
	data.SQLiteObjName = SQLitePrefix + "Obj"

	return writeFileFromTemplate(filepath.Join(dirPath, "sqlite_get.go"), SQLiteGetFile, data)
}
