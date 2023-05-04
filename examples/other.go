package main

import (
	"encoding/json"
	"io"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// DoRequest ...
func DoRequest(req *http.Request, proxyURL *url.URL) map[string]interface{} {
	client := &http.Client{
		Timeout: time.Duration(30 * time.Second),
	}
	res, err := client.Do(req)
	if err != nil {
		log.Print(err)
	}
	defer res.Body.Close()
	var result map[string]interface{}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Print(err)
	}
	if err := json.Unmarshal(body, &result); err != nil {
		log.Print(err)
	}
	return result
}

// GetImageFromWeb ...
func GetImageFromWeb() string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	offset := strconv.Itoa(rand.Intn(1000))

	//blogName := "mincang.tumblr.com"
	blogName := "kawaii-sexy-love.tumblr.com"

	resp, err := http.Get("http://api.tumblr.com/v2/blog/" + blogName + "/posts/photo?api_key=3bllCUqUSGV73R3O7BKWq5m9mIERjzMORIPnkxv90s3KsqOCH4&limit=1&offset=" + offset)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	data := make(map[string]interface{})
	body, _ := io.ReadAll(resp.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err)
	}

	response := data["response"].(map[string]interface{})
	posts := response["posts"].([]interface{})
	post := posts[0].(map[string]interface{})
	photos := post["photos"].([]interface{})
	photo := photos[0].(map[string]interface{})
	altSizes := photo["alt_sizes"].([]interface{})
	firstSize := altSizes[0].(map[string]interface{})
	imageURL := firstSize["url"].(string)
	return imageURL
}
