package utils

import (
	"strconv"
	"strings"
)

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
