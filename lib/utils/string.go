package utils

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"math/rand"
	"time"
)

// RandomRangeString 获取指定范围内的随机数
func RandomRangeString(min, max int) string {
	if max == 0 {
		max = 20
	}
	rand.Seed(time.Now().UnixNano())
	size := min + rand.Intn(max)
	alpha := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var buffer bytes.Buffer
	for i := 0; i < size; i++ {
		buffer.WriteByte(alpha[rand.Intn(len(alpha))])
	}
	return buffer.String()
}

// SecretKey 将字符串转换为符合冰蝎、哥斯拉加密要求的 md5[0:16] 后的结果
func SecretKey(pwd string) []byte {
	return []byte(pass2MD5(pwd))
}

// 获取前十六位 md5 值
func pass2MD5(input string) string {
	md5hash := md5.New()
	md5hash.Write([]byte(input))
	return hex.EncodeToString(md5hash.Sum(nil))[0:16]
}

func MD5(input string) string {
	md5hash := md5.New()
	md5hash.Write([]byte(input))
	return hex.EncodeToString(md5hash.Sum(nil))
}

func JsonStrToMap(jsonStr string) (map[string]string, error) {
	m := make(map[string]string)
	err := json.Unmarshal([]byte(jsonStr), &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func MapToJsonStr(m map[string]string) (string, error) {
	jsonByte, err := json.Marshal(m)
	if err != nil {
		return "", err
	}
	return string(jsonByte), nil
}
