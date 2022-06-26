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

//type Params map[string]string
//
//func (p Params) Check() {
//}

//func (p Params) Get(key string) string {
//	return ""
//}
