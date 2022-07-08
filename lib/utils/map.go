package utils

import (
	"encoding/json"
	"errors"
	"fmt"
)

func ToMapParams(params interface{}) (map[string]string, error) {
	raw, err := json.Marshal(params)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("marshal params failed: %s", err))
	}
	var p = map[string]string{}
	err = json.Unmarshal(raw, &p)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("unmarshal map params failed: %s", err))
	}

	return p, nil
}

func JsonStrToMap(jsonStr string) (map[string]string, error) {
	m := make(map[string]string)
	err := json.Unmarshal([]byte(jsonStr), &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func MapToJsonStr(m map[string]string) (string, error) {
	jsonByte, err := json.Marshal(m)
	if err != nil {
		return "", err
	}
	return string(jsonByte), nil
}
