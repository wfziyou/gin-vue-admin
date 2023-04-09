package utils

import (
	"strconv"
	"strings"
)

func SplitToInt32List(str string, sep string) (i32List []int32) {
	if str == "" {
		return
	}
	strList := strings.Split(str, sep)
	if len(strList) == 0 {
		return
	}
	for _, item := range strList {
		if item == "" {
			continue
		}
		val, err := strconv.ParseInt(item, 10, 32)
		if err != nil {
			// logs.CtxError(ctx, "ParseInt fail, err=%v, str=%v, sep=%v", err, str, sep) // 此处打印出log报错信息
			continue
		}
		i32List = append(i32List, int32(val))
	}
	return
}
func SplitToUint64List(str string, sep string) (i32List []uint64) {
	if str == "" {
		return
	}
	strList := strings.Split(str, sep)
	if len(strList) == 0 {
		return
	}
	for _, item := range strList {
		if item == "" {
			continue
		}
		val, err := strconv.ParseUint(item, 10, 32)
		if err != nil {
			// logs.CtxError(ctx, "ParseInt fail, err=%v, str=%v, sep=%v", err, str, sep) // 此处打印出log报错信息
			continue
		}
		i32List = append(i32List, val)
	}
	return
}

// 类型转换
func InterfaceToInt(v interface{}) (i int) {
	switch v := v.(type) {
	case int:
		i = v
	default:
		i = 0
	}
	return
}

func IdsToArr(data string) (_ids []uint64) {
	var slice = strings.Split(data, ",")
	length := len(slice)
	ids := make([]uint64, 0, length)
	for i := 0; i < length; i++ {
		id, _ := strconv.ParseUint(slice[i], 10, 64)
		ids = append(ids, id)
	}
	return ids
}
