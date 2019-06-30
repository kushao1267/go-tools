package examples

import (
	"fmt"

	"github.com/kushao1267/go-tools/magic"
)

func GetExtFromFile(filename string) {
	magic.Init()
	fmt.Println("GetExtFromFile: ", magic.FromFile(filename, false))
}
