package shell

const (
	Ping = iota
	BasicInfo
	Exec
)

type IParams interface {
	SetDefaultAndCheckValue() error
	//ToString() string
	//ToMap() string
}
