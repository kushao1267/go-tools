package magic

/* puremagic的简化版，无confidence，仅返回可能性最大的file extension*/

import (
	"encoding/hex"
	"encoding/json"
	"log"
	"reflect"
	"strconv"
)

type PureMagic struct {
	byteMatch []byte
	offset    int
	extension string
	mimeType  string
	name      string
}

var HeaderArray []PureMagic
var FooterArray []PureMagic

func Init() {
	_, HeaderArray, FooterArray = loadMagicData()
}

// fileDetails Grab the start and end of the file
func fileDetails(filename) {
}

func InterfaceSlice(slice interface{}) []interface{} {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		panic("InterfaceSlice() given a non-slice type")
	}

	ret := make([]interface{}, s.Len())

	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}

	return ret
}

// loadMagicData 内部加载magic数据的方法
func loadMagicData() (e error, headers, footers []PureMagic) {
	m := make(map[string]interface{})
	if e := json.Unmarshal([]byte(magicDataJson), m); e != nil {
		log.Println("loadMagicData error:", e.Error())
		return e, headers, footers
	}

	for k, t := range m {
		var l []PureMagic
		slicedT := InterfaceSlice(t)
		for _, pure := range slicedT {
			typedPure := InterfaceSlice(pure)
			dst := strconv.QuoteToASCII(typedPure[0].(string)) // 转ASCII,过滤非ASCII
			byteM, _ := hex.DecodeString(dst)                  // hex -> byte
			l = append(l, PureMagic{
				byteM,
				typedPure[1].(int),
				typedPure[2].(string),
				typedPure[3].(string),
				typedPure[4].(string),
			})
		}
		if k == "headers" {
			headers = l
		} else if k == "footers" {
			footers = l
		}
	}
	return nil, headers, footers
}

// maxLengths 获取MagicHeaderArray和MagicFooterArray的最长bytes长度
func maxLengths() (maxHeaderLength, maxFooterLength int) {
	for _, header := range HeaderArray {
		length := len(header.byteMatch) + header.offset
		if length > maxHeaderLength {
			maxHeaderLength = length
		}
	}
	for _, footer := range FooterArray {
		length := len(footer.byteMatch) + footer.offset
		if length > maxFooterLength {
			maxFooterLength = length
		}
	}
	return maxHeaderLength, maxFooterLength
}

