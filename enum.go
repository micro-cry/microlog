package microlog

// // // // // // // // // //

const (
	ColumUndefined ColumType = 0
	ColumBool      ColumType = 1
	ColumByte      ColumType = 2
	ColumBytes     ColumType = 3
	ColumString    ColumType = 4
	ColumDateTime  ColumType = 5
)

var ColumMap = map[ColumType]string{
	ColumUndefined: "undefined",
	ColumBool:      "bool",
	ColumByte:      "byte",
	ColumBytes:     "bytes",
	ColumString:    "string",
	ColumDateTime:  "datetime",
}

// //

const (
	KeyNone    KeyType = 0
	KeyPrimary KeyType = 1
	KeyIndex   KeyType = 2
)

var KeyMap = map[KeyType]string{
	KeyNone:    "none",
	KeyPrimary: "primary",
	KeyIndex:   "index",
}
