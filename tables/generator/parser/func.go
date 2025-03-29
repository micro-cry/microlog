package parser

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"microlog/tables/generator"
	"os"
	"path/filepath"
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

func readFile(pathToFile string) (map[string]ColumObj, error) {
	data, err := os.ReadFile(pathToFile)
	if err != nil {
		return nil, err
	}

	var rawData map[string]map[string]interface{}
	err = yaml.Unmarshal(data, &rawData)
	if err != nil {
		return nil, err
	}

	result := make(map[string]ColumObj)

	for key, props := range rawData {
		var col ColumObj

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

		result[key] = col
	}

	return result, nil
}

// //

func fileName(pathToFile string) string {
	fullName := filepath.Base(pathToFile)
	ext := filepath.Ext(fullName)
	return fullName[:len(fullName)-len(ext)]
}

//

func parseColumType(name string) generator.ColumType {
	for t, n := range generator.ColumMap {
		if n == name {
			return t
		}
	}
	return 0
}

func parseKeyType(name string) generator.KeyType {
	for t, n := range generator.KeyMap {
		if n == name {
			return t
		}
	}
	return 0
}
