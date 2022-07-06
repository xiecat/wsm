package dynamic

import (
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
)

// ReplaceClassStrVar 替换参数
func ReplaceClassStrVar(bs64class []byte, oldVar, newVar string) ([]byte, error) {
	baseHexCode := hex.EncodeToString(bs64class)
	oldLength := fmt.Sprintf("%04x", len(oldVar))
	hexOldVar := oldLength + hex.EncodeToString([]byte(oldVar))
	oldPos := strings.LastIndex(baseHexCode, hexOldVar)
	if oldPos > -1 {
		newLength := fmt.Sprintf("%04x", len(newVar))
		hexNewVar := newLength + hex.EncodeToString([]byte(newVar))
		// 只替换第一次出现的变量，防止类似 user = "sql" 也被替换掉
		retCode := baseHexCode[:oldPos] + strings.Replace(baseHexCode[oldPos:], hexOldVar, hexNewVar, 1)
		retByte, err := hex.DecodeString(retCode)
		if err != nil {
			return nil, err
		}
		return retByte, nil
	}
	return nil, errors.New("class 字节码变量替换失败")
}

// ReplaceClassName 动态替换类名
func ReplaceClassName(classContent []byte, old, new string) []byte {
	classContent = bytes.ReplaceAll(classContent,
		MergeBytes([]byte{(byte)(len(old) + 2), 76},
			[]byte(old)),
		MergeBytes([]byte{(byte)(len(new) + 2), 76},
			[]byte(new)))
	classContent = bytes.ReplaceAll(classContent,
		MergeBytes([]byte{(byte)(len(old))}, []byte(old)),
		MergeBytes([]byte{(byte)(len(new))}, []byte(new)),
	)
	return classContent
}

// ReplaceSourceFile 尝试替换一下 SourceFile 为随机
func ReplaceSourceFile(classContent []byte, old, new string) []byte {
	if !strings.HasSuffix(old, ".java") {
		old = old + ".java"
	}
	if !strings.HasSuffix(new, ".java") {
		old = old + ".java"
	}
	classContent = replaceClassString(classContent, old, new)
	return classContent
}

// ReplaceFuncName 尝试替换一下函数名字为随机
func ReplaceFuncName(classContent []byte, old, new string) []byte {
	classContent = replaceClassString(classContent, old, new)
	return classContent
}

// 替换 class 字节码中的 string
func replaceClassString(classContent []byte, old, new string) []byte {
	classContent = bytes.Replace(classContent,
		MergeBytes([]byte{00, byte(len(old))}, []byte(old)),
		MergeBytes([]byte{00, byte(len(new))}, []byte(new)), 1)
	return classContent
}
