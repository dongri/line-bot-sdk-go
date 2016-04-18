package LineBot

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

// URL Shortener
const (
	GoogleURLShortener = "https://www.googleapis.com/urlshortener/v1/url"
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

// URLShortener return goo.gl url
func URLShortener(APIKey, longURL string) string {
	apiURL := GoogleURLShortener + "?key=" + APIKey
	type params struct {
		LongURL string `json:"longUrl"`
	}
	p := params{LongURL: longURL}
	b, err := json.Marshal(p)
	if err != nil {
		return longURL
	}
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(b))
	if err != nil {
		return longURL
	}
	req.Header.Add("Content-Type", "application/json; charset=UTF-8")
	body, err := DoRequest(req, nil)
	if err != nil {
		return longURL
	}
	var result = map[string]interface{}{}
	if err := json.Unmarshal(body, &result); err != nil {
		return longURL
	}
	return result["id"].(string)
}
