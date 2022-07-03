package httpx

type Result struct {
	RawBody []byte
	Status  int
}

func NewResult(data []byte) *Result {
	return &Result{RawBody: data}
}

func (r Result) Parser() {
	//TODO implement me
	panic("implement me")
}
