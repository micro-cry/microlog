package file_go

import (
	"bytes"
	"fmt"
	"microlog/tables/generator"
	"path/filepath"
)

// // // // // // // // // //

func init() {
	generatorArr = append(generatorArr, generateSQLiteGet)
}

func generateSQLiteGet(dirPath string, table *generator.InfoTableObj) error {
	var buf bytes.Buffer
	setHeaderGo(filepath.Base(dirPath), &buf)

	importArr := []string{}

	if len(importArr) > 0 {
		buf.WriteString("import (\n")
		for _, line := range importArr {
			buf.WriteString(fmt.Sprintf("\t\"%s\"\n", line))
		}
		buf.WriteString(")\n")
	}

	setSeparator(&buf, 8)

	// //

	// //

	return writeGoFile(filepath.Join(dirPath, "sqlite_get.go"), buf.Bytes())
}
