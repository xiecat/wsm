package shell

type IResult interface {
	Parser()
	GetRaw() string
}
