package gzip

import (
	"bytes"
	"compress/gzip"
	"io"
)

func Compress(src []byte) ([]byte, error) {
	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)
	if _, err := gz.Write(src); err != nil {
		return nil, err
	}
	if err := gz.Flush(); err != nil {
		return nil, err
	}
	if err := gz.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func DeCompress(ret []byte) ([]byte, error) {
	var reader *gzip.Reader
	var err error
	reader, err = gzip.NewReader(bytes.NewBuffer(ret))
	if err != nil {
		return nil, err
	}
	var bufBytes bytes.Buffer
	_, err = io.Copy(&bufBytes, reader)
	reader.Close()
	if err != nil {
		return bufBytes.Bytes(), err
	}
	return bufBytes.Bytes(), nil
}
