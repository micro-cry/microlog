package microlog

import "fmt"

// // // // // // // // // //

func (b ColumType) Byte() byte {
	return byte(b)
}

func (b ColumType) String() string {
	return ColumMap[b]
}

// //

func (b KeyType) Byte() byte {
	return byte(b)
}

func (b KeyType) String() string {
	return KeyMap[b]
}

// //

func (column *InfoColumObj) TypeString() string {
	switch column.Type {

	case ColumBool, ColumByte, ColumString:
		return column.Type.String()

	case ColumBytes:
		if column.Length == 0 {
			return "[]byte"
		} else {
			return fmt.Sprintf("[%d]byte", column.Length)
		}

	case ColumDateTime:
		return "time.Time"
	}

	return "any"
}
