package shell

type IResult interface {
	Parser() error
	ToMap() map[string]string
	ToString() string
}
