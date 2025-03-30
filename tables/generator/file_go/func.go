package file_go

import (
	"bytes"
	"fmt"
	"go/format"
	"microlog/tables/generator"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"unicode"
)

// // // // // // // // // //

func clearOldDir(pathToDir string) error {
	items, err := os.ReadDir(pathToDir)
	if err != nil {
		return err
	}

	for _, item := range items {
		if item.IsDir() && len(item.Name()) > len(DirPrefix) {
			if item.Name()[:len(DirPrefix)] == DirPrefix {
				fullPath := filepath.Join(pathToDir, item.Name())
				err = os.RemoveAll(fullPath)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func createDir(pathToDir, dirName string) (string, error) {
	newPath := filepath.Join(pathToDir, DirPrefix+dirName)
	err := os.Mkdir(newPath, 0755)
	if err != nil {
		return "", err
	}

	return newPath, nil
}

func writeGoFile(pathToFile string, data []byte) error {
	file, err := os.OpenFile(pathToFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	formatted, err := format.Source(data)
	if err != nil {
		panic(err)
	}

	_, err = file.Write(formatted)
	return err
}

func writeFileFromTemplate(pathToFile string, textTemplate string, dataTemplate any) error {
	fileName := filepath.Base(pathToFile)

	t, err := template.New(pathToFile).Parse(textTemplate)
	if err != nil {
		return fmt.Errorf("init template [%s]: %s", fileName, err.Error())
	}

	var buf bytes.Buffer
	err = t.Execute(&buf, dataTemplate)
	if err != nil {
		return fmt.Errorf("filling template [%s]: %s", fileName, err.Error())
	}

	file, err := os.OpenFile(pathToFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("open file [%s]: %s", fileName, err.Error())
	}
	defer file.Close()

	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		return fmt.Errorf("format template [%s]: %s", fileName, err.Error())
	}

	_, err = file.Write(formatted)
	if err != nil {
		return fmt.Errorf("write file [%s]: %s", fileName, err.Error())
	}

	return nil
}

// //

func goNamespace(s string) string {
	if len(s) == 0 {
		return s
	}

	runes := []rune(s)
	first := unicode.ToUpper(runes[0])
	rest := strings.ToLower(string(runes[1:]))
	return string(first) + rest
}

//

func nameObj(tableName string) string {
	tableName = goNamespace(tableName)
	return tableName + "Obj"
}

func nameTableObj(tableName string) string {
	tableName = goNamespace(tableName)
	return tableName + "TableObj"
}

func nameColumType(l uint32, t generator.ColumType) string {
	switch t {

	case generator.ColumBool, generator.ColumByte, generator.ColumString:
		return t.String()

	case generator.ColumBytes:
		if l == 0 {
			return "[]byte"
		} else {
			return fmt.Sprintf("[%d]byte", l)
		}

	case generator.ColumDateTime:
		return "time.Time"
	}

	return "any"
}
