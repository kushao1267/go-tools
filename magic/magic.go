package magic

/* Inspired by puremagic*/

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"log"
	"os"
	"reflect"
	"sort"
	"strings"
)

type PureMagic struct {
	byteMatch string
	offset    int
	extension []string
	mimeType  string
	name      string
}

type PureMagicWithConfidence struct {
	pureMagic  PureMagic
	confidence float32
}

var HeaderArray []PureMagic
var FooterArray []PureMagic

// Init 使用magic之前必须调用
func Init() {
	_, HeaderArray, FooterArray = loadMagicData()
}
func check(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}

// fileDetails Grab the start and end of the file
func fileDetails(filename string) (string, string) {
	maxHeaderLength, maxFooterLength := maxLengths()
	var head, foot []byte
	f, err := os.Open(filename)
	check(err)
	head = make([]byte, maxHeaderLength)
	_, err1 := f.Read(head)
	check(err1)
	_, err2 := f.Seek(-int64(maxFooterLength), 2)
	if err2 != nil {
		_, _ = f.Seek(0, 0)
	}
	foot = make([]byte, maxFooterLength)
	_, _ = f.Read(foot)
	return string(head), string(foot)
}

// stringDetails Grab the start and end of the string
func stringDetails(s string) (head, foot string) {
	maxHeaderLength, maxFooterLength := maxLengths()
	return s[:maxHeaderLength], s[-maxFooterLength:]
}

func confidence(matches []PureMagic, ext string) (results []PureMagicWithConfidence) {
	for _, match := range matches {
		var con float32
		if len(match.extension) > 9 {
			con = 0.8
		} else {
			con = 0.1 * float32(len(match.extension))
		}
		if ext == match.extension[0] {
			con = 0.9
		}
		results = append(results, PureMagicWithConfidence{pureMagic: match, confidence: con})
	}
	// sorted by confidence
	sort.Slice(results, func(i, j int) bool {
		return results[i].confidence > results[j].confidence
	})
	return results
}

// magic: Discover what type of file it is based on the incoming string
func magic(header, footer, ext string, mime bool) []string {
	if len(header) == 0 {
		return []string{""}
	}
	infos, _ := identifyAll(header, footer, ext)
	info := infos[0]
	if mime {
		return []string{info.pureMagic.mimeType}
	}
	return info.pureMagic.extension
}

func identifyAll(header, footer, ext string) ([]PureMagicWithConfidence, error) {
	var matches []PureMagic
	for _, magicRow := range HeaderArray {
		start := magicRow.offset
		end := magicRow.offset + len(magicRow.byteMatch)
		if end > len(header) {
			continue
		}
		if header[start:end] == magicRow.byteMatch {
			matches = append(matches, magicRow)
		}
	}
	for _, magicRow := range FooterArray {
		start := magicRow.offset
		// 解决golang无负数切片的问题
		index_start := len(footer) + start
		if index_start < 0 {
			index_start = 0
		} else {
			index_start = start
		}

		if footer[index_start:] == magicRow.byteMatch {
			matches = append(matches, magicRow)
		}
	}
	if len(matches) == 0 {
		return []PureMagicWithConfidence{}, errors.New("Could not identify file")
	}
	return confidence(matches, ext), nil
}

// contains whether slice contains element
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// Scan a filename for it's extension.
func extFromFilename(filename string) string {
	splits := strings.Split(strings.ToLower(filename), ".")
	if len(splits) < 2 {
		return ""
	}
	base := splits[1]
	ext := "." + splits[2]
	var chainArray []PureMagic
	var allExts []string
	chainArray = append(chainArray, HeaderArray...)
	chainArray = append(chainArray, FooterArray...)
	for _, ca := range chainArray {
		allExts = append(allExts, ca.extension...)
	}
	// if long_ext in all_exts:
	// return long_ext
	if strings.HasPrefix(base[len(base)-4:], ".") {
		// For double extensions like like .tar.gz
		longExt := base[len(base)-4:] + ext
		if contains(allExts, longExt) {
			return longExt
		}
	}
	return ext
}

// ConvertInterface2Slice 将interface转为slice类型
func ConvertInterface2Slice(slice interface{}) (error, []interface{}) {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		return errors.New("ConvertInterface2Slice() given a non-slice type"), []interface{}{}
	}

	ret := make([]interface{}, s.Len())

	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}

	return nil, ret
}

// loadMagicData 内部加载magic数据的方法
func loadMagicData() (e error, headers, footers []PureMagic) {
	m := make(map[string]interface{})
	if e := json.Unmarshal([]byte(magicDataJson), &m); e != nil {
		log.Println("loadMagicData error:", e.Error())
		return e, headers, footers
	}

	for k, t := range m {
		var l []PureMagic
		// 库内数据，认为无error
		_, slicedT := ConvertInterface2Slice(t)
		for _, pure := range slicedT {
			_, typedPure := ConvertInterface2Slice(pure)
			byteM, _ := hex.DecodeString(typedPure[0].(string)) // hex -> byte
			var extensions []string
			for _, ext := range typedPure[2].([]interface{}) {
				extensions = append(extensions, ext.(string))
			}

			l = append(l, PureMagic{
				string(byteM),
				int(typedPure[1].(float64)),
				extensions,
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

// MagicString: Reads in file, return list of possible matches(extension), highest confidence first
func MagicString(s, filename string) (error, []PureMagicWithConfidence) {
	if len(s) == 0 {
		return errors.New("Input was empty."), []PureMagicWithConfidence{}
	}
	head, foot := stringDetails(s)
	var ext string
	if len(filename) > 0 {
		ext = extFromFilename(filename)
	} else {
		ext = ""
	}
	infos, _ := identifyAll(head, foot, ext)
	sort.Slice(infos, func(i, j int) bool {
		return infos[i].confidence > infos[j].confidence
	})
	return nil, infos
}

// MagicFile: Reads in file, return list of possible matches(extension), highest confidence first
func MagicFile(filename string) (error, []PureMagicWithConfidence) {
	head, foot := fileDetails(filename)
	if len(head) == 0 {
		return errors.New("Input was empty."), []PureMagicWithConfidence{}
	}
	infos, _ := identifyAll(head, foot, extFromFilename(filename))
	sort.Slice(infos, func(i, j int) bool {
		return infos[i].confidence > infos[j].confidence
	})
	return nil, infos
}

// FromString: Reads in string, return extension
func FromString(s, filename string, mime bool) string {
	head, foot := stringDetails(s)
	var ext string
	if len(filename) > 0 {
		ext = extFromFilename(filename)

	} else {
		ext = ""
	}
	return magic(head, foot, ext, mime)[0]
}

// FromFile: Reads from file, return extension
func FromFile(filename string, mime bool) string {
	head, foot := fileDetails(filename)
	return magic(head, foot, "", mime)[0]
}
