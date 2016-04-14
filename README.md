# LINE Bot Client for Golang

# Start using it

1. Download and install it:

```go
$ go get github.com/dongri/line-bot-client
```

2. Import it in your code:

```go
$ import line "github.com/dongri/line-bot-client"
```

# Examples

```go
package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	line "github.com/dongri/line-bot-client"
)

var botClient *line.Client

func main() {
	channelID := os.Getenv("LINE_CHANNEL_ID")
	channelSecret := os.Getenv("LINE_CHANNEL_SECRET")
	mid := os.Getenv("LINE_MID")
	proxyURL := getProxyURL() // can set nil if not need

	botClient = line.NewClient(line.EndPoint, channelID, channelSecret, mid, proxyURL)

	http.HandleFunc("/callback", callbackHandler)
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	m, err := botClient.ReceiveMessageAndOperation(r.Body)
	if err != nil {
		log.Print(err)
	}
	for _, result := range m.Result {
		from := result.Content.From

		if result.Content.OpType == line.OpTypeAdded {
			from = result.Content.Params[0]
			botClient.SendText([]string{from}, "友達追加してくれてありがとうね！")
			log.Print("Added as friend")
			return
		}

		if result.Content.OpType == line.OpTypeBlocked {
			from = result.Content.Params[0]
			botClient.SendText([]string{from}, "あらら、いつでもブロック解除できますよー")
			log.Print("Blocked account")
			return
		}

		// Get User Profile
		fromUser, err := botClient.GetUserProfiles(from)
		if err != nil {
			log.Print(err)
		}
		displayName := fromUser.Contacts[0].DisplayName

		// Send Text
		text := result.Content.Text
		sentResult, err := botClient.SendText([]string{from}, text+"\n\nBy "+displayName)
		if err != nil {
			log.Print(err)
		}
		if len(sentResult.Failed) == 0 {
			log.Print("Failed")
			return
		}

		// Send Sticker
		STKID := "2"
		STKPKGID := "1"
		STKVER := "100"
		STKTXT := "Text"
		sentResult, err = botClient.SendSticker([]string{from}, STKID, STKPKGID, STKVER, STKTXT)
		if err != nil {
			log.Print(err)
		}
		if len(sentResult.Failed) == 0 {
			log.Print("Failed")
			return
		}

		// Send Image
		imageOriginalContentURL := "https://example.com/test_original.png"
		imagePreviewImageURL := "https://example.com/test_preview.png"
		sentResult, err = botClient.SendImage([]string{from}, imageOriginalContentURL, imagePreviewImageURL)
		if err != nil {
			log.Print(err)
		}
		if len(sentResult.Failed) == 0 {
			log.Print("Failed")
			return
		}

		// Send Video ....
		videoOriginalContentURL := "https://example.com/test.mp4"
		videoPreviewImageURL := "http://example.com/test.png"
		sentResult, err = botClient.SendVideo([]string{from}, videoOriginalContentURL, videoPreviewImageURL)
		if err != nil {
			log.Print(err)
		}
		if len(sentResult.Failed) == 0 {
			log.Print("Failed")
			return
		}

		// Send Audio ....
		audioOriginalContentURL := "https://example.com/test.mp3"
		audlen := "240000"
		sentResult, err = botClient.SendAudio([]string{from}, audioOriginalContentURL, audlen)
		if err != nil {
			log.Print(err)
		}
		if len(sentResult.Failed) == 0 {
			log.Print("Failed")
			return
		}

		// Send Contact ....
		sentResult, err = botClient.SendContact([]string{from}, from, displayName)
		if err != nil {
			log.Print(err)
		}
		if len(sentResult.Failed) == 0 {
			log.Print("Failed")
			return
		}

		// Send Rich Message
		downloadURL := "https://farm1.staticflickr.com/715/22658699705_7591e8d0a6_b.jpg"
		altText := "リスト画面に表示される文字列"
		markupJSON := getMarkupJSON()
		sentResult, err = botClient.SendRichMessage([]string{from}, downloadURL, altText, markupJSON)
		if err != nil {
			log.Print(err)
		}
		if len(sentResult.Failed) == 0 {
			log.Print("Failed")
			return
		}

		// Send MultipleMessage
		messageNotified := 0
		var contents []line.Content
		textContent := new(line.Content)
		textContent.Text = "hoge"
		contents = append(contents, *textContent)

		imageContent := new(line.Content)
		imageContent.OriginalContentURL = "https://farm1.staticflickr.com/715/22658699705_7591e8d0a6_b.jpg"
		contents = append(contents, *imageContent)
		botClient.SendMultipleMessage([]string{from}, messageNotified, contents)

	}

	log.Print("Success")
}

func getProxyURL() *url.URL {
	proxyURL, err := url.Parse(os.Getenv("PROXY_URL"))
	if err != nil {
		log.Fatal(err)
	}
	return proxyURL
}

func getMarkupJSON() string {
	return "{\"canvas\": {\"width\": 1024,\"height\": 576,\"initialScene\": \"scene1\"},\"images\": {\"image1\": {\"x\": 0,\"y\": 0,\"w\": 1024,\"h\": 576}},\"actions\": {\"openHomepage\": {\"type\": \"web\",\"text\": \"Open link1.\",\"params\": {\"linkUri\": \"http://dongri.github.io/\"}},\"showItem\": {\"type\": \"web\",\"text\": \"Open link2.\",\"params\": {\"linkUri\": \"https://dongri.github.io/post/2016-02-22-the-programmer-hierarchy/\"}}},\"scenes\": {\"scene1\": {\"draws\": [{\"image\": \"image1\",\"x\": 0,\"y\": 0,\"w\": 1024,\"h\": 576}],\"listeners\": [{\"type\": \"touch\",\"params\": [0, 0, 1024, 250],\"action\": \"openHomepage\"}, {\"type\": \"touch\",\"params\": [0, 250, 1024, 326],\"action\": \"showItem\"}]}}}"
}
```
