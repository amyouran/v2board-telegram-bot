package stringutil

import (
	"math/rand"
)

func GetRandomString(strArray []string) string {
	// 生成随机索引
	randomIndex := rand.Intn(len(strArray))
	// 返回随机字符串
	return strArray[randomIndex]
}
