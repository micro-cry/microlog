package generators

import (
	_ "embed"
	"microlog"
	"microlog/generator_go"
	"path/filepath"
)

// // // // // // // // // //

type SQLiteOtherObj struct {
	PackageName   string
	SQLiteObjName string
}

// //

func (data *SQLiteOtherObj) Generator(dirPath string, table *microlog.InfoTableObj) error {
	data.PackageName = filepath.Base(dirPath)
	data.SQLiteObjName = generator_go.SQLitePrefix + "Obj"

	return writeFileFromTemplate(filepath.Join(dirPath, "sqlite_other.go"), microlog.FileGoSqliteOther.Data, data)
}
