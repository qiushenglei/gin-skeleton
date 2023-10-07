package utils

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

// Str2MD5 returns a MD5 hash in string form of the passed-in `s`
func Str2MD5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

// StringConcat returns the concatenation of `base` and `strs` strings seperated by `sep`
func StringConcat(sep string, strs ...string) string {
	return strings.Join(strs, sep)
}

func GenerateRandomString(l int, prefix string) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	temp, _ := uuid.NewUUID()
	uid := temp.String()

	//return prefix + string(result) + fmt.Sprintf("%x", md5.Sum([]byte(uid)))
	return prefix + fmt.Sprintf("%x", md5.Sum([]byte(uid)))
}

// StringJoin 生成sql的in条件
func StringJoin(glue string, pieces []string) (ret string) {
	for index, word := range pieces {
		ret += "'" + word + "'"
		if index < len(pieces)-1 {
			ret += glue
		}
	}
	return
}

// IntJoin  int join
func IntJoin(glue string, pieces []int) (ret string) {
	for index, word := range pieces {
		ret += strconv.Itoa(int(word))
		if index < len(pieces)-1 {
			ret += glue
		}
	}
	return
}
func Int64Join(glue string, pieces []int64) (ret string) {
	for index, word := range pieces {
		ret += strconv.Itoa(int(word))
		if index < len(pieces)-1 {
			ret += glue
		}
	}
	return
}

// AddQuotes 单引号包裹字符串
func AddQuotes(str string) string {
	return "'" + str + "'"
}

// 字符串拼接（非sql）
func StringJoinNoSql(glue string, pieces []string) (ret string) {
	for index, word := range pieces {
		ret += word
		if index < len(pieces)-1 {
			ret += glue
		}
	}
	return
}
