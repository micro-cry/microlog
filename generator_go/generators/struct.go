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

	PackageName string
	ImportArr   []string

	Obj      *microlog.TemplateStructObj
	TableObj *microlog.TemplateStructObj
}

// //

func (data *StructObj) Generator(dirPath string, table *microlog.InfoTableObj) error {
	data.PackageName = filepath.Base(dirPath)
	data.Global = microlog.FileGoStruct.NewTemplate()

	data.ImportArr = make([]string, 0)
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
		lineBuf := new(microlog.StructLineObj)

		lineBuf.Name = microlog.NameValGo(column.Name)

		if column.Children == nil {
			lineBuf.Type = column.TypeString()
		} else {
			lineBuf.Type = fmt.Sprintf(
				"*%s%s.%sObj\t",
				generator_go.DirPrefix,
				column.Children.Table.Name,
				microlog.NameValGo(column.Children.Table.Name),
			)
		}

		lineBuf.Reflect = fmt.Sprintf("json:\"%s\"", column.Name)

		data.Obj.LinesArr = append(data.Obj.LinesArr, lineBuf)
	}

	for _, column := range table.Columns {
		lineBuf := new(microlog.StructLineObj)

		lineBuf.Name = microlog.NameValGo(column.Name)

		if column.Children == nil {
			lineBuf.Type = column.TypeString()

		} else {
			lineBuf.Type = column.Children.Column.TypeString()
		}

		if column.Key != microlog.KeyNone {
			lineBuf.Comment = column.Key.String()
		} else if column.Children != nil {
			lineBuf.Comment = "*" + microlog.KeyIndex.String()
		}

		data.TableObj.LinesArr = append(data.TableObj.LinesArr, lineBuf)
	}

	// //

	return writeFileFromTemplate(filepath.Join(dirPath, data.Global.NameGoFile()), data.Global.TemplateText(), data)
}
