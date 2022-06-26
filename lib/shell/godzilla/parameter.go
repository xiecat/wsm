package godzilla

import (
	"bytes"
	"github.com/Go0p/wsm/lib/dynamic"
)

type Parameter struct {
	HashMap map[string]interface{}
	Size    int
}

func (p *Parameter) AddString(key, value string) {
	p.addParameterString(key, value)
}

func (p *Parameter) AddBytes(key string, value []byte) {
	p.addParameterByteArray(key, value)
}

func (p *Parameter) addParameterString(key, value string) {
	p.addParameterByteArray(key, []byte(value))
}

func (p *Parameter) addParameterByteArray(key string, value []byte) {
	p.HashMap[key] = value
	p.Size += len(value)
}

func (p *Parameter) Serialize() []byte {
	var outputStream bytes.Buffer
	for key, value := range p.HashMap {
		outputStream.Write([]byte(key))
		outputStream.WriteByte(2)
		// 根据这个判断 value 的长度
		outputStream.Write(dynamic.IntToBytes(len(value.([]byte))))
		outputStream.Write(value.([]byte))
	}
	return outputStream.Bytes()
}
