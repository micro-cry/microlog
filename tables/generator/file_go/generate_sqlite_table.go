package file_go

import (
	"bytes"
	"microlog/tables/generator"
	"path/filepath"
)

// // // // // // // // // //

func init() {
	generatorArr = append(generatorArr, generateSQLiteTable)
}

func generateSQLiteTable(dirPath string, table *generator.InfoTableObj) error {
	var buf bytes.Buffer
	setHeaderGo(filepath.Base(dirPath), &buf)

	importArr := []string{}

	setImports(&buf, importArr)
	setSeparator(&buf, 8)

	// //

	// //

	return writeGoFile(filepath.Join(dirPath, "sqlite_table.go"), buf.Bytes())
}
