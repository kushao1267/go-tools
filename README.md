## go-tools

> Golang tools for Human

golang常见的一些小工具，代码片段，建议COPY使用

## 目前支持
- excel导入、导出工具
- 发送邮件

## 安装
`go get -u github.com/kushao1267/go-tools`

## 使用方式

### excel

支持excel的各种导入、导出方式，封装极易使用的API接口。

1.通过内置类型构造excel数据
```go
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
```

2.通过Json字符串构造excel数据
```go
excelStr := `
{
    "name":"测试xlsx导出",
    "sheets":
        [{
            "name":"sheet1",
            "titles":["title1","title2","title3"],
            "values":[[1,2,3],
                    ["你","好","吗"],
                    ["A","B","C"]]
        },{
            "name":"sheet2",
            "titles":["title3","title4","title5"],
            "values":[[4,5,6],
                    ["我","很","好"],
                    ["D","E","F"]]
        }]
}`
_, exc := excel.JsonToExcel([]byte(excelStr))
```

3.导出excel文件

导出后将在当前目录下找到新建的.xlsx文件

```go
if e1 := exc.ExportXlsx(); e1 != nil {
    fmt.Printf("ExportXlsx error: %s", e1.Error())
}
```

4.导入excel文件

导出后将在当前目录下找到新建的.xlsx文件

```go
exc := &excel.Excel{}
if e := exc.Import("./测试xlsx导出.xlsx"); e != nil {
    fmt.Println(e.Error())
}
_, s := exc.ExportString()
fmt.Println(s)
```
output:
```shell
{"name":"测试xlsx导出.xlsx","sheets":[{"name":"sheet1","titles":["title1","title2","title3"],"values":[["1","2","3"],["你","好","吗"],["A","B","C"]]},{"name":"sheet2","titles":["title3","title4","title5"],"values":[["4","5","6"],["我","很","好"],["D","E","F"]]}]}
```

5.导出excel.Excel变量为Json

```go
_, j := exc.ExportJson()
	fmt.Println(j)
```

6.导出excel.Excel变量为String

```go
_, s := exc.ExportString()
	fmt.Println(s)
```

## License

MIT
