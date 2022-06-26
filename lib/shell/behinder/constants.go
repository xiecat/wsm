package behinder

const (
	ENCRYPT_TYPE_AES = iota
	ENCRYPT_TYPE_XOR
)

type PayloadName string

const (
	EchoGo PayloadName = "EchoGo"
)
