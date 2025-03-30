package file_go

import (
	"fmt"
	"microlog/tables/generator"
	"path/filepath"
)

// // // // // // // // // //

func init() {
	generatorArr = append(generatorArr, generateValues)
}

func generateValues(dirPath string, table *generator.InfoTableObj) error {
	buf := newBuf(filepath.Base(dirPath))

	importArr := []string{
		"microlog/tables",
	}

	buf.WriteImports(importArr)
	buf.WriteSeparator(8)

	// //

	buf.WriteString("const (\n")
	buf.WriteString(fmt.Sprintf("\tTable = \"%s\"\n\n", table.Name))

	for _, column := range table.Columns {
		goName := fmt.Sprintf("Name%s", goNamespace(column.Name))
		buf.WriteString(fmt.Sprintf("\t%s %s = \"%s\"\n", goName, TypeColumnName, column.Name))
	}
	buf.WriteString(")\n\n")

	buf.WriteString("var NameToTypeMap = map[tables.ColumnNameInterface]string {\n")
	for _, column := range table.Columns {
		goName := fmt.Sprintf("Name%s", goNamespace(column.Name))
		buf.WriteString(fmt.Sprintf("\t%s: \"", goName))

		if column.Children == nil {
			buf.WriteString(nameColumType(column.Length, column.Type))
		} else {
			buf.WriteString(nameColumType(column.Children.Column.Length, column.Children.Column.Type))
		}
		buf.WriteString("\",\n")
	}
	buf.WriteString("}\n")

	// //

	return writeGoFile(filepath.Join(dirPath, "values.go"), buf.Bytes())
}
