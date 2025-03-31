package generator_go

import _ "embed"

//go:embed templates/values.tmpl
var ValuesFile string

//go:embed templates/struct.tmpl
var StructFile string

//go:embed templates/sqlite_table.tmpl
var SQLiteTableFile string

//go:embed templates/sqlite_other.tmpl
var SQLiteOtherFile string

//go:embed templates/sqlite_get.tmpl
var SQLiteGetFile string

//go:embed templates/func.tmpl
var FuncFile string

//go:embed templates/functest.tmpl
var FuncTestFile string

//go:embed templates/sqlite.tmpl
var SQLiteFile string
