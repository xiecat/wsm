// Package tools 解密一些冰蝎、哥斯拉的流量
package tools

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go0p/wsm/lib/encrypt"
	"github.com/go0p/wsm/lib/shell"
	"github.com/go0p/wsm/lib/shell/behinder"
	"github.com/go0p/wsm/lib/utils"
	"strings"
)

// DecryptBehinderPcap 解密冰蝎流量
func DecryptBehinderPcap(src string, pass string, script shell.ScriptType) (string, string, error) {
	if strings.Contains(src, "\n") {
		src = strings.ReplaceAll(src, "\n", "")
	}
	key := utils.SecretKey(pass)
	raw := strings.SplitN(src, "485454502f312e3120", 2)
	req, resp := raw[0], "485454502f312e3120"+raw[1]
	reqStr, err := decryptRequest(req, key, script)
	if err != nil {
		return "", "", err
	}
	respStr, err := decryptResponse(resp, key, script)
	if err != nil {
		return "", "", err
	}
	return string(reqStr), respStr, nil
}

func decryptRequest(req string, key []byte, script shell.ScriptType) ([]byte, error) {
	raw := strings.SplitN(req, "0d0a0d0a", 2)
	body, err := hex.DecodeString(raw[1])
	if err != nil {
		return nil, err
	}
	decrypto, err := restitutePayload(body, key, script)
	if err != nil {
		return nil, err
	}
	return decrypto, nil
}

func decryptResponse(resp string, key []byte, script shell.ScriptType) (string, error) {
	raw := strings.SplitN(resp, "0d0a0d0a", 2)
	body, err := hex.DecodeString(raw[1])
	if err != nil {
		return "", err
	}
	decrypto, err := behinder.Decrypto(body, key, script, "false", 0, 0, 0)
	if err != nil {
		return "", err
	}
	result := make(map[string]string, 2)
	if err = json.Unmarshal(decrypto, &result); err == nil {
		for k, v := range result {
			value, err := base64.StdEncoding.DecodeString(v)
			if err == nil {
				result[k] = string(value)
			}
		}
	}
	str, err := utils.MapToJsonStr(result)
	if err != nil {
		return "", err
	}
	return str, nil
}

func restitutePayload(body, key []byte, script shell.ScriptType) ([]byte, error) {
	switch script {
	case shell.JavaScript, shell.JspxScript:
		deBody, err := base64.StdEncoding.DecodeString(string(body))
		if err != nil {
			return nil, err
		}
		decrypt, err := encrypt.AESECBDecrypt(deBody, key)
		if err != nil {
			return nil, err
		}
		return decrypt, nil
	case shell.PhpScript:
		deBody, err := base64.StdEncoding.DecodeString(string(body))
		if err != nil {
			return nil, err
		}
		decrypt, err := encrypt.AESCBCDecrypt(deBody, key, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
		if err != nil {
			decrypt = encrypt.Xor(deBody, key)
		}
		return decrypt, nil
	case shell.CsharpScript:
		decrypt, err := encrypt.AESCBCDecrypt(body, key, key)
		if err != nil {
			return nil, err
		}
		return decrypt, nil
	case shell.AspScript:
		decrypt := encrypt.Xor(body, key)
		return decrypt, nil
	}
	return nil, errors.New(fmt.Sprintf("get %s payload error", script))
}
