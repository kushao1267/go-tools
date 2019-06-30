package examples

import (
	"fmt"

	"github.com/kushao1267/go-tools/excel"
)

func LoadXlsxExample() {
	// 导入Excel文件
	exc := &excel.Excel{}
	if e := exc.Load("./测试xlsx导出.xlsx"); e != nil {
		fmt.Println(e.Error())
	}
	_, s := exc.DumpString()
	fmt.Println(s)
}
