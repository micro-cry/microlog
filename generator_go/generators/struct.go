package generators

import (
	_ "embed"
	"fmt"
	"microlog"
	"microlog/generator_go"
	"path/filepath"
	"strings"
)

// // // // // // // // // //

type StructObj struct {
	PackageName    string
	GoTableName    string
	ColumnNameType string

	ImportArr []string

	ObjArr      []string
	TableObjArr []string
}

// //

func (data *StructObj) Generator(dirPath string, table *microlog.InfoTableObj) error {
	data.PackageName = filepath.Base(dirPath)
	data.ColumnNameType = generator_go.TypeColumnName
	data.GoTableName = microlog.NameValGo(table.Name)

	data.ImportArr = make([]string, 0)
	data.ObjArr = make([]string, 0)
	data.TableObjArr = make([]string, 0)

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
		var strBuf strings.Builder

		strBuf.WriteString(microlog.NameValGo(column.Name))
		strBuf.WriteString("\t")

		if column.Children == nil {
			strBuf.WriteString(column.TypeString())
			strBuf.WriteString("\t")

		} else {
			strBuf.WriteString(fmt.Sprintf(
				"*%s%s.%sObj\t",
				generator_go.DirPrefix,
				column.Children.Table.Name,
				microlog.NameValGo(column.Children.Table.Name),
			))
		}

		strBuf.WriteString(fmt.Sprintf("`json:\"%s\"`\t", column.Name))

		data.ObjArr = append(data.ObjArr, strBuf.String())
	}

	for _, column := range table.Columns {
		var strBuf strings.Builder

		strBuf.WriteString(microlog.NameValGo(column.Name))
		strBuf.WriteString("\t")

		if column.Children == nil {
			strBuf.WriteString(column.TypeString())
			strBuf.WriteString("\t")

		} else {
			strBuf.WriteString(column.Children.Column.TypeString())
			strBuf.WriteString("\t")
		}

		if column.Key != microlog.KeyNone {
			strBuf.WriteString(fmt.Sprintf("//*%s", column.Key.String()))
		} else if column.Children != nil {
			strBuf.WriteString(fmt.Sprintf("//%s", microlog.KeyIndex.String()))
		}

		data.TableObjArr = append(data.TableObjArr, strBuf.String())
	}

	// //

	return writeFileFromTemplate(filepath.Join(dirPath, "struct.go"), microlog.FileGoStruct.Data, data)
}
