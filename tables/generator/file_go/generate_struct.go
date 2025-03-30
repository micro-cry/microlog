package file_go

import (
	"bytes"
	"fmt"
	"microlog/tables/generator"
	"path/filepath"
)

// // // // // // // // // //

func generateStruct(dirPath string, table *generator.InfoTableObj) error {
	importArr := make([]string, 0)
	mapInc := map[generator.ColumType]string{
		generator.ColumDateTime: "time",
	}

	for _, column := range table.Columns {
		s, o := mapInc[column.Type]
		if o && s != "" {
			importArr = append(importArr, s)
			mapInc[column.Type] = ""
		}

		if column.Children != nil {
			importArr = append(importArr, fmt.Sprintf("microlog/tables/%s%s", DirPrefix, column.Children.Table.Name))
		}
	}

	// //

	var buf bytes.Buffer
	setHeaderGo(filepath.Base(dirPath), &buf)

	if len(importArr) > 0 {
		buf.WriteString("import (\n")
		for _, line := range importArr {
			buf.WriteString(fmt.Sprintf("\t\"%s\"\n", line))
		}
		buf.WriteString(")\n")
	}

	setSeparator(&buf, 8)

	// //

	buf.WriteString("type " + TypeColumnName + " string\n\n")

	buf.WriteString("type ")
	buf.WriteString(nameObj(table.Name))
	buf.WriteString(" struct {\n")

	for _, column := range table.Columns {
		buf.WriteString("\t")

		buf.WriteString(goNamespace(column.Name))
		buf.WriteString("\t")

		if column.Children == nil {
			setColumTypeToString(&buf, column.Length, column.Type)
			buf.WriteString("\t")

		} else {
			buf.WriteString(fmt.Sprintf(
				"*%s%s.%sObj\t",
				DirPrefix,
				column.Children.Table.Name,
				goNamespace(column.Children.Table.Name),
			))
		}

		buf.WriteString(fmt.Sprintf("`json:\"%s\"`\t", column.Name))

		buf.WriteString("\n")
	}

	buf.WriteString("}\n")

	// //

	buf.WriteString("type ")
	buf.WriteString(nameTableObj(table.Name))
	buf.WriteString(" struct {\n")

	for _, column := range table.Columns {
		buf.WriteString("\t")

		buf.WriteString(goNamespace(column.Name))
		buf.WriteString("\t")

		if column.Children == nil {
			setColumTypeToString(&buf, column.Length, column.Type)
			buf.WriteString("\t")

		} else {
			setColumTypeToString(&buf, column.Children.Column.Length, column.Children.Column.Type)
			buf.WriteString("\t")
		}

		if column.Key != generator.KeyNone {
			buf.WriteString(fmt.Sprintf("//%s", column.Key.String()))
		}

		buf.WriteString("\n")
	}

	buf.WriteString("}\n")

	// //

	return writeGoFile(filepath.Join(dirPath, "struct.go"), buf.Bytes())
}
