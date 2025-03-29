package tables

import (
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

	t.Log(tables)
}
