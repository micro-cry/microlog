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
	data.TableConstName = generator_go.TableConstName
	data.MapName = generator_go.TableMapName

	data.ColumnTypeName = generator_go.TypeColumnName
	data.DataObjName = generator_go.NameValGo(table.Name) + "Obj"
	data.DataTableObjName = generator_go.NameValGo(table.Name) + "TableObj"
	data.ParentComment = ""

	data.ChildrenArr = make([]string, 0)
	data.ParentArr = make([]string, 0)

	// //

	for _, column := range table.Columns {
		var strBuf strings.Builder

		if column.Children == nil {
			strBuf.WriteString(fmt.Sprintf(
				"objTable.%s = obj.%s",
				generator_go.NameValGo(column.Name), generator_go.NameValGo(column.Name),
			))
		} else {
			data.ParentComment = "// warning!!!\n// method does not create a complete structure, but only transfers those values that were in the original structure!"

			strBuf.WriteString(fmt.Sprintf(
				"objTable.%s = obj.%s.%s",
				generator_go.NameValGo(column.Name), generator_go.NameValGo(column.Name), generator_go.NameValGo(column.Children.Column.Name),
			))
		}

		data.ChildrenArr = append(data.ChildrenArr, strBuf.String())
	}

	for _, column := range table.Columns {
		if column.Children == nil {
			data.ParentArr = append(data.ParentArr, fmt.Sprintf(
				"objTable.%s = obj.%s",
				generator_go.NameValGo(column.Name), generator_go.NameValGo(column.Name),
			))
		}
	}

	// //

	return writeFileFromTemplate(filepath.Join(dirPath, "func.go"), generator_go.FuncFile, data)
}
