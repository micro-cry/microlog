package main

import (
	"fmt"
	"microlog/tables/generator/file_go"
	"microlog/tables/generator/parser"
	"os"
	"path/filepath"
)

// // // // // // // // // //

func main() {
	rootDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	fmt.Println("Generate Go\t\t", rootDir)

	if filepath.Base(rootDir) != "microlog" {
		panic("start is not in `microlog` directory")
	}

	filesPath := filepath.Join(rootDir, "tables")

	// //

	tables, err := parser.Dir(filesPath)
	if err != nil {
		panic(err)
	}

	if len(tables) == 0 {
		panic("no tables: " + filesPath)
	}

	filesPath = filepath.Join(filesPath, "struct")
	err = file_go.Generate(tables, filesPath)
	if err != nil {
		panic(err)
	}
}
