package tables

import "database/sql"

// // // // // // // // // //

type ColumnNameInterface interface {
	String() string
	StringSQL() string

	Type() string
	TableName() string
}

type DataInterface interface {
	JSON() ([]byte, error)
	Children() DataTableInterface
}

type DataTableInterface interface {
	JSON() ([]byte, error)
	Parent() DataInterface
	TableName() string
	TableColumns() map[ColumnNameInterface]string
}

// //

type GoFuncGeneratorInterface interface {
	TableCheck(*sql.DB) error
	TableCreate(*sql.DB) error
	TableClear(*sql.DB) error
	TableOptimize(*sql.DB) error

	TableGetSize(*sql.DB) (uint64, error)
}
