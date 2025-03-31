package parser

import (
	"microlog/tables"
	"testing"
)

// // // // // // // // // //

func TestFileName(t *testing.T) {
	input := "/path/to/testfile.txt"
	expected := "testfile"
	result := fileName(input)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestParseColumType(t *testing.T) {
	expected := tables.ColumByte
	result := parseColumType(expected.String())
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestParseKeyType(t *testing.T) {
	expected := tables.KeyPrimary
	result := parseKeyType(expected.String())
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
