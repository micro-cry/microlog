package file_go

import (
	"bytes"
	"microlog/tables/generator"
	"path/filepath"
)

// // // // // // // // // //

func generateFunc(dirPath string, table *generator.InfoTableObj) error {
	var buf bytes.Buffer
	setHeaderGo(filepath.Base(dirPath), &buf)

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

	// //

	return writeGoFile(filepath.Join(dirPath, "func.go"), buf.Bytes())
}
