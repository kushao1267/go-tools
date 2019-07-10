package examples

import (
	"fmt"

	"github.com/kushao1267/go-tools/magic"
)

func GetExtFromFile(filename string) {
	magic.Init()
	fmt.Println("GetExtFromFile: ", magic.FromFile(filename, false))
}

func GetExtListFromFile(filename string) {
	magic.Init()
	var err error
	var list []magic.PureMagicWithConfidence
	if err, list = magic.MagicFile(filename); err !=nil {
		fmt.Println("MagicFile error: ", err.Error())
	}
	fmt.Printf("GetExtListFromFile result:\n %v", list)
}