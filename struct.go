package microlog

// // // // // // // // // //

type ColumType byte

type KeyType byte

// //

type InfoColumObj struct {
	Name     string
	Length   uint32
	Type     ColumType
	Key      KeyType
	Children *InfoColumChildrenObj
}

type InfoColumChildrenObj struct {
	Table  *InfoTableObj
	Column *InfoColumObj
}

type InfoTableObj struct {
	Name    string
	Columns []InfoColumObj
}

// // // //

type EmbedTemplateObj struct {
	Path string
	Type string
	Name string
	Data string
}

// // //

type GlobalDocInfoObj struct {
	TemplatePath   string
	GenerationTime string
	Params         map[string]string

	embeddedTemplate *EmbedTemplateObj
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
