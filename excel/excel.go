package excel

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
)

const FileSuffix = ".xlsx"

var ALPHABET = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

// getFileNameFromPath 从路径中获取文件名
func getFileNameFromPath(path string) string {
	splits := strings.Split(path, "/")
	return splits[len(splits)-1]
}

// convertToInterfaceSlice 将String类型的slice转为interface的slice
func convertStr2InterSlice(s []string) []interface{} {
	interSlice := make([]interface{}, len(s))
	for i, v := range s {
		interSlice[i] = v
	}
	return interSlice
}

// getColumnsID获取所有列标识
func getColumnsID() (allColumnsID []string) {
	nums := []string{""}
	nums = append(nums, ALPHABET...) // 导出的Excel最多支持26 * 27列
	for _, n := range nums {
		for _, ID := range ALPHABET {
			allColumnsID = append(allColumnsID, n+ID)
		}
	}
	return allColumnsID
}

// Value 每行数据结构, 元素数必须与title数保持一致，超过则以title为准
type Value []interface{}

// Sheet 页数据对象
type Sheet struct {
	Name   string   `json:"name"`   // sheet name
	Titles []string `json:"titles"` // 该sheet的title列表，从左至右
	Values []Value  `json:"values"` // 该sheet的数据行
}

// Excel 文件数据对象
type Excel struct {
	// Excel的文件名
	Name string `json:"name"`
	// 页
	Sheets []Sheet `json:"sheets"`
}

// LoadMap 将map结构转为Excel便于插入
func (excel *Excel) LoadMap(m map[string]interface{}) error {
	jsonString, e := json.Marshal(m)
	if e != nil {
		return e
	}
	e1 := excel.LoadJson(jsonString)
	return e1
}

// LoadJson 将Json转为Excel便于插入
func (excel *Excel) LoadJson(j []byte) error {
	e := json.Unmarshal(j, excel)
	return e
}

// Load 导入xlsx文件
func (excel *Excel) Load(filename string) error {
	f, err := excelize.OpenFile(filename)
	if err != nil {
		return err
	}
	excel.Name = getFileNameFromPath(filename)
	// Get all the rows by sheet.
	sheetMap := f.GetSheetMap()
	for _, sheetName := range sheetMap {
		sheet := Sheet{}
		sheet.Name = sheetName
		rows, _ := f.GetRows(sheetName)
		for index, row := range rows {
			if index == 0 {
				// 写入titles
				sheet.Titles = row
			} else {
				sheet.Values = append(sheet.Values, convertStr2InterSlice(row))
			}
		}
		excel.Sheets = append(excel.Sheets, sheet)
	}
	return nil
}

// DumpJson 将Excel类型转为Json
func (excel Excel) DumpJson() (error, []byte) {
	excelJson, e := json.Marshal(excel)
	return e, excelJson
}

// DumpString 将Excel类型转为String
func (excel Excel) DumpString() (error, string) {
	e, j := excel.DumpJson()
	return e, string(j[:])
}

// Dump 导出为xlsx文件
func (excel Excel) Dump() (e error) {
	// 名字纠正
	if !strings.HasSuffix(excel.Name, FileSuffix) {
		excel.Name += FileSuffix
	}

	f := excelize.NewFile()
	allColumnsID := getColumnsID()
	// set sheets
	for index, sheet := range excel.Sheets {
		if index == 0 {
			f.SetSheetName("Sheet1", sheet.Name)
		} else {
			f.NewSheet(sheet.Name)
		}
		// set titles
		for dIndex, dValue := range sheet.Titles {
			_ = f.SetCellValue(sheet.Name, allColumnsID[dIndex]+"1", dValue)
		}
		// set values
		for dIndex, dValue := range sheet.Values {
			number := strconv.Itoa(dIndex + 2)
			for dIndex1, dValue1 := range dValue {
				_ = f.SetCellValue(sheet.Name, allColumnsID[dIndex1]+number, dValue1)
			}
		}
	}
	// 保存到当前目录
	e = f.SaveAs(excel.Name)
	return e
}
