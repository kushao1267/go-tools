package funcs

import "unicode/utf8"

// StringInSlice 判断字段串是否在slice中
func StringInSlice(target string, list []string) bool {
	for _, elem := range list {
		if elem == target {
			return true
		}
	}
	return false
}

// RemoveDuplicateString 字符串列表去重
func RemoveDuplicateString(intSlice []string) []string {
	keys := make(map[string]bool)
	var list []string
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

// RemoveEmptyStringInSlice 从slice中滤掉指定字符串
func RemoveEmptyStringInSlice(list []string, target string) []string {
	var resultList []string
	for _, elem := range list {
		if elem != target {
			resultList = append(resultList, elem)
		}
	}
	return resultList
}


// ValidUTF8 滤除非UTF8的字符
func ValidUTF8(s string) string {
	if !utf8.ValidString(s) {
		v := make([]rune, 0, len(s))
		for i, r := range s {
			if r == utf8.RuneError {
				_, size := utf8.DecodeRuneInString(s[i:])
				if size == 1 {
					continue
				}
			}
			v = append(v, r)
		}
		return string(v)
	}
	return s
}