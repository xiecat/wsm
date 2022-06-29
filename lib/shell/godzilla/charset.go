package godzilla

import (
	"github.com/yuin/charsetutil"
	"log"
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

// 猜的太不准了，中文加英文必猜错
func (e *EncodingCharset) chardet(data []byte) error {
	if e.charset == Chardet {
		guess, err := charsetutil.GuessBytes(data)
		if err != nil {
			return err
		}
		e.charset = guess.Charset()
	}
	return nil
}

func (e *EncodingCharset) CharsetEncode(input string) ([]byte, error) {
	b, err := charsetutil.EncodeString(input, "gbk")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return b, nil
}

func (e *EncodingCharset) CharsetDecode(input []byte) (string, error) {
	e.chardet(input)
	b, err := charsetutil.DecodeBytes(input, "gbk")
	if err != nil {
		log.Println(err)
		return "", err
	}
	return b, nil
}
