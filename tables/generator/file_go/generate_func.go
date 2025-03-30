package file_go

import (
	"bytes"
	"fmt"
	"microlog/tables/generator"
	"path/filepath"
)

// // // // // // // // // //

func generateFunc(dirPath string, table *generator.InfoTableObj) error {
	var buf bytes.Buffer
	setHeaderGo(filepath.Base(dirPath), &buf)

	importArr := []string{
		"encoding/json",
		"microlog/tables",
	}

	if len(importArr) > 0 {
		buf.WriteString("import (\n")
		for _, line := range importArr {
			buf.WriteString(fmt.Sprintf("\t\"%s\"\n", line))
		}
		buf.WriteString(")\n")
	}

	setSeparator(&buf, 8)

	// //

	buf.WriteString("func (name " + TypeColumnName + ") String() string {\n")
	buf.WriteString("\treturn string(name)\n")
	buf.WriteString("}\n\n")

	buf.WriteString("func (name " + TypeColumnName + ") StringSQL() string {\n")
	buf.WriteString("\treturn \"`\"+Table+\".\"+string(name)+\"`\"\n")
	buf.WriteString("}\n\n")

	buf.WriteString("func (name " + TypeColumnName + ") Type() string {\n")
	buf.WriteString("\treturn NameToTypeMap[name]\n")
	buf.WriteString("}\n\n")

	buf.WriteString("func (name " + TypeColumnName + ") TableName() string {\n")
	buf.WriteString("\treturn Table\n")
	buf.WriteString("}\n\n")

	//

	setSeparator(&buf, 4)

	buf.WriteString("func (obj *")
	buf.WriteString(nameObj(table.Name))
	buf.WriteString(") JSON() ([]byte, error) {\n")
	buf.WriteString("\treturn json.Marshal(obj)\n}\n\n")

	buf.WriteString("func (obj *")
	buf.WriteString(nameObj(table.Name))
	buf.WriteString(") Children() tables.DataTableInterface {\n")
	buf.WriteString("\tobjTable := new(")
	buf.WriteString(nameTableObj(table.Name) + ")\n")
	//todo генератор по реальным
	buf.WriteString("\treturn objTable\n")
	buf.WriteString("}\n\n")

	//

	setSeparator(&buf, 2)

	buf.WriteString("func (objTable *")
	buf.WriteString(nameTableObj(table.Name))
	buf.WriteString(") JSON() ([]byte, error) {\n")
	buf.WriteString("\treturn json.Marshal(objTable)\n}\n\n")

	buf.WriteString("func (objTable *")
	buf.WriteString(nameTableObj(table.Name))
	buf.WriteString(") TableName() string {\n")
	buf.WriteString("\treturn Table\n}\n\n")

	buf.WriteString("func (objTable *")
	buf.WriteString(nameTableObj(table.Name))
	buf.WriteString(") TableColumns() map[tables.ColumnNameInterface]string {\n")
	buf.WriteString("\treturn NameToTypeMap\n}\n\n")

	buf.WriteString("func (objTable *")
	buf.WriteString(nameTableObj(table.Name))
	buf.WriteString(") Parent() tables.DataInterface {\n")
	buf.WriteString("\tobj := new(")
	buf.WriteString(nameObj(table.Name) + ")\n")
	//todo генератор по реальным
	buf.WriteString("\treturn obj\n")
	buf.WriteString("}\n\n")

	// //

	return writeGoFile(filepath.Join(dirPath, "func.go"), buf.Bytes())
}
