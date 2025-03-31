package generators

import (
	_ "embed"
	"microlog"
	"microlog/generator_go"
	"path/filepath"
	"time"
)

// // // // // // // // // //

type SQLiteGetObj struct {
	Global        GlobalDocInfoObj
	SQLiteObjName string

	Data *TemplateStructObj
	Map  *TemplateMapObj
}

// //

func (data *SQLiteGetObj) Generator(dirPath string, table *microlog.InfoTableObj) error {
	data.Global.PackageName = filepath.Base(dirPath)
	data.Global.TemplatePath = "sqlite_get.tmpl"
	data.Global.GenerationTime = time.Now().Format(time.RFC3339)

	data.Global.Params = make(map[string]string)
	data.Global.Params["ver"] = "'" + microlog.GlobalVersion + "'"
	data.Global.Params["name"] = "'" + microlog.GlobalName + "'"
	data.Global.Params["commit_hash"] = "'" + microlog.GlobalHash[32:] + "'"
	data.Global.Params["commit_date"] = "'" + microlog.GlobalDateUpdate + "'"

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

	return writeFileFromTemplate(filepath.Join(dirPath, "sqlite_get.go"), generator_go.SQLiteGetFile, data)
}
