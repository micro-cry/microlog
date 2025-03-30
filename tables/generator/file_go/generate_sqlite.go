package file_go

import (
	"microlog/tables/generator"
	"path/filepath"
)

// // // // // // // // // //

const objSQLite = "SQLiteObj"

func init() {
	generatorArr = append(generatorArr, generateSQLite)
}

func generateSQLite(dirPath string, table *generator.InfoTableObj) error {
	buf := newBuf(filepath.Base(dirPath))

	importArr := []string{
		"context",
		"database/sql",
	}

	buf.WriteImports(importArr...)
	buf.WriteSeparator(8)

	// //

	buf.WriteString("type " + objSQLite + " struct {\n")
	buf.WriteString("\tdb *sql.DB\n")
	buf.WriteString("\tctx context.Context\n")
	buf.WriteString("}\n\n")

	buf.WriteString("func DriverSQLite(db *sql.DB, ctx context.Context) *" + objSQLite + " {\n")
	buf.WriteString("\tobj := new(" + objSQLite + ")\n")
	buf.WriteString("\tobj.db = db\n")
	buf.WriteString("\tobj.ctx = ctx\n")
	buf.WriteString("\treturn obj\n")
	buf.WriteString("}\n\n")

	//

	buf.WriteSeparator(2)

	buf.WriteString("func (obj *" + objSQLite + ") Ping() error {\n")
	buf.WriteString("\treturn obj.db.PingContext(obj.ctx)\n")
	buf.WriteString("}\n\n")

	// //

	return writeGoFile(filepath.Join(dirPath, "sqlite.go"), buf.Bytes())
}
