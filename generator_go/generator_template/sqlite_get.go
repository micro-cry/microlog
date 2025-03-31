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
	Map  *TemplateMapObj
}

// //

func (data *SQLiteGetObj) Generator(dirPath string, table *microlog.InfoTableObj) error {
	data.PackageName = filepath.Base(dirPath)
	data.SQLiteObjName = SQLitePrefix + "Obj"

	//

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

	//

	data.Map = new(TemplateMapObj)
	data.Map.ValuesArr = make([]*MapLineObj, 0)

	data.Map.NameMap = "TestName"
	data.Map.CommentMap = "test comment"
	data.Map.TypeKey = "int"
	data.Map.TypeValue = "string"

	data.Map.ValuesArr = append(data.Map.ValuesArr, &MapLineObj{
		Key:   "1",
		Value: "\"TestName\"",
	})

	//

	return writeFileFromTemplate(filepath.Join(dirPath, "sqlite_get.go"), SQLiteGetFile, data)
}
