package wsm

import (
	"encoding/base64"
	"encoding/json"
)

type Result struct {
	Raw []byte

	Status int
}

func NewResult(data []byte) *Result {
	return &Result{Raw: data}
}

func (r *Result) Parser() {
	result := make(map[string]string, 2)
	if err := json.Unmarshal(r.Raw, &result); err == nil {
		for k, v := range result {
			value, err := base64.StdEncoding.DecodeString(v)
			if err == nil {
				result[k] = string(value)
			}
		}
	}
}

func (r *Result) GetRaw() string {
	return string(r.Raw)
}
