package examples

import (
	"fmt"

	"github.com/kushao1267/go-tools/magic"
)

func GetExtFromString(s string) {
	magic.Init()
	fmt.Println("GetExtFromString: ", magic.FromString(s, "", false))
}
