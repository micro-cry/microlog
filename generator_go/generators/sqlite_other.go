package generators

import (
	_ "embed"
	"microlog"
	"microlog/generator_go"
	"path/filepath"
)

// // // // // // // // // //

type SQLiteOtherObj struct {
	Global *microlog.GlobalDocInfoObj

	PackageName   string
	SQLiteObjName string

	rootDirName string
}

// //

func NewSQLiteOther(rootDirName string) generator_go.GeneratorInterface {
	obj := new(SQLiteOtherObj)
	obj.rootDirName = rootDirName
	return obj
}

func (data *SQLiteOtherObj) Generator(dirPath string, table *microlog.InfoTableObj) error {
	data.PackageName = filepath.Base(dirPath)
	data.Global = microlog.FileGoSqliteOther.NewTemplate()

	data.SQLiteObjName = generator_go.SQLitePrefix + "Obj"

	return writeFileFromTemplate(filepath.Join(dirPath, data.Global.NameGoFile()), data.Global.TemplateText(), data)
}
