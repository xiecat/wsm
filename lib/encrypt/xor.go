package encrypt

func Xor(src, key []byte) []byte {
	dst := make([]byte, len(src))
	for i := 0; i < len(src); i++ {
		dst[i] = src[i] ^ key[(i+1)&15]
	}
	return dst
}
