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

	buf.WriteImports("microlog/tables")
	buf.WriteSeparator(8)

	// //

	buf.WriteLine("const (")
	buf.WritePadLine(1, fmt.Sprintf("Table = \"%s\"\n", table.Name))

	for _, column := range table.Columns {
		goName := fmt.Sprintf("Name%s", goNamespace(column.Name))
		buf.WritePadLine(1, goName, " ", TypeColumnName, " = \"", column.Name, "\"")
	}
	buf.WriteLine(")\n")

	buf.WriteLine("var NameToTypeMap = map[tables.ColumnNameInterface]string {")
	for _, column := range table.Columns {
		goName := fmt.Sprintf("Name%s", goNamespace(column.Name))
		buf.WritePadString(1, goName, ": \"")

		if column.Children == nil {
			buf.WriteString(nameColumType(column.Length, column.Type))
		} else {
			buf.WriteString(nameColumType(column.Children.Column.Length, column.Children.Column.Type))
		}
		buf.WriteString("\"", ",", "\n")
	}
	buf.WriteLine("}")

	// //

	return writeGoFile(filepath.Join(dirPath, "values.go"), buf.Bytes())
}
