## Quick Start

### 1.通过文件识别其扩展名
```go
import "github.com/kushao1267/go-tools/magic"

magic.Init()
fmt.Println("GetExtFromFile: ", magic.FromFile(filename, false))
```

输出结果:
```
> GetExtFromFile:  .png
```

### 1.通过字符串识别其扩展名
```go
import "github.com/kushao1267/go-tools/magic"

magic.Init()
fmt.Println("GetExtFromString: ", magic.FromString(s, "", false))
```

输出结果:
```
> GetExtFromString:  .png
```
