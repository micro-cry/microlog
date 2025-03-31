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

//

type MapLineObj struct {
	Key     string
	Value   string
	Comment string
}
type TemplateMapObj struct {
	NameMap    string
	CommentMap string
	TypeKey    string
	TypeValue  string
	ValuesArr  []*MapLineObj
}
