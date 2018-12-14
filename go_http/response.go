package go_http

import (
	"io/ioutil"
	"net/http"
)

type Response struct {
	Resp *http.Response
}

func (r Response) GetBodyString() string {
	if r.Resp.Body != nil {
		by, err := ioutil.ReadAll(r.Resp.Body)
		defer r.Resp.Body.Close()
		if err != nil {
			return err.Error()
		}
		return string(by)
	}
	return ""
}
