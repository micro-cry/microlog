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
	generatorArr = append(generatorArr, generateValues)
}

func generateValues(dirPath string, table *generator.InfoTableObj) error {
	data := generator_template.ValuesObj{
		PackageName:    filepath.Base(dirPath),
		TableName:      table.Name,
		TableConstName: "Table",
		MapName:        "NameToTypeMap",
	}

	// //

	for _, column := range table.Columns {
		data.ConstArr = append(data.ConstArr, fmt.Sprintf(
			"Name%s %s = \"%s\"",
			goNamespace(column.Name), TypeColumnName, column.Name,
		))
	}

	for _, column := range table.Columns {
		var strBuf strings.Builder

		strBuf.WriteString(fmt.Sprintf("Name%s: ", goNamespace(column.Name)))
		strBuf.WriteString("\"")

		if column.Children == nil {
			strBuf.WriteString(nameColumType(column.Length, column.Type))
		} else {
			strBuf.WriteString(nameColumType(column.Children.Column.Length, column.Children.Column.Type))
		}

		strBuf.WriteString("\",")

		data.MapArr = append(data.MapArr, strBuf.String())
	}

	// //

	return writeFileFromTemplate(filepath.Join(dirPath, "values.go"), generator_template.ValuesFile, data)
}
