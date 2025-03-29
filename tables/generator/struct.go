package generator

// // // // // // // // // //

type ColumType byte

type KeyType byte

// //

type InfoColumObj struct {
	Name     string
	Length   uint32
	Type     ColumType
	Key      KeyType
	Children *InfoColumObj
}

type InfoTableObj struct {
	Name    string
	Columns []*InfoColumObj
}
