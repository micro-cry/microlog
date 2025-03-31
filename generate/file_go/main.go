package main

import (
	"fmt"
	"microlog/file_go"
	"microlog/parser"
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

	filesPath := filepath.Join(rootDir, "yml")

	// //

	tables, err := parser.Dir(filesPath)
	if err != nil {
		panic(err)
	}

	if len(tables) == 0 {
		panic("no tables: " + filesPath)
	}

	err = file_go.Generate(tables, "microlog", filepath.Join(rootDir, "generate_temp"))
	if err != nil {
		panic(err)
	}
}
