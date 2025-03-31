package generators

import (
	_ "embed"
	"microlog"
	"microlog/generator_go"
	"path/filepath"
)

// // // // // // // // // //

type SQLiteObj struct {
	PackageName   string
	SQLiteObjName string
}

// //

func (data *SQLiteObj) Generator(dirPath string, table *microlog.InfoTableObj) error {
	data.PackageName = filepath.Base(dirPath)
	data.SQLiteObjName = generator_go.SQLitePrefix + "Obj"

	return writeFileFromTemplate(filepath.Join(dirPath, "sqlite.go"), generator_go.SQLiteFile, data)
}
