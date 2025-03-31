package generators

import (
	_ "embed"
	"microlog"
	"microlog/generator_go"
	"path/filepath"
)

// // // // // // // // // //

type SQLiteTableObj struct {
	Global *microlog.GlobalDocInfoObj

	PackageName   string
	SQLiteObjName string
}

// //

func (data *SQLiteTableObj) Generator(dirPath string, table *microlog.InfoTableObj) error {
	data.PackageName = filepath.Base(dirPath)
	data.Global = microlog.FileGoSqliteTable.NewTemplate()

	data.SQLiteObjName = generator_go.SQLitePrefix + "Obj"

	return writeFileFromTemplate(filepath.Join(dirPath, data.Global.NameGoFile()), data.Global.TemplateText(), data)
}
