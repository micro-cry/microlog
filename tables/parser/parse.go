package parser

import (
	"errors"
	"fmt"
	"microlog/tables"
	"path/filepath"
	"sort"
	"strings"
)

// // // // // // // // // //

func Dir(pathToDir string) ([]tables.InfoTableObj, error) {
	filesArr := scanDir(pathToDir)
	if len(filesArr) == 0 {
		return nil, errors.New("no files found in dir " + pathToDir)
	}

	tablesMap := make(map[string][]tables.InfoColumObj)
	childMap := make(map[string]string)
	retTables := make([]tables.InfoTableObj, 0)

	//

	sort.Strings(filesArr)
	for _, file := range filesArr {
		path := filepath.Join(pathToDir, file)
		bufArr, err := readFile(path)
		if err != nil {
			return nil, err
		}

		tableName := fileName(file)
		_, ok := tablesMap[tableName]
		if ok {
			return nil, errors.New("duplicate table " + tableName)
		}

		columArr := make([]tables.InfoColumObj, 0)
		for _, obj := range bufArr {

			colum := tables.InfoColumObj{
				Name:   obj.Name,
				Length: obj.Len,
				Type:   parseColumType(obj.Type),
				Key:    parseKeyType(obj.Key),
			}
			if obj.Children != "" {
				colum.Children = new(tables.InfoColumChildrenObj)
				childMap[tableName+"."+obj.Name] = obj.Children
			}

			columArr = append(columArr, colum)
		}

		tablesMap[tableName] = columArr
		retTables = append(retTables, tables.InfoTableObj{Name: tableName, Columns: columArr})
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

				if column.Children.Column == nil {
					return nil, fmt.Errorf(`column "%s.%s" has invalid children`, tableName, column.Name)
				}
				if column.Children.Table == nil {
					return nil, fmt.Errorf(`table "%s.%s" has invalid children`, tableName, column.Name)
				}
			}
		}
	}

	return retTables, nil
}
