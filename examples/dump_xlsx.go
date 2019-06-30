package examples

import (
	"fmt"

	"github.com/kushao1267/go-tools/excel"
)

func DumpXlsxExample() {
	// data construct method: 1
	exc := excel.Excel{Name: "测试xlsx导出"}
	sheet1 := excel.Sheet{
		Name:   "sheet1",
		Titles: []string{"title1", "title2", "title3"},
		Values: []excel.Value{
			[]interface{}{1, 2, 3},
			[]interface{}{"你", "好", "吗"},
			[]interface{}{"A", "B", "C"},
		},
	}
	sheet2 := excel.Sheet{
		Name:   "sheet2",
		Titles: []string{"title3", "title4", "title5"},
		Values: []excel.Value{
			[]interface{}{4,5,6},
			[]interface{}{"我", "很", "好"},
			[]interface{}{"D", "E", "F"},
		},
	}
	exc.Sheets = []excel.Sheet{sheet1, sheet2}

	// // data construct method: 2
	// excelStr := `
	// {
	// 	"name":"测试xlsx导出",
	// 	"sheets":
	// 		[{
	// 			"name":"sheet1",
	// 			"titles":["title1","title2","title3"],
	// 			"values":[[1,2,3],
	// 					["你","好","吗"],
	// 					["A","B","C"]]
	// 		},{
	// 			"name":"sheet2",
	// 			"titles":["title3","title4","title5"],
	// 			"values":[[4,5,6],
	// 					["我","很","好"],
	// 					["D","E","F"]]
	// 		}]
	// }`
	// exc := excel.Excel{}
	// _ = exc.LoadJson([]byte(excelStr))

	// export to xlsx file
	if e1 := exc.Dump(); e1 != nil {
		fmt.Printf("Dump error: %s", e1.Error())
	}

	// export to json
	// _, j := exc.DumpJson()

	// export to string
	_, s := exc.DumpString()
	fmt.Println(s)
}
