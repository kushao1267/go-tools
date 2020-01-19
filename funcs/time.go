package funcs

import (
	"strconv"
	"time"
)

// GetNowTimeStamp 获取时间戳
func GetNowTimeStamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// GetNowTimeStampStr 获取时间戳，单位ms，类型为string
func GetNowTimeStampStr() string {
	return strconv.FormatInt(GetNowTimeStamp(), 10)
}

// GetNowTimeStr 获取当前时间格式化字符串
func GetNowTimeStr() string{
	return time.Now().Format("2006-01-02 15:04:05")
}