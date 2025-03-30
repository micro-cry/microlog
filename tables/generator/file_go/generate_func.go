package file_go

import (
	"fmt"
	"microlog/tables/generator"
	"microlog/tables/generator/file_go/generator_template"
	"path/filepath"
	"strings"
)

// // // // // // // // // //

func init() {
	generatorArr = append(generatorArr, generateFunc)
}

func generateFunc(dirPath string, table *generator.InfoTableObj) error {
	data := generator_template.FuncObj{
		PackageName:    filepath.Base(dirPath),
		TableConstName: "Table",
		MapName:        "NameToTypeMap",

		ColumnTypeName:   TypeColumnName,
		DataObjName:      nameObj(table.Name),
		DataTableObjName: nameTableObj(table.Name),
	}

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

	return writeFileFromTemplate(filepath.Join(dirPath, "func.go"), generator_template.FuncFile, data)
}
