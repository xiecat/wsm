package httpx

import (
	"bytes"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

type HttpResponse struct {
	RawBody []byte
	*http.Response
}

func newHttpResponse(resp *http.Response) (*HttpResponse, error) {
	httpResp := &HttpResponse{
		Response: resp,
	}
	defer resp.Body.Close()
	raw, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "read body fail")
	}
	raw = bytes.TrimRight(raw, "\r\n\r\n")
	httpResp.RawBody = raw
	return httpResp, nil
}
