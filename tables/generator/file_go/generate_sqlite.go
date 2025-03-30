package file_go

import (
	"bytes"
	"microlog/tables/generator"
	"path/filepath"
)

// // // // // // // // // //

const objSQLite = "SQLiteObj"

func init() {
	generatorArr = append(generatorArr, generateSQLite)
}

func generateSQLite(dirPath string, table *generator.InfoTableObj) error {
	var buf bytes.Buffer
	setHeaderGo(filepath.Base(dirPath), &buf)

	importArr := []string{
		"context",
		"database/sql",
	}

	setImports(&buf, importArr)
	setSeparator(&buf, 8)

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

	setSeparator(&buf, 2)

	buf.WriteString("func (obj *" + objSQLite + ") Ping() error {\n")
	buf.WriteString("\treturn obj.db.PingContext(obj.ctx)\n")
	buf.WriteString("}\n\n")

	// //

	return writeGoFile(filepath.Join(dirPath, "sqlite.go"), buf.Bytes())
}
