package behinder

import (
	"encoding/base64"
	"errors"
	"github.com/go0p/wsm/lib/encrypt"
	"github.com/go0p/wsm/lib/shell"
)

func Encrypto(bs, key []byte, encryptType int, types shell.ScriptType) ([]byte, error) {
	var result []byte
	var err error
	if types == shell.JavaScript || types == shell.JspxScript {
		result, err = encryptForJava(bs, key)
	} else if types == shell.PhpScript {
		result, err = encryptForPhp(bs, key, encryptType)
		if encryptType == ENCRYPT_TYPE_AES {
			result = []byte(base64.StdEncoding.EncodeToString(result))
		}
	} else if types == shell.CsharpScript {
		result, err = encryptForCSharp(bs, key)
	} else if types == shell.AspScript {
		result = encryptForAsp(bs, key)
	}
	return result, err
}

func Decrypto(raw, key []byte, types shell.ScriptType, notEncrypt string, encryptType, prefixLen, suffixLen int) ([]byte, error) {
	var result, targetBts []byte
	var err error
	if (suffixLen != 0 || prefixLen != 0) && len(raw)-prefixLen >= suffixLen {
		targetBts = raw[prefixLen : len(raw)-suffixLen]
	} else {
		targetBts = raw
	}
	if types == shell.JavaScript || types == shell.JspxScript {
		if notEncrypt != "true" {
			result, err = decryptForJava(targetBts, key)
		} else {
			result = targetBts
		}
	} else if types == shell.PhpScript {
		result, err = decryptForPhp(targetBts, key, encryptType)
	} else if types == shell.CsharpScript {
		result, err = decryptForCSharp(targetBts, key)
	} else if types == shell.AspScript {
		result = decryptForAsp(targetBts, key)
	}
	return result, err
}

func encryptForCSharp(bs []byte, key []byte) ([]byte, error) {
	encrypted, err := encrypt.AESCBCEncrypt(bs, key, key)
	if err != nil {
		return nil, err
	}
	return encrypted, nil
}

func decryptForCSharp(bs []byte, key []byte) ([]byte, error) {
	decrypted, err := encrypt.AESCBCDecrypt(bs, key, key)
	if err != nil {
		return nil, err
	}
	return decrypted, nil
}

func encryptForJava(bs, key []byte) ([]byte, error) {
	encrypted, err := encrypt.AESECBEncrypt(bs, key)
	if err != nil {
		return nil, err
	}
	return encrypted, nil
}

func decryptForJava(ct, key []byte) ([]byte, error) {
	decrypted, err := encrypt.AESECBDecrypt(ct, key)
	if err != nil {
		return nil, err
	}
	return decrypted, nil
}

func encryptForPhp(pt, key []byte, encryptType int) ([]byte, error) {
	var encrypted []byte
	var err error
	if encryptType == ENCRYPT_TYPE_AES {
		encrypted, err = encrypt.AESCBCEncrypt(pt, key, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
		if err != nil {
			return nil, err
		}
	} else if encryptType == ENCRYPT_TYPE_XOR {
		encrypted = encryptForAsp(pt, key)
	}
	return encrypted, nil
}
func decryptForPhp(ct, key []byte, encryptType int) ([]byte, error) {
	var decrypted []byte
	if encryptType == ENCRYPT_TYPE_AES {
		c, err := base64.StdEncoding.DecodeString(string(ct))
		if err != nil {
			return nil, err
		}
		decrypted, err = encrypt.AESCBCDecrypt(c, key, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
		if err != nil {
			return nil, err
		}
	} else if encryptType == ENCRYPT_TYPE_XOR {
		decrypted = decryptForAsp(ct, key)
	} else {
		return nil, errors.New("encryptType is not support")
	}
	return decrypted, nil
}

func encryptForAsp(bs []byte, key []byte) []byte {
	return encrypt.Xor(bs, key)
}

func decryptForAsp(bs, key []byte) []byte {
	return encrypt.Xor(bs, key)
}
