package godzilla

import (
	"fmt"
	"github.com/olaure/chardet"
	"testing"
)

func Test_charsetEncode(t *testing.T) {
	data := []byte("你好哈哈哈")
	detected := chardet.Detect(data)
	fmt.Printf(
		"Detectected character set : %v language : %v with confidence %v\n",
		detected.Encoding, detected.Language, detected.Confidence,
	)
}
