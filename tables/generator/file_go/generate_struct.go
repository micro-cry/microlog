package file_go

import (
	"bytes"
	"fmt"
	"microlog/tables/generator"
	"microlog/tables/generator/file_go/generator_template"
	"path/filepath"
	"strings"
	"text/template"
)

// // // // // // // // // //

func init() {
	generatorArr = append(generatorArr, generateStruct)
}

func generateStruct(dirPath string, table *generator.InfoTableObj) error {
	data := generator_template.StructObj{
		PackageName:    filepath.Base(dirPath),
		ColumnNameType: TypeColumnName,
		GoTableName:    goNamespace(table.Name),
	}

	mapInc := map[generator.ColumType]string{
		generator.ColumDateTime: "time",
	}

	for _, column := range table.Columns {
		s, o := mapInc[column.Type]
		if o && s != "" {
			data.ImportArr = append(data.ImportArr, s)
			mapInc[column.Type] = ""
		}

		if column.Children != nil {
			data.ImportArr = append(data.ImportArr, fmt.Sprintf("microlog/tables/%s%s", DirPrefix, column.Children.Table.Name))
		}
	}

	// //

	for _, column := range table.Columns {
		var strBuf strings.Builder

		strBuf.WriteString(goNamespace(column.Name))
		strBuf.WriteString("\t")

		if column.Children == nil {
			strBuf.WriteString(nameColumType(column.Length, column.Type))
			strBuf.WriteString("\t")

		} else {
			strBuf.WriteString(fmt.Sprintf(
				"*%s%s.%sObj\t",
				DirPrefix,
				column.Children.Table.Name,
				goNamespace(column.Children.Table.Name),
			))
		}

		strBuf.WriteString(fmt.Sprintf("`json:\"%s\"`\t", column.Name))

		data.ObjArr = append(data.ObjArr, strBuf.String())
	}

	for _, column := range table.Columns {
		var strBuf strings.Builder

		strBuf.WriteString(goNamespace(column.Name))
		strBuf.WriteString("\t")

		if column.Children == nil {
			strBuf.WriteString(nameColumType(column.Length, column.Type))
			strBuf.WriteString("\t")

		} else {
			strBuf.WriteString(nameColumType(column.Children.Column.Length, column.Children.Column.Type))
			strBuf.WriteString("\t")
		}

		if column.Key != generator.KeyNone {
			strBuf.WriteString(fmt.Sprintf("//*%s", column.Key.String()))
		} else if column.Children != nil {
			strBuf.WriteString(fmt.Sprintf("//%s", generator.KeyIndex.String()))
		}

		data.TableObjArr = append(data.TableObjArr, strBuf.String())
	}

	// //

	t, err := template.New(DirPrefix + table.Name).Parse(generator_template.StructFile)
	if err != nil {
		return fmt.Errorf("generate struct template error: %v", err)
	}

	var buf bytes.Buffer
	err = t.Execute(&buf, data)
	if err != nil {
		return err
	}

	return writeFileFromTemplate(filepath.Join(dirPath, "struct.go"), generator_template.StructFile, data)
}
