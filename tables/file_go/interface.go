package file_go

import (
	"microlog/tables"
)

// // // // // // // // // //

type GeneratorInterface interface {
	Generator(dirPath string, table *tables.InfoTableObj) error
}
