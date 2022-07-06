package wsm

import (
	"encoding/base64"
	"encoding/json"
	"github.com/Go0p/wsm/lib/utils"
)

type bResult struct {
	Raw    []byte
	Body   map[string]string
	Status bool
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

func newBResult(raw []byte) *bResult {
	return &bResult{Raw: raw}
}

func (b *bResult) Parser() error {
	result := make(map[string]string, 2)
	err := json.Unmarshal(b.Raw, &result)
	if err != nil {
		return err
	}
	for k, v := range result {
		value, err := base64.StdEncoding.DecodeString(v)
		if err != nil {
			return err
		}
		result[k] = string(value)
	}
	b.Body = result
	return nil
}

type gResult struct {
	Raw    []byte
	Body   string
	Status bool
}

func newGResult(raw []byte) *gResult {
	return &gResult{Raw: raw}
}

func (g gResult) Parser() error {
	//TODO implement me
	panic("implement me")
}

func (g gResult) ToMap() map[string]string {
	//TODO implement me
	panic("implement me")
}

func (g gResult) ToString() string {
	return string(g.Raw)
}
