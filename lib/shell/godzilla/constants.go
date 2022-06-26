package godzilla

type CrypticType string

const (
	JAVA_AES_BASE64   CrypticType = "JAVA_AES_BASE64"
	JAVA_AES_RAW      CrypticType = "JAVA_AES_RAW"
	CSHARP_AES_BASE64 CrypticType = "CSHARP_AES_BASE64"
	CSHARP_AES_RAW    CrypticType = "CSHARP_AES_RAW"
	PHP_XOR_BASE64    CrypticType = "PHP_XOR_BASE64"
	PHP_XOR_RAW       CrypticType = "PHP_XOR_RAW"
)
