package tables

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
	Ping() error

	TableCheck() error
	TableCreate() error
	TableClear() error
	TableOptimize() error

	GetSize() (uint64, error)
	GetFromUID([]byte) (DataTableInterface, error)
	GetFromUIDs(...[]byte) ([]DataTableInterface, error)
	GetFromIndexes(map[ColumnNameInterface]any) ([]DataTableInterface, error)
	GetFromIndexesLimit(indexes map[ColumnNameInterface]any, indent, size uint32) ([]DataTableInterface, error)

	Add(...DataTableInterface) error
	EditFromUID([]byte, map[ColumnNameInterface]any) error

	DeleteFromUID([]byte) error
	DeleteFromUIDs(...[]byte) error
}
