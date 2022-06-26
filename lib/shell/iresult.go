package shell

type IResult interface {
	Parser()
}

type Result map[string]string

func (r Result) Parser() {

}
