package parser

import (
	"errors"
	"fmt"
	"microlog/tables/generator"
	"path/filepath"
	"strings"
)

// // // // // // // // // //

func Dir(pathToDir string) ([]generator.InfoTableObj, error) {
	filesArr := scanDir(pathToDir)
	if len(filesArr) == 0 {
		return nil, errors.New("no files found in dir " + pathToDir)
	}

	tablesMap := make(map[string][]generator.InfoColumObj)
	childMap := make(map[string]string)
	retTables := make([]generator.InfoTableObj, 0)

	//

	for _, file := range filesArr {
		path := filepath.Join(pathToDir, file)
		bufMap, err := readFile(path)
		if err != nil {
			return nil, err
		}

		tableName := fileName(file)
		_, ok := tablesMap[tableName]
		if ok {
			return nil, errors.New("duplicate table " + tableName)
		}

		columArr := make([]generator.InfoColumObj, 0)
		for name, obj := range bufMap {

			colum := generator.InfoColumObj{
				Name:   name,
				Length: obj.Len,
				Type:   parseColumType(obj.Type),
				Key:    parseKeyType(obj.Key),
			}
			if obj.Children != "" {
				colum.Children = new(generator.InfoColumChildrenObj)
				childMap[tableName+"."+name] = obj.Children
			}

			columArr = append(columArr, colum)
		}

		tablesMap[tableName] = columArr
		retTables = append(retTables, generator.InfoTableObj{Name: tableName, Columns: columArr})
	}

	//

	for tableName, columArr := range tablesMap {
		for _, column := range columArr {
			if column.Children != nil {
				buf := strings.Split(childMap[tableName+"."+column.Name], ".")
				if len(buf) != 2 {
					return nil, fmt.Errorf(`column "%s.%s" has invalid children`, tableName, column.Name)
				}

				for _, tableBuf := range retTables {
					if tableBuf.Name == buf[0] {
						for _, bufColumn := range tableBuf.Columns {
							if bufColumn.Name == buf[1] {
								column.Children.Column = &bufColumn
								column.Children.Table = &tableBuf
								break
							}
						}
					}
					if column.Children.Column != nil {
						break
					}
				}
			}
		}
	}

	return retTables, nil
}
