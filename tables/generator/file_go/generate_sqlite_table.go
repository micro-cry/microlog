package file_go

import (
	"microlog/tables/generator"
	"path/filepath"
)

// // // // // // // // // //

func init() {
	generatorArr = append(generatorArr, generateSQLiteTable)
}

func generateSQLiteTable(dirPath string, table *generator.InfoTableObj) error {
	buf := newBuf(filepath.Base(dirPath))

	importArr := []string{}

	buf.WriteImports(importArr)
	buf.WriteSeparator(8)

	// //

	// //

	return writeGoFile(filepath.Join(dirPath, "sqlite_table.go"), buf.Bytes())
}
