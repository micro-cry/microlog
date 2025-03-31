package generator_go

// // // // // // // // // //

type EmbedTemplateObj struct {
	Path string
	Data string
}

// // //

type GlobalDocInfoObj struct {
	TemplatePath   string
	GenerationTime string
	Params         map[string]string

	PackageName string
}

//

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
