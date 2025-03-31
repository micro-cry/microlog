package generator_go

import _ "embed"

// // // // // // // // // //

var (
	//go:embed templates/values.tmpl
	ValuesFile string

	//go:embed templates/struct.tmpl
	StructFile string

	//go:embed templates/sqlite_table.tmpl
	SQLiteTableFile string

	//go:embed templates/sqlite_other.tmpl
	SQLiteOtherFile string

	//go:embed templates/sqlite_get.tmpl
	SQLiteGetFile string

	//go:embed templates/func.tmpl
	FuncFile string

	//go:embed templates/functest.tmpl
	FuncTestFile string

	//go:embed templates/sqlite.tmpl
	SQLiteFile string
)
