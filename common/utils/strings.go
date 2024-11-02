package utils

import (
	"crypto/rand"
	"strconv"
	"time"
)

func CreateCn() string {

	b := make([]byte, 10)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}

	// 将字节转换为整数
	intFromBytes := int64(0)
	for _, v := range b {
		intFromBytes = (intFromBytes << 8) | int64(v)
	}

	// 格式化时间
	formatted := time.Now().Format("060102150405")

	str := strconv.FormatInt(intFromBytes, 10)
	return formatted + str[1:5]

}
