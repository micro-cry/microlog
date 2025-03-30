package file_go

import (
	"bytes"
	"fmt"
	"math/rand"
	"microlog/tables/generator"
	"path/filepath"
)

// // // // // // // // // //

func init() {
	generatorArr = append(generatorArr, generateFuncTest)
}

func generateFuncTest(dirPath string, table *generator.InfoTableObj) error {
	var buf bytes.Buffer
	setHeaderGo(filepath.Base(dirPath), &buf)

	buf.WriteString("import \"testing\"")

	setSeparator(&buf, 8)

	// //

	column := table.Columns[rand.Intn(len(table.Columns))].Name
	goName := fmt.Sprintf("Name%s", goNamespace(column))
	columnSQL := "`" + table.Name + "." + column + "`"

	//

	buf.WriteString("func TestString(t *testing.T) {\n")
	buf.WriteString(fmt.Sprintf(
		"if %s.String() != \"%s\" {t.Fatalf(\"%s\", \"%s\", %s.String())}\n",
		goName, column, "expected %q, got %q", column, goName,
	))
	buf.WriteString("}\n\n")

	buf.WriteString("func TestStringSQL(t *testing.T) {\n")
	buf.WriteString(fmt.Sprintf("expected := \"%s\"\n", columnSQL))
	buf.WriteString(fmt.Sprintf(
		"if %s.StringSQL() != expected {t.Fatalf(\"%s\", expected, %s.StringSQL())}\n",
		goName, "expected %q, got %q", goName,
	))
	buf.WriteString("}\n\n")

	buf.WriteString("func TestTableName(t *testing.T) {\n")
	buf.WriteString(fmt.Sprintf(
		"if %s.TableName() != \"%s\" {t.Fatalf(\"%s\", \"%s\", %s.TableName())}\n",
		goName, table.Name, "expected %q, got %q", table.Name, goName,
	))
	buf.WriteString("}\n\n")

	//

	setSeparator(&buf, 4)

	buf.WriteString("func BenchmarkString(b *testing.B) {\n")
	buf.WriteString(fmt.Sprintf(
		"for i := 0; i < b.N; i++ {_ = %s.String()}\n",
		goName,
	))
	buf.WriteString("}\n\n")

	buf.WriteString("func BenchmarkStringSQL(b *testing.B) {\n")
	buf.WriteString(fmt.Sprintf(
		"for i := 0; i < b.N; i++ {_ = %s.StringSQL()}\n",
		goName,
	))
	buf.WriteString("}\n\n")

	buf.WriteString("func BenchmarkType(b *testing.B) {\n")
	buf.WriteString(fmt.Sprintf(
		"for i := 0; i < b.N; i++ {_ = %s.Type()}\n",
		goName,
	))
	buf.WriteString("}\n\n")

	buf.WriteString("func BenchmarkTableName(b *testing.B) {\n")
	buf.WriteString(fmt.Sprintf(
		"for i := 0; i < b.N; i++ {_ = %s.TableName()}\n",
		goName,
	))
	buf.WriteString("}\n\n")

	// //

	return writeGoFile(filepath.Join(dirPath, "func_test.go"), buf.Bytes())
}
