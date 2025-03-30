package file_go

import (
	"microlog/tables/generator"
	"microlog/tables/generator/file_go/generator_template"
	"path/filepath"
)

// // // // // // // // // //

func init() {
	generatorArr = append(generatorArr, generateSQLiteTable)
}

func generateSQLiteTable(dirPath string, table *generator.InfoTableObj) error {
	data := generator_template.SQLiteTableObj{
		PackageName:   filepath.Base(dirPath),
		SQLiteObjName: "SQLiteObj",
	}

	// //

	return writeFileFromTemplate(filepath.Join(dirPath, "sqlite_table.go"), generator_template.SQLiteTableFile, data)
}
