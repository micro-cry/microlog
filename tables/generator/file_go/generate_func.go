package file_go

import (
	"fmt"
	"microlog/tables/generator"
	"path/filepath"
)

// // // // // // // // // //

func init() {
	generatorArr = append(generatorArr, generateFunc)
}

func generateFunc(dirPath string, table *generator.InfoTableObj) error {
	buf := newBuf(filepath.Base(dirPath))

	importArr := []string{
		"encoding/json",
		"microlog/tables",
	}

	buf.WriteImports(importArr)
	buf.WriteSeparator(8)

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

	buf.WriteSeparator(4)

	buf.WriteString("func (obj *")
	buf.WriteString(nameObj(table.Name))
	buf.WriteString(") JSON() ([]byte, error) {\n")
	buf.WriteString("\treturn json.Marshal(obj)\n}\n\n")

	buf.WriteString("func (obj *")
	buf.WriteString(nameObj(table.Name))
	buf.WriteString(") Children() tables.DataTableInterface {\n")
	buf.WriteString("\tobjTable := new(")
	buf.WriteString(nameTableObj(table.Name) + ")\n")
	for _, column := range table.Columns {
		if column.Children == nil {
			buf.WriteString(fmt.Sprintf(
				"\tobjTable.%s = obj.%s\n",
				goNamespace(column.Name), goNamespace(column.Name),
			))
		} else {
			buf.WriteString(fmt.Sprintf(
				"\tobjTable.%s = obj.%s.%s\n",
				goNamespace(column.Name), goNamespace(column.Name), goNamespace(column.Children.Column.Name),
			))
		}
	}
	buf.WriteString("\treturn objTable\n")
	buf.WriteString("}\n\n")

	//

	buf.WriteSeparator(2)

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

	buf.WriteString("//warning!!! \n//method does not create a complete structure, but only transfers those values that were in the original structure! \n//in case of nesting, there will be nil-values.\n")
	buf.WriteString("func (objTable *")
	buf.WriteString(nameTableObj(table.Name))
	buf.WriteString(") Parent() tables.DataInterface {\n")
	buf.WriteString("\tobj := new(")
	buf.WriteString(nameObj(table.Name) + ")\n")
	for _, column := range table.Columns {
		if column.Children == nil {
			buf.WriteString(fmt.Sprintf(
				"\tobjTable.%s = obj.%s\n",
				goNamespace(column.Name), goNamespace(column.Name),
			))
		}
	}
	buf.WriteString("\treturn obj\n")
	buf.WriteString("}\n\n")

	// //

	return writeGoFile(filepath.Join(dirPath, "func.go"), buf.Bytes())
}
