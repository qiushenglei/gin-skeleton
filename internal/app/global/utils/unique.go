package utils

import (
	"github.com/bwmarrin/snowflake"
	"math/rand"
	"time"
)

func GenerateUniqueNumberByRand() int64 {
	// 使用当前时间戳作为基础值
	base := time.Now().Unix()

	// 生成一个随机数作为后缀，确保唯一性
	rand.Seed(time.Now().UnixNano())
	suffix := rand.Int63n(1000) // 这里假设后缀范围在 0 到 999

	// 组合基础值与后缀得到唯一值
	uniqueValue := base*1000 + suffix

	return uniqueValue
}

func GenerateUniqueNumberBySnowFlake() string {
	n, _ := snowflake.NewNode(1)
	id := n.Generate()
	return id.String()
}
