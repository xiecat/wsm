package charset

import (
	"fmt"
	"github.com/gogs/chardet"
	"github.com/yuin/charsetutil"
	"testing"
)

func Test_charsetEncode(t *testing.T) {
	data := "你好"
	e := charsetutil.MustEncodeString(data, "gbk")
	det := chardet.NewTextDetector()
	cs, err := det.DetectAll(e)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v\n", cs)
	d := charsetutil.MustDecodeBytes(e, "gbk")
	fmt.Println(d)
}
