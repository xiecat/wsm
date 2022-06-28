package godzilla

import (
	"bytes"
	"fmt"
	"github.com/qiniu/iconv"
	"github.com/saintfish/chardet"
	"io/ioutil"
	"log"
	"strings"
)

const (
	Chardet         = "auto"
	UTF8CharSet     = "UTF-8"
	GBKCharSet      = "GBK"
	GB2312CharSet   = "GB2312"
	BIG5CharSet     = "BIG5"
	ISO88591CharSet = "ISO-8859-1"
	GB18030CharSet  = "GB18030"
	// Latin1CharSet ISO-8859-1 的别名，有些环境下写作 Latin-1
	Latin1CharSet = "latin1"
	UTF16CharSet  = "UTF16"
	AsciiCharSet  = "ascii"
	Cp850CharSet  = "cp850"
)

type EncodingCharset struct {
	charset string
}

func (e *EncodingCharset) GetCharset() string {
	return e.charset
}
func (e *EncodingCharset) SetCharset(c string) {
	e.charset = c
}

func (e *EncodingCharset) tryChardet(data []byte) (string, error) {
	det := chardet.NewTextDetector()
	charsetMap, err := det.DetectBest(data)
	fmt.Println(charsetMap)
	if err != nil {
		return "", err
	}
	return charsetMap.Charset, nil
}

func (e *EncodingCharset) CharsetEncode(input string) ([]byte, error) {
	// convert gbk to utf8
	cd, err := iconv.Open("gbk", "utf-8")
	if err != nil {
		fmt.Println("iconv.Open failed!")
		return nil, err
	}
	defer cd.Close()
	i := strings.NewReader(input)
	bufSize := 0 // default if zero
	r := iconv.NewReader(cd, i, bufSize)

	buf, err := ioutil.ReadAll(r)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return buf, nil
}

func (e *EncodingCharset) CharsetDecode(input []byte) (string, error) {
	if e.charset == Chardet {
		e.charset, _ = e.tryChardet(input)
		log.Println(e.charset)
	}
	// convert gbk to utf8
	cd, err := iconv.Open("utf-8", e.charset)
	if err != nil {
		fmt.Println("iconv.Open failed!")
		return "", err
	}
	defer cd.Close()
	i := bytes.NewReader(input)
	bufSize := 0 // default if zero
	r := iconv.NewReader(cd, i, bufSize)

	buf, err := ioutil.ReadAll(r)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return string(buf), nil
}
