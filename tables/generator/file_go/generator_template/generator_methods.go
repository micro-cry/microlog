package generator_template

import (
	"bytes"
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"unicode"
)

// // // // // // // // // //

func goNamespace(s string) string {
	if len(s) == 0 {
		return s
	}

	runes := []rune(s)
	first := unicode.ToUpper(runes[0])
	rest := strings.ToLower(string(runes[1:]))
	return string(first) + rest
}

// //

func writeFileFromTemplate(pathToFile string, textTemplate string, dataTemplate any) error {
	fileName := filepath.Base(pathToFile)

	t, err := template.New(fileName).Parse(textTemplate)
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
