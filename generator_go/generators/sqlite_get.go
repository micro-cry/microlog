package generators

import (
	_ "embed"
	"microlog"
	"microlog/generator_go"
	"path/filepath"
)

// // // // // // // // // //

type SQLiteGetObj struct {
	Global *microlog.GlobalDocInfoObj

	SQLiteObjName string
	PackageName   string

	Data *microlog.TemplateStructObj
	Map  *microlog.TemplateMapObj
}

// //

func (data *SQLiteGetObj) Generator(dirPath string, table *microlog.InfoTableObj) error {
	data.PackageName = filepath.Base(dirPath)
	data.Global = microlog.FileGoSqliteGet.NewTemplate()

	data.SQLiteObjName = generator_go.SQLitePrefix + "Obj"

	//

	data.Data = new(microlog.TemplateStructObj)
	data.Data.LinesArr = make([]*microlog.StructLineObj, 0)

	data.Data.NameStruct = "TestName"
	data.Data.CommentStruct = "test comment"

	data.Data.LinesArr = append(data.Data.LinesArr, &microlog.StructLineObj{
		Name:    "SomeName",
		Type:    "string",
		Reflect: "asasasa",
		Comment: "test comment",
	})

	//

	data.Map = new(microlog.TemplateMapObj)
	data.Map.ValuesArr = make([]*microlog.MapLineObj, 0)

	data.Map.NameMap = "TestName"
	data.Map.CommentMap = "test comment"
	data.Map.TypeKey = "int"
	data.Map.TypeValue = "string"

	data.Map.ValuesArr = append(data.Map.ValuesArr, &microlog.MapLineObj{
		Key:   "1",
		Value: "\"TestName\"",
	})

	//

	return writeFileFromTemplate(filepath.Join(dirPath, data.Global.NameGoFile()), data.Global.TemplateText(), data)
}
