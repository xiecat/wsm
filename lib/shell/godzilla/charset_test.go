package godzilla

import (
	"fmt"
	"github.com/yuin/charsetutil"
	"testing"
)

func Test_charsetEncode(t *testing.T) {
	data := "你好"
	e := charsetutil.MustEncodeString(data, "gbk")
	cs, err := charsetutil.GuessBytes(e)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(
		"Detectected character set : %v\n"+
			"language : %v\n"+
			"with confidence %v\n",
		cs.Charset(), cs.Language(), cs.Confidence(),
	)
	d := charsetutil.MustDecodeBytes(e, "gbk")
	fmt.Println(d)
}
