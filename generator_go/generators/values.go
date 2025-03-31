package generators

import (
	_ "embed"
	"microlog"
	"microlog/generator_go"
	"path/filepath"
)

// // // // // // // // // //

type ValuesObj struct {
	Global *microlog.GlobalDocInfoObj

	PackageName string
	TableName   string
	ImportArr   []string

	ConstArr []*microlog.StructLineObj
	Map      *microlog.TemplateMapObj
}

// //

func (data *ValuesObj) Generator(dirPath string, table *microlog.InfoTableObj) error {
	data = new(ValuesObj)

	data.PackageName = filepath.Base(dirPath)
	data.Global = microlog.FileGoValues.NewTemplate()

	data.TableName = table.Name

	data.Map = new(microlog.TemplateMapObj)
	data.Map.NameMap = generator_go.TableMapName
	data.Map.TypeKey = "microlog.ColumnNameInterface"
	data.Map.TypeValue = "string"

	data.ImportArr = append(data.ImportArr, "microlog")

	// //

	for _, column := range table.Columns {
		lineConst := new(microlog.StructLineObj)
		lineMap := new(microlog.MapLineObj)

		lineConst.Name = microlog.NameValGo(column.Name)
		lineMap.Key = microlog.NameValGo(column.Name)

		lineConst.Type = generator_go.TypeColumnName
		lineConst.Value = column.Name

		if column.Children == nil {
			lineMap.Value = column.TypeString()
		} else {
			lineMap.Value = column.Children.Column.TypeString()
		}

		data.ConstArr = append(data.ConstArr, lineConst)
		data.Map.ValuesArr = append(data.Map.ValuesArr, lineMap)
	}

	// //

	return writeFileFromTemplate(filepath.Join(dirPath, data.Global.NameGoFile()), data.Global.TemplateText(), data)
}
