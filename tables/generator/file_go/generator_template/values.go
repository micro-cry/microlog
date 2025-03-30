package generator_template

import (
	_ "embed"
	"fmt"
	"microlog/tables/generator"
	"path/filepath"
	"strings"
)

// // // // // // // // // //

//go:embed values.tmpl
var ValuesFile string

type ValuesObj struct {
	PackageName    string
	TableName      string
	TableConstName string
	MapName        string

	ConstArr []string
	MapArr   []string
}

// //

func (data *ValuesObj) Generator(dirPath string, table *generator.InfoTableObj) error {
	data.PackageName = filepath.Base(dirPath)
	data.TableName = table.Name
	data.TableConstName = "Table"
	data.MapName = "NameToTypeMap"

	data.ConstArr = make([]string, 0)
	data.MapArr = make([]string, 0)

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

	return writeFileFromTemplate(filepath.Join(dirPath, "values.go"), ValuesFile, data)
}
