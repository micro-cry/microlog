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
	Global        microlog.GlobalDocInfoObj
	SQLiteObjName string

	Data *microlog.TemplateStructObj
	Map  *microlog.TemplateMapObj
}

// //

func (data *SQLiteGetObj) Generator(dirPath string, table *microlog.InfoTableObj) error {
	data.Global.PackageName = filepath.Base(dirPath)
	data.Global.TemplatePath = "go-sqlite_get.tmpl"
	data.Global.GenerationTime = time.Now().Format(time.RFC3339)

	data.Global.Params = make(map[string]string)
	data.Global.Params["ver"] = "'" + microlog.GlobalVersion + "'"
	data.Global.Params["name"] = "'" + microlog.GlobalName + "'"
	data.Global.Params["commit_hash"] = "'" + microlog.GlobalHash[32:] + "'"
	data.Global.Params["commit_date"] = "'" + microlog.GlobalDateUpdate + "'"

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

	return writeFileFromTemplate(filepath.Join(dirPath, "sqlite_get.go"), microlog.FileGoSqliteGet.Data, data)
}
