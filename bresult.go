package wsm

import (
	"encoding/base64"
	"github.com/xiecat/wsm/lib/utils"
)

type bResult struct {
	Raw    []byte
	Body   map[string]string
	Status bool
}

func newBResult(raw []byte) *bResult {
	return &bResult{Raw: raw}
}

func (b *bResult) Parser() error {
	result, err := utils.JsonStrToMap(string(b.Raw))
	if err != nil {
		return err
	}
	for k, v := range result {
		value, _ := base64.StdEncoding.DecodeString(v)
		//if err != nil {
		//	result[k] = string(value)
		//}
		result[k] = string(value)
	}
	b.Body = result
	return nil
}

func (b *bResult) ToMap() map[string]string {
	return b.Body
}

func (b *bResult) ToString() string {
	str, err := utils.MapToJsonStr(b.Body)
	if err != nil {
		return ""
	}
	return str
}
