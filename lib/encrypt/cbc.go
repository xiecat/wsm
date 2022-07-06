package encrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"errors"
)

func AESCBCEncrypt(src, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(src) == 0 {
		return nil, errors.New("plain content empty")
	}
	cbc := cipher.NewCBCEncrypter(block, iv)
	content := pKCS5Padding(src, block.BlockSize())
	crypted := make([]byte, len(content))
	cbc.CryptBlocks(crypted, content)

	return crypted, nil
}

func AESCBCDecrypt(src []byte, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(src) == 0 {
		return nil, errors.New("plain content empty")
	}
	cbc := cipher.NewCBCDecrypter(block, iv)
	decrypted := make([]byte, len(src))
	cbc.CryptBlocks(decrypted, src)

	return pKCS5Trimming(decrypted), nil
}

func pKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func pKCS5Trimming(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}
