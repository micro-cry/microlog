package parser

import (
	"bytes"
	"fmt"
	"gopkg.in/yaml.v3"
	"microlog"
	"os"
	"path/filepath"
	"strings"
)

// // // // // // // // // //

func scanDir(pathToDir string) []string {
	files, err := os.ReadDir(pathToDir)
	if err != nil {
		return nil
	}

	var filesArr []string

	for _, file := range files {
		if !file.IsDir() {
			if filepath.Ext(file.Name()) == ".yml" {
				filesArr = append(filesArr, file.Name())
			}
		}
	}

	return filesArr
}

func readFile(pathToFile string) ([]*ColumObj, error) {
	data, err := os.ReadFile(pathToFile)
	if err != nil {
		return nil, err
	}

	var rawData map[string]map[string]interface{}
	err = yaml.Unmarshal(data, &rawData)
	if err != nil {
		return nil, err
	}

	bufMap := make(map[string]*ColumObj)
	for key, props := range rawData {
		col := new(ColumObj)
		col.Name = key

		if l, ok := props["len"]; ok {
			switch v := l.(type) {
			case int:
				col.Len = uint32(v)
			case int64:
				col.Len = uint32(v)
			case float64:
				col.Len = uint32(v)
			default:
				fmt.Printf("unknown type for Len field in the record %s\n", key)
			}
		}

		if t, ok := props["type"]; ok {
			if str, ok := t.(string); ok {
				col.Type = str
			}
		}

		if k, ok := props["key"]; ok {
			if str, ok := k.(string); ok {
				col.Key = str
			}
		}

		if c, ok := props["children"]; ok {
			if str, ok := c.(string); ok {
				col.Children = str
			}
		}

		if _, ok := bufMap[key]; ok {
			return nil, fmt.Errorf("duplicate colum %s", key)
		}

		bufMap[key] = col
	}

	// //

	var keys []string
	lines := bytes.Split(data, []byte("\n"))
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		trimmed := strings.TrimSpace(string(line))
		if len(trimmed) > 0 && line[0] != ' ' && strings.HasSuffix(trimmed, ":") {
			key := strings.TrimSuffix(trimmed, ":")
			keys = append(keys, key)
		}
	}

	// //

	bufArr := make([]*ColumObj, 0)
	for _, key := range keys {
		bufArr = append(bufArr, bufMap[key])
	}

	return bufArr, nil
}

// //

func fileName(pathToFile string) string {
	fullName := filepath.Base(pathToFile)
	ext := filepath.Ext(fullName)
	return fullName[:len(fullName)-len(ext)]
}

//

func parseColumType(name string) microlog.ColumType {
	for t, n := range microlog.ColumMap {
		if n == name {
			return t
		}
	}
	return 0
}

func parseKeyType(name string) microlog.KeyType {
	for t, n := range microlog.KeyMap {
		if n == name {
			return t
		}
	}
	return 0
}
