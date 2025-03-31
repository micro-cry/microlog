package generators

import (
	_ "embed"
	"fmt"
	"math/rand"
	"microlog"
	"microlog/generator_go"
	"path/filepath"
)

// // // // // // // // // //

type FuncTestObj struct {
	Global *microlog.GlobalDocInfoObj

	PackageName   string
	GoName        string
	TableName     string
	ColumnName    string
	ColumnNameSQL string

	rootDirName string
}

// //

func NewFuncTest(rootDirName string) generator_go.GeneratorInterface {
	obj := new(FuncTestObj)
	obj.rootDirName = rootDirName
	return obj
}

func (data *FuncTestObj) Generator(dirPath string, table *microlog.InfoTableObj) error {
	column := table.Columns[rand.Intn(len(table.Columns))].Name

	data.PackageName = filepath.Base(dirPath)
	data.Global = microlog.FileGoFuncTest.NewTemplate()

	data.GoName = fmt.Sprintf("%s%s", generator_go.ColumnNamePrefix, microlog.NameValGo(column))
	data.TableName = table.Name
	data.ColumnName = column
	data.ColumnNameSQL = "`" + table.Name + "." + column + "`"

	return writeFileFromTemplate(filepath.Join(dirPath, data.Global.NameGoFile()), data.Global.TemplateText(), data)
}
