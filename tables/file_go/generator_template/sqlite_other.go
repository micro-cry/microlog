package generator_template

import (
	_ "embed"
	"microlog/tables"
	"path/filepath"
)

// // // // // // // // // //

//go:embed sqlite_other.tmpl
var SQLiteOtherFile string

type SQLiteOtherObj struct {
	PackageName   string
	SQLiteObjName string
}

// //

func (data *SQLiteOtherObj) Generator(dirPath string, table *tables.InfoTableObj) error {
	data.PackageName = filepath.Base(dirPath)
	data.SQLiteObjName = SQLitePrefix + "Obj"

	return writeFileFromTemplate(filepath.Join(dirPath, "sqlite_other.go"), SQLiteOtherFile, data)
}
