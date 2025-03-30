package generator_template

import (
	_ "embed"
	"fmt"
	"microlog/tables/generator"
	"path/filepath"
	"strings"
)

// // // // // // // // // //

//go:embed func.tmpl
var FuncFile string

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

func (data *FuncObj) Generator(dirPath string, table *generator.InfoTableObj) error {
	data.PackageName = filepath.Base(dirPath)
	data.TableConstName = "Table"
	data.MapName = "NameToTypeMap"

	data.ColumnTypeName = TypeColumnName
	data.DataObjName = nameObj(table.Name)
	data.DataTableObjName = nameTableObj(table.Name)

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

	return writeFileFromTemplate(filepath.Join(dirPath, "func.go"), FuncFile, data)
}
