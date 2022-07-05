package httpx

import (
	"bytes"
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
		return nil, err
	}
	raw = bytes.TrimRight(raw, "\r\n\r\n")
	httpResp.RawBody = raw
	return httpResp, nil
}
