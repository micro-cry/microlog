package generator_md

import (
	"bytes"
	_ "embed"
	"fmt"
	"microlog"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// // // // // // // // // //

//go:embed md-tables-struct.tmpl
var MdTablesFile string

type TableObj struct {
	Title          string
	ColumnsNameArr []string
	RowsArr        [][]string
}

type MdTablesObj struct {
	TablesArr []*TableObj
}

// //

func Generate(tablesArr []microlog.InfoTableObj, pathToFile string) error {
	data := new(MdTablesObj)

	for _, table := range tablesArr {
		tableBuf := new(TableObj)

		tableBuf.Title = table.Name
		tableBuf.RowsArr = [][]string{[]string{}, []string{}}

		for _, column := range table.Columns {
			tableBuf.ColumnsNameArr = append(tableBuf.ColumnsNameArr, column.Name)
			var row1Buf strings.Builder
			var row2Buf strings.Builder

			if column.Children == nil {
				if column.Length > 0 {
					row1Buf.WriteString(fmt.Sprintf("[%d]", column.Length))
				}
				row1Buf.WriteString(column.Type.String())
			} else {
				row1Buf.WriteString(fmt.Sprintf("%s.%s", column.Children.Table.Name, column.Children.Column.Name))
			}

			if column.Key != microlog.KeyNone {
				row2Buf.WriteString(fmt.Sprintf("_%s_", column.Key.String()))
			} else if column.Children != nil {
				row2Buf.WriteString(fmt.Sprintf("_*%s_", microlog.KeyIndex.String()))
			} else {
				row2Buf.WriteString("-")
			}

			tableBuf.RowsArr[0] = append(tableBuf.RowsArr[0], row1Buf.String())
			tableBuf.RowsArr[1] = append(tableBuf.RowsArr[1], row2Buf.String())
		}

		data.TablesArr = append(data.TablesArr, tableBuf)
	}

	// //

	return writeFileFromTemplate(pathToFile, MdTablesFile, data)
}

// //

func writeFileFromTemplate(pathToFile string, textTemplate string, dataTemplate any) error {
	fileName := filepath.Base(pathToFile)

	t, err := template.New(fileName).Parse(textTemplate)
	if err != nil {
		return fmt.Errorf("init template [%s]: %s", fileName, err.Error())
	}

	var buf bytes.Buffer
	err = t.Execute(&buf, dataTemplate)
	if err != nil {
		return fmt.Errorf("filling template [%s]: %s", fileName, err.Error())
	}

	file, err := os.OpenFile(pathToFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("open file [%s]: %s", fileName, err.Error())
	}
	defer file.Close()

	_, err = file.Write(buf.Bytes())
	if err != nil {
		return fmt.Errorf("write file [%s]: %s", fileName, err.Error())
	}

	return nil
}
