package generator_template

import (
	_ "embed"
	"microlog"
	"path/filepath"
)

// // // // // // // // // //

//go:embed sqlite_get.tmpl
var SQLiteGetFile string

type SQLiteGetObj struct {
	PackageName   string
	SQLiteObjName string

	Data *TemplateStructObj
}

// //

func (data *SQLiteGetObj) Generator(dirPath string, table *microlog.InfoTableObj) error {
	data.PackageName = filepath.Base(dirPath)
	data.SQLiteObjName = SQLitePrefix + "Obj"

	data.Data = new(TemplateStructObj)
	data.Data.LinesArr = make([]*StructLineObj, 0)

	data.Data.NameStruct = "TestName"
	data.Data.CommentStruct = "test comment"

	data.Data.LinesArr = append(data.Data.LinesArr, &StructLineObj{
		Name:    "SomeName",
		Type:    "string",
		Reflect: "asasasa",
		Comment: "test comment",
	})

	return writeFileFromTemplate(filepath.Join(dirPath, "sqlite_get.go"), SQLiteGetFile, data)
}
