package file_go

import "microlog/tables/generator"

// // // // // // // // // //

type GeneratorInterface interface {
	Generator(dirPath string, table *generator.InfoTableObj) error
}
