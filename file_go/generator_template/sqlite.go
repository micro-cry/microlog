package generator_template

import (
	_ "embed"
	"microlog"
	"path/filepath"
)

// // // // // // // // // //

//go:embed sqlite.tmpl
var SQLiteFile string

type SQLiteObj struct {
	PackageName   string
	SQLiteObjName string
}

// //

func (data *SQLiteObj) Generator(dirPath string, table *microlog.InfoTableObj) error {
	data.PackageName = filepath.Base(dirPath)
	data.SQLiteObjName = SQLitePrefix + "Obj"

	return writeFileFromTemplate(filepath.Join(dirPath, "sqlite.go"), SQLiteFile, data)
}
