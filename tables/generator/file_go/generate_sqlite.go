package file_go

import (
	"microlog/tables/generator"
	"microlog/tables/generator/file_go/generator_template"
	"path/filepath"
)

// // // // // // // // // //

func init() {
	generatorArr = append(generatorArr, generateSQLite)
}

func generateSQLite(dirPath string, table *generator.InfoTableObj) error {
	data := generator_template.SQLiteObj{
		PackageName:   filepath.Base(dirPath),
		SQLiteObjName: "SQLiteObj",
	}

	// //

	return writeFileFromTemplate(filepath.Join(dirPath, "sqlite.go"), generator_template.SQLiteFile, data)
}
