package encrypt

func Xor(src, key []byte) []byte {
	for i := 0; i < len(src); i++ {
		src[i] = src[i] ^ key[(i+1)&15]
	}
	return src
}
