package generator_template

// // // // // // // // // //

type StructLineObj struct {
	Name    string
	Type    string
	Reflect string
	Comment string
}
type TemplateStructObj struct {
	NameStruct    string
	CommentStruct string
	LinesArr      []*StructLineObj
}
