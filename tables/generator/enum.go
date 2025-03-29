package generator

// // // // // // // // // //

const (
	ColumBool     ColumType = 1
	ColumByte     ColumType = 2
	ColumBytes    ColumType = 3
	ColumString   ColumType = 4
	ColumDateTime ColumType = 5
)

var ColumMap = map[ColumType]string{
	ColumBool:     "bool",
	ColumByte:     "byte",
	ColumBytes:    "bytes",
	ColumString:   "string",
	ColumDateTime: "datetime",
}

// //

const (
	KeyNone    KeyType = 0
	KeyPrimary KeyType = 1
)

var KeyMap = map[KeyType]string{
	KeyNone:    "none",
	KeyPrimary: "primary",
}
