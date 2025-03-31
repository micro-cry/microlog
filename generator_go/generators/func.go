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

type FuncObj struct {
	PackageName    string
	TableConstName string
	MapName        string

	ColumnTypeName   string
	DataObjName      string
	DataTableObjName string
	ParentComment    string

	ChildrenArr []string
	ParentArr   []string
}

// //

func (data *FuncObj) Generator(dirPath string, table *microlog.InfoTableObj) error {
	data.PackageName = filepath.Base(dirPath)
	data.TableConstName = TableConstName
	data.MapName = TableMapName

	data.ColumnTypeName = TypeColumnName
	data.DataObjName = nameObj(table.Name)
	data.DataTableObjName = nameTableObj(table.Name)
	data.ParentComment = ""

	data.ChildrenArr = make([]string, 0)
	data.ParentArr = make([]string, 0)

	// //

	for _, column := range table.Columns {
		var strBuf strings.Builder

		if column.Children == nil {
			strBuf.WriteString(fmt.Sprintf(
				"objTable.%s = obj.%s",
				goNamespace(column.Name), goNamespace(column.Name),
			))
		} else {
			data.ParentComment = "// warning!!!\n// method does not create a complete structure, but only transfers those values that were in the original structure!"

			strBuf.WriteString(fmt.Sprintf(
				"objTable.%s = obj.%s.%s",
				goNamespace(column.Name), goNamespace(column.Name), goNamespace(column.Children.Column.Name),
			))
		}

		data.ChildrenArr = append(data.ChildrenArr, strBuf.String())
	}

	for _, column := range table.Columns {
		if column.Children == nil {
			data.ParentArr = append(data.ParentArr, fmt.Sprintf(
				"objTable.%s = obj.%s",
				goNamespace(column.Name), goNamespace(column.Name),
			))
		}
	}

	// //

	return writeFileFromTemplate(filepath.Join(dirPath, "func.go"), generator_go.FuncFile, data)
}
