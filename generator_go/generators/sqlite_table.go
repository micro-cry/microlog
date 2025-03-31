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

	rootDirName string
}

// //

func NewSQLiteTable(rootDirName string) generator_go.GeneratorInterface {
	obj := new(SQLiteTableObj)
	obj.rootDirName = rootDirName
	obj.Global = microlog.FileGoSqliteTable.NewTemplate()
	return obj
}

func (data *SQLiteTableObj) Generator(dirPath string, table *microlog.InfoTableObj) error {
	data.PackageName = filepath.Base(dirPath)
	data.SQLiteObjName = generator_go.SQLitePrefix + "Obj"

	return writeFileFromTemplate(filepath.Join(dirPath, data.Global.NameGoFile()), data.Global.TemplateText(), data)
}
