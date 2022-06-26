package dynamic

import (
	"bytes"
	"encoding/binary"
	"math/rand"
	"strings"
	"time"
)

// MatchData 找字节
func MatchData(srcData, dataToFind []byte) int {
	iDataLen := len(srcData)
	iDataToFindLen := len(dataToFind)
	//bGotData := false
	iMatchDataCntr := 0

	for i := 0; i < iDataLen; i++ {
		if srcData[i] == dataToFind[iMatchDataCntr] {
			iMatchDataCntr++
			//bGotData = true
		} else if srcData[i] == dataToFind[0] {
			iMatchDataCntr = 1
		} else {
			iMatchDataCntr = 0
			//bGotData = false
		}

		if iMatchDataCntr == iDataToFindLen {
			return i - len(dataToFind) + 1
		}
	}
	return -1
}

// GetIndexAndLastIndex 返回正常密文起始值、结束值
func GetIndexAndLastIndex(src, substr []byte) (index int, endIndex int) {
	if bytes.Contains(src, substr) {
		index = bytes.Index(src, substr)
		endIndex = index + len(substr)
	} else {
		return -1, -1
	}
	return
}

// MergeBytes 合并 byte 数组
func MergeBytes(a, b []byte) []byte {
	return append(a, b...)
}

// InStrSlice 判断字符串是否在数组中
func InStrSlice(array []string, str string) bool {
	for _, e := range array {
		if e == str {
			return true
		}
	}

	return false
}

// IntToBytes int 转 bytes， 小端模式
func IntToBytes(value int) []byte {
	src := []byte{(byte)(value & 255), (byte)(value >> 8 & 255), (byte)(value >> 16 & 255), (byte)(value >> 24 & 255)}
	return src
}

func intToBytes(n int) []byte {
	x := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.LittleEndian, x)
	return bytesBuffer.Bytes()
}

// BytesToInt 字节转换成整形
func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)

	var x int32
	binary.Read(bytesBuffer, binary.LittleEndian, &x)

	return int(x)
}

// RandomClassName 随机类名
func RandomClassName() string {
	domainAs := []string{"com", "net", "org", "sun"}
	rand.Seed(time.Now().UnixNano())
	domainB := strings.ToLower(randomAlpha(rand.Intn(5) + 3))
	domainC := strings.ToLower(randomAlpha(rand.Intn(5) + 3))
	domainD := strings.ToLower(randomAlpha(rand.Intn(5) + 3))
	className := randomAlpha(rand.Intn(7) + 4)
	className = strings.ToUpper(className[0:1]) + strings.ToLower(className[1:])
	domainAIndex := rand.Intn(4)
	domainA := domainAs[domainAIndex]
	randomSegments := rand.Intn(3) + 3
	var randomName string
	switch randomSegments {
	case 3:
		randomName = domainA + "/" + domainB + "/" + className
		break
	case 4:
		randomName = domainA + "/" + domainB + "/" + domainC + "/" + className
		break
	case 5:
		randomName = domainA + "/" + domainB + "/" + domainC + "/" + domainD + "/" + className
		break
	default:
		randomName = domainA + "/" + domainB + "/" + domainC + "/" + domainD + "/" + className
	}

	return randomName
}

func randomAlpha(size int) string {
	alpha := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var buffer bytes.Buffer

	//rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		buffer.WriteByte(alpha[rand.Intn(len(alpha))])
	}
	return buffer.String()
}
