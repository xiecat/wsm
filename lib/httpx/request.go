package httpx

import (
	"bytes"
	"github.com/pkg/errors"
	"github.com/xiecat/wsm/lib/shell"
	"github.com/xiecat/wsm/lib/shell/godzilla"
	"net/http"
)

type ReqClient struct {
	client *http.Client
	header map[string]string
	script shell.ScriptType
	crypto godzilla.CrypticType
}

func NewClient(proxyUrl string, header map[string]string, script shell.ScriptType, crypt godzilla.CrypticType) *ReqClient {
	if len(header["User-Agent"]) == 0 {
		header = randUa(header)
	}
	return &ReqClient{
		client: newClient(proxyUrl),
		header: header,
		script: script,
		crypto: crypt,
	}
}

func randUa(header map[string]string) map[string]string {
	header["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.93 Safari/537.36"
	return header
}

func (r *ReqClient) DoHttpRequest(url string, data []byte) (*HttpResponse, error) {
	resp, err := r.sendPayload(url, data)
	if err != nil {
		return nil, err
	}
	result, err := newHttpResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *ReqClient) sendPayload(u string, data []byte) (*http.Response, error) {
	request, err := http.NewRequest(http.MethodPost, u, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	for k, v := range r.header {
		request.Header.Set(k, v)
	}
	resp, err := r.client.Do(request)
	if err != nil {
		return nil, errors.Wrap(err, "http client request fail")

	}
	return resp, nil
}
