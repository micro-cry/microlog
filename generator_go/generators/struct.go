package generators

import (
	_ "embed"
	"fmt"
	"microlog"
	"microlog/generator_go"
	"path/filepath"
)

// // // // // // // // // //

type StructObj struct {
	Global *microlog.GlobalDocInfoObj

	PackageName    string
	TypeColumnName string
	ImportArr      []string

	Obj      *microlog.TemplateStructObj
	TableObj *microlog.TemplateStructObj
}

// //

func (data *StructObj) Generator(dirPath string, table *microlog.InfoTableObj) error {
	data.PackageName = filepath.Base(dirPath)
	data.Global = microlog.FileGoStruct.NewTemplate()
	data.TypeColumnName = generator_go.TypeColumnName

	data.TableObj = new(microlog.TemplateStructObj)
	data.Obj = new(microlog.TemplateStructObj)

	data.Obj.NameStruct = microlog.NameValGo(table.Name)
	data.TableObj.NameStruct = microlog.NameValGo(table.Name)

	// //

	mapInc := map[microlog.ColumType]string{
		microlog.ColumDateTime: "time",
	}

	for _, column := range table.Columns {
		s, o := mapInc[column.Type]
		if o && s != "" {
			data.ImportArr = append(data.ImportArr, s)
			mapInc[column.Type] = ""
		}

		if column.Children != nil {
			data.ImportArr = append(data.ImportArr, fmt.Sprintf("microlog/tables/%s%s", generator_go.DirPrefix, column.Children.Table.Name))
		}
	}

	// //

	for _, column := range table.Columns {
		lineObj := new(microlog.StructLineObj)
		lineTableObj := new(microlog.StructLineObj)

		lineObj.Name = microlog.NameValGo(column.Name)
		lineTableObj.Name = microlog.NameValGo(column.Name)

		if column.Children == nil {
			lineObj.Type = column.TypeString()
			lineTableObj.Type = column.TypeString()
		} else {
			lineObj.Type = fmt.Sprintf(
				"*%s%s.%sObj\t",
				generator_go.DirPrefix,
				column.Children.Table.Name,
				microlog.NameValGo(column.Children.Table.Name),
			)
			lineTableObj.Type = column.Children.Column.TypeString()
		}

		lineObj.Reflect = fmt.Sprintf("json:\"%s\"", column.Name)

		if column.Key != microlog.KeyNone {
			lineTableObj.Comment = column.Key.String()
		} else if column.Children != nil {
			lineTableObj.Comment = "*" + microlog.KeyIndex.String()
		}

		data.Obj.LinesArr = append(data.Obj.LinesArr, lineObj)
		data.TableObj.LinesArr = append(data.TableObj.LinesArr, lineTableObj)
	}

	// //

	return writeFileFromTemplate(filepath.Join(dirPath, data.Global.NameGoFile()), data.Global.TemplateText(), data)
}
