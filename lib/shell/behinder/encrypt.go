package behinder

import (
	"encoding/base64"
	"github.com/Go0p/wsm/lib/encrypt"
	"github.com/Go0p/wsm/lib/shell"
	"log"
)

func Encrypto(bs, key []byte, encryptType int, types shell.ScriptType) []byte {
	var result []byte
	if types == shell.JavaScript {
		result = encryptForJava(bs, key)
	} else if types == shell.PhpScript {
		result = encryptForPhp(bs, key, encryptType)
		if encryptType == ENCRYPT_TYPE_AES {
			result = []byte(base64.StdEncoding.EncodeToString(result))
		}
	} else if types == shell.CsharpScript {
		result = encryptForCSharp(bs, key)
	} else if types == shell.AspScript {
		result = encryptForAsp(bs, key)
	}
	return result
}

func Decrypto(raw, key []byte, types shell.ScriptType, notEncrypt string, encryptType, prefixLen, suffixLen int) []byte {
	var result []byte
	var targetBts []byte
	if (suffixLen != 0 || prefixLen != 0) && len(raw)-prefixLen >= suffixLen {
		targetBts = raw[prefixLen : len(raw)-suffixLen]
	} else {
		targetBts = raw
	}
	if types == shell.JavaScript {
		if notEncrypt != "true" {
			result = decryptForJava(targetBts, key)
		} else {
			result = targetBts
		}
	} else if types == shell.PhpScript {
		result = decryptForPhp(targetBts, key, encryptType)
	} else if types == shell.CsharpScript {
		result = decryptForCSharp(targetBts, key)
	} else if types == shell.AspScript {
		result = decryptForAsp(targetBts, key)
	}
	return result
}

func encryptForCSharp(bs []byte, key []byte) []byte {
	encrypted, err := encrypt.AESCBCEncrypt(bs, key, key)
	if err != nil {
		log.Println(err)
	}
	return encrypted
}

func decryptForCSharp(bs []byte, key []byte) []byte {
	decrypted, err := encrypt.AESCBCDecrypt(bs, key, key)
	if err != nil {
		log.Println(err)
	}
	return decrypted
}

func encryptForJava(bs, key []byte) []byte {
	encrypted, err := encrypt.AESECBEncrypt(bs, key)
	if err != nil {
		log.Println(err)
	}
	return encrypted
}

func decryptForJava(ct, key []byte) []byte {
	decrypted, err := encrypt.AESECBDecrypt(ct, key)
	if err != nil {
		log.Println(err)
	}
	return decrypted
}

func encryptForPhp(pt, key []byte, encryptType int) []byte {
	var encrypted []byte
	var err error
	if encryptType == ENCRYPT_TYPE_AES {
		encrypted, err = encrypt.AESCBCEncrypt(pt, key, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
		if err != nil {
			log.Println(err)
		}
	} else if encryptType == ENCRYPT_TYPE_XOR {
		encrypted = encryptForAsp(pt, key)
	}
	return encrypted
}
func decryptForPhp(ct, key []byte, encryptType int) []byte {
	var decrypted []byte
	if encryptType == ENCRYPT_TYPE_AES {
		c, err := base64.StdEncoding.DecodeString(string(ct))
		if err != nil {
			log.Print("PHP CBC Resp Base64,")
			//panic(err)
		}
		decrypted, err = encrypt.AESCBCDecrypt(c, key, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
		if err != nil {
			log.Println(err)
		}
	} else if encryptType == ENCRYPT_TYPE_XOR {
		decrypted = decryptForAsp(ct, key)
	}
	return decrypted
}

func encryptForAsp(bs []byte, key []byte) []byte {
	return encrypt.Xor(bs, key)
}

func decryptForAsp(bs, key []byte) []byte {
	return encrypt.Xor(bs, key)
}
