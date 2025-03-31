package microlog

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
