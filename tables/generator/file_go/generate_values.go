package file_go

import (
	"bytes"
	"fmt"
	"microlog/tables/generator"
	"path/filepath"
)

// // // // // // // // // //

func generateValues(dirPath string, table *generator.InfoTableObj) error {
	var buf bytes.Buffer
	setHeaderGo(filepath.Base(dirPath), &buf)

	setSeparator(&buf, 8)

	// //

	buf.WriteString("const (\n")
	buf.WriteString(fmt.Sprintf("\tTable = \"%s\"\n\n", table.Name))

	for _, column := range table.Columns {
		goName := fmt.Sprintf("Name%s", goNamespace(column.Name))
		buf.WriteString(fmt.Sprintf("\t%s %s = \"%s\"\n", goName, TypeColumnName, column.Name))
	}
	buf.WriteString(")\n\n")

	buf.WriteString("var NameToTypeMap = map[" + TypeColumnName + "]string {\n")
	for _, column := range table.Columns {
		goName := fmt.Sprintf("Name%s", goNamespace(column.Name))
		buf.WriteString(fmt.Sprintf("\t%s: \"", goName))

		if column.Children == nil {
			setColumTypeToString(&buf, column.Length, column.Type)
		} else {
			setColumTypeToString(&buf, column.Children.Column.Length, column.Children.Column.Type)
		}
		buf.WriteString("\",\n")
	}
	buf.WriteString("}\n")

	// //

	return writeGoFile(filepath.Join(dirPath, "values.go"), buf.Bytes())
}
