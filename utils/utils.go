package utils

import (
	"fmt"
	"math/rand"
)

func TrafficConvert(byteCount int64) string {
	kb := int64(1024)
	mb := int64(1048576)
	gb := int64(1073741824)

	if byteCount > gb {
		return fmt.Sprintf("%.2f GB", float64(byteCount)/float64(gb))
	} else if byteCount > mb {
		return fmt.Sprintf("%.2f MB", float64(byteCount)/float64(mb))
	} else if byteCount > kb {
		return fmt.Sprintf("%.2f KB", float64(byteCount)/float64(kb))
	} else if byteCount < 0 {
		return "0"
	} else {
		return fmt.Sprintf("%.2f B", float64(byteCount))
	}
}

// 生成指定区间范围内的随机数
func GenerateRandomNumber(min, max int64) int64 {
	return rand.Int63n(max-min+1) + min
}
