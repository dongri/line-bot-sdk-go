package linebot

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// URL Shortener
const (
	GoogleURLShortener = "https://www.googleapis.com/urlshortener/v1/url"
)

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
	//body, err := DoRequest(req)
	// if err != nil {
	// 	return longURL
	// }
	// var result = map[string]interface{}{}
	// if err := json.Unmarshal(body, &result); err != nil {
	// 	return longURL
	// }
	// return result["id"].(string)
	return ""
}
