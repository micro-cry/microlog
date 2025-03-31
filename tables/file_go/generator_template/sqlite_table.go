package generator_template

import (
	_ "embed"
	"microlog/tables"
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

func (data *SQLiteTableObj) Generator(dirPath string, table *tables.InfoTableObj) error {
	data.PackageName = filepath.Base(dirPath)
	data.SQLiteObjName = SQLitePrefix + "Obj"

	return writeFileFromTemplate(filepath.Join(dirPath, "sqlite_table.go"), SQLiteTableFile, data)
}
