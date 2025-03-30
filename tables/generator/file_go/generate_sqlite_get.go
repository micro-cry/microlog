package file_go

import (
	"microlog/tables/generator"
	"microlog/tables/generator/file_go/generator_template"
	"path/filepath"
)

// // // // // // // // // //

func init() {
	generatorArr = append(generatorArr, generateSQLiteGet)
}

func generateSQLiteGet(dirPath string, table *generator.InfoTableObj) error {
	data := generator_template.SQLiteGetObj{
		PackageName:   filepath.Base(dirPath),
		SQLiteObjName: "SQLiteObj",
	}

	// //

	return writeFileFromTemplate(filepath.Join(dirPath, "sqlite_get.go"), generator_template.SQLiteGetFile, data)
}
