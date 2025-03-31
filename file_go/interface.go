package file_go

import (
	"microlog"
)

// // // // // // // // // //

type GeneratorInterface interface {
	Generator(dirPath string, table *microlog.InfoTableObj) error
}
