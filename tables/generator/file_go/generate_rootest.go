package file_go

import (
	"bytes"
	"fmt"
	"microlog/tables/generator"
	"path/filepath"
)

// // // // // // // // // //

func generateRootTest(pathToDir string, pathsMap map[string]*generator.InfoTableObj) error {
	importArr := make([]string, 0)
	importArr = append(importArr, "testing")

	for path, _ := range pathsMap {
		path = fmt.Sprintf("microlog/tables/%s", filepath.Base(path))
		importArr = append(importArr, path)
	}

	// //

	var buf bytes.Buffer
	setHeaderGo(filepath.Base(pathToDir), &buf)

	if len(importArr) > 0 {
		buf.WriteString("import (\n")
		for _, line := range importArr {
			buf.WriteString(fmt.Sprintf("\t\"%s\"\n", line))
		}
		buf.WriteString(")\n")
	}

	setSeparator(&buf, 8)

	// //

	for path, table := range pathsMap {
		key := filepath.Base(path)
		path = fmt.Sprintf("microlog/tables/%s", key)

		buf.WriteString(fmt.Sprintf("func Test%s(t *testing.T) {\n", goNamespace(table.Name)))
		buf.WriteString(fmt.Sprintf(
			"\tif %s.Table != \"%s\" {t.Error(\"table name should be empty\")}\n",
			key, table.Name,
		))
		buf.WriteString("}\n\n")
	}

	// //

	return writeGoFile(filepath.Join(pathToDir, "table_generator_test.go"), buf.Bytes())
}
