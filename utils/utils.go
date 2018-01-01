package utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func NewHTTPRequest(method, url string, body []byte, cookie *http.Cookie) (*http.Request, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if nil != err {
		req = nil
	}

	if nil != cookie {
		req.AddCookie(cookie)
	}

	return req, err
}

func MakeHTTPQuery(req *http.Request) (*http.Response, error) {

	client := &http.Client{}
	resp, err := client.Do(req)

	return resp, err
}

func PrintResponseBody(resp *http.Response) {
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)
}
