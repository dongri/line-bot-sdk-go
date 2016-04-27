package linebot

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

// DoRequest ...
func DoRequest(req *http.Request, proxyURL *url.URL) ([]byte, error) {
	client := &http.Client{
		Timeout: time.Duration(30 * time.Second),
	}
	if proxyURL != nil {
		client.Transport = &http.Transport{Proxy: http.ProxyURL(proxyURL)}
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	return body, err
}
