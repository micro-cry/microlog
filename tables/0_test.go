package tables

import (
	"microlog/tables/generator/file_md"
	"microlog/tables/generator/parser"
	"testing"
)

// // // // // // // // // //

func TestGenerate(t *testing.T) {
	tables, err := parser.Dir("./")

	if err != nil {
		t.Error(err)
		return
	}

	t.Log(file_md.Generate(tables, "tables structure.md"))
}
