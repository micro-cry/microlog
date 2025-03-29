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

	setSeparator(&buf, 16)

	// //

	// //

	return writeGoFile(filepath.Join(dirPath, "func.go"), buf.Bytes())
}
