// Copyright 2019 Conflux Foundation. All rights reserved.
// Conflux is free software and distributed under GNU General Public License.
// See http://www.gnu.org/licenses/

package mock

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
)

type HttpClientMock struct {
	GetHandler map[string]string
}

func (c *HttpClientMock) SetHandler(path, responseBody string) {
	if c.GetHandler == nil {
		c.GetHandler = make(map[string]string)
	}
	c.GetHandler[path] = responseBody
}

func (c *HttpClientMock) Get(url string) (resp *http.Response, err error) {
	bodyStr, ok := c.GetHandler[url]
	if !ok {
		bodyStr, ok = c.GetHandler[""]
		if !ok {
			return nil, errors.New("Not exist mock response string of url :" + url)
		}
	}

	// var rsp http.Response
	rsp := http.Response{
		Body: ioutil.NopCloser(bytes.NewBufferString(bodyStr)),
	}
	return &rsp, nil
}
