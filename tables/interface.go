package tables

// // // // // // // // // //

type ColumnNameInterface interface {
	String() string
	StringSQL() string

	Type() string
	TableName() string
}
