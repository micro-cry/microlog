package generator_md

import (
	"bytes"
	"fmt"
	"microlog"
	"os"
)

// // // // // // // // // //

func Generate(tablesArr []microlog.InfoTableObj, pathToFile string) error {
	var buf bytes.Buffer
	buf.WriteString("# The overall structure of the tables\n")
	buf.WriteString("This file is generated automatically\n\n")
	buf.WriteString("---\n\n")

	// //

	for _, table := range tablesArr {
		buf.WriteString(fmt.Sprintf("## %s\n\n", table.Name))
		lineBuf := []string{"|", "|", "|", "|"}

		for _, column := range table.Columns {
			lineBuf[0] += fmt.Sprintf(" %s |", column.Name)
			lineBuf[1] += "--|"

			lineBuf[2] += " "
			if column.Children == nil {

				if column.Length > 0 {
					lineBuf[2] += fmt.Sprintf("[%d]", column.Length)
				}
				lineBuf[2] += column.Type.String()

			} else {
				lineBuf[2] += fmt.Sprintf("%s.%s", column.Children.Table.Name, column.Children.Column.Name)
			}
			lineBuf[2] += " |"

			if column.Key != microlog.KeyNone {
				lineBuf[3] += fmt.Sprintf(" _%s_ |", column.Key.String())
			} else if column.Children != nil {
				lineBuf[3] += fmt.Sprintf(" _*%s_ |", microlog.KeyIndex.String())
			} else {
				lineBuf[3] += " - |"
			}

		}

		for _, line := range lineBuf {
			buf.WriteString(line + "\n")
		}
		buf.WriteString("\n\n")
	}

	// //

	file, err := os.OpenFile(pathToFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(buf.Bytes())
	return err
}
