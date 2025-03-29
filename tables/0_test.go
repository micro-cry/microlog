package tables

import (
	"encoding/json"
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

	data, err := json.MarshalIndent(tables, "", "  ")
	if err != nil {
		t.Error(err)
	}
	t.Log(string(data))
}
