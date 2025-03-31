package microlog

import _ "embed"

// // // // // // // // // //

var (

	//go:embed templates/md-tables-struct.tmpl
	MdTablesFile string

	//go:embed templates/go-values.tmpl
	ValuesFile string

	//go:embed templates/go-struct.tmpl
	StructFile string

	//go:embed templates/go-sqlite_table.tmpl
	SQLiteTableFile string

	//go:embed templates/go-sqlite_other.tmpl
	SQLiteOtherFile string

	//go:embed templates/go-sqlite_get.tmpl
	SQLiteGetFile string

	//go:embed templates/go-func.tmpl
	FuncFile string

	//go:embed templates/go-functest.tmpl
	FuncTestFile string

	//go:embed templates/go-sqlite.tmpl
	SQLiteFile string
)
