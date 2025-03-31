package main

import (
	"fmt"
	"microlog/generator_md"
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

	fmt.Println("Generate MD\t\t", rootDir)

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

	err = generator_md.Generate(tables, "TEMP/tables structure.md")
	if err != nil {
		panic(err)
	}
}
