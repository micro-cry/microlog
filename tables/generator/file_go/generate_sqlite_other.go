package file_go

import (
	"microlog/tables/generator"
	"path/filepath"
)

// // // // // // // // // //

func init() {
	generatorArr = append(generatorArr, generateSQLiteOther)
}

func generateSQLiteOther(dirPath string, table *generator.InfoTableObj) error {
	buf := newBuf(filepath.Base(dirPath))

	importArr := []string{}

	buf.WriteImports(importArr)
	buf.WriteSeparator(8)

	// //

	// //

	return writeGoFile(filepath.Join(dirPath, "sqlite_other.go"), buf.Bytes())
}
