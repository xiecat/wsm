package shell

const (
	Ping = iota
	BasicInfo
	Exec
)

type IParams interface {
	Check()
	//Get(key string) string
}
