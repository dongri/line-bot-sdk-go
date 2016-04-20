# The LINE bot SDK for Go (Golang)

## Start using it

* Download and install it:

```go
$ go get github.com/dongri/line-bot-sdk-go
```

* Import it in your code:

```go
$ import line "github.com/dongri/line-bot-sdk-go"
```

## Examples

```go
package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/dongri/line-bot-sdk-go"
)

var botClient *LineBot.Client

func main() {
	channelID := os.Getenv("LINE_CHANNEL_ID")
	channelSecret := os.Getenv("LINE_CHANNEL_SECRET")
	mid := os.Getenv("LINE_MID")
	proxyURL := getProxyURL() // can set nil if not need

	botClient = LineBot.NewClient(LineBot.EndPoint, channelID, channelSecret, mid, proxyURL)

	// EventHandler
	var myEvent LineBot.EventHandler = NewEventHandler()
	botClient.SetEventHandler(myEvent)

	http.HandleFunc("/callback", callbackHandler)
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("=== callback ===")
}

func getProxyURL() *url.URL {
	proxyURL, err := url.Parse(os.Getenv("PROXY_URL"))
	if err != nil {
		log.Fatal(err)
	}
	return proxyURL
}

// BotEventHandler ...
type BotEventHandler struct {
}

// NewEventHandler ...
func NewEventHandler() *BotEventHandler {
	return &BotEventHandler{}
}

// OnAddedAsFriendOperation ...
func (be *BotEventHandler) OnAddedAsFriendOperation(mids []string) {
	botClient.SendText(mids, "友達追加してくれてありがとうね！")
}

// OnBlockedAccountOperation ...
func (be *BotEventHandler) OnBlockedAccountOperation(mids []string) {
	botClient.SendText(mids, "あらら,,, (このメッセージは届かない)")
}

// OnTextMessage ...
func (be *BotEventHandler) OnTextMessage(from, text string) {
	botClient.SendText([]string{from}, text)
}

// OnImageMessage ...
func (be *BotEventHandler) OnImageMessage(from string) {
	log.Print("=== Received Image ===")
}

// OnVideoMessage ...
func (be *BotEventHandler) OnVideoMessage(from string) {
	log.Print("=== Received Video ===")
}

// OnAudioMessage ...
func (be *BotEventHandler) OnAudioMessage(from string) {
	log.Print("=== Received Audio ===")
}

// OnLocationMessage ...
func (be *BotEventHandler) OnLocationMessage(from, title, address string, latitude, longitude float64) {
	log.Print("=== Received Location ===")
}

// OnStickerMessage ...
func (be *BotEventHandler) OnStickerMessage(from, stickerPackageID, stickerID, stickerVersion, stickerText string) {
	log.Print("=== Received Sticker ===")
}

// OnContactMessage ...
func (be *BotEventHandler) OnContactMessage(from, MID, displayName string) {
	log.Print("=== Received Contact ===")
}
```

#### Get user Profile

```go
fromUser, err := botClient.GetUserProfiles(from)
if err != nil {
	log.Print(err)
}
displayName := fromUser.Contacts[0].DisplayName
```

#### Send Text

```go
botClient.SendText([]string{from}, "Hello!")
```

#### Send Image

```go
originalContentURL := "http://example.com//robot_original.jpg"
previewImageURL := "http://example.com//robot_preview.jpg"
botClient.SendImage([]string{from}, originalContentURL, previewImageURL)
```

#### Send Video

```go
videoOriginalContentURL := "https://example.com/video.mp4"
videoPreviewImageURL := "http://example.com/video.png"
botClient.SendVideo([]string{from}, videoOriginalContentURL, videoPreviewImageURL)
```

#### Send Audio

```go
audioOriginalContentURL := "https://example.com/test.mp3"
audlen := "240000"
botClient.SendAudio([]string{from}, audioOriginalContentURL, audlen)
```

#### Send Location

```go
address := "Minato-ku, Tokyo 107-0062"
latitude := 35.665525
longitude := 139.717945
title := "俺んち"
botClient.SendLocation([]string{from}, address, latitude, longitude, title)
```

### Send Sticker

```go
stkID := "2"
stkpkgID := "1"
stkVer := "100"
stkText := "happy"
botClient.SendSticker([]string{from}, stkID, stkpkgID, stkVer, stkText)
```

#### Send Contact

```go
botClient.SendContact([]string{from}, from, displayName)
```

#### Send Rich Message

```go
func getMarkupJSON() string {
	return "{\"canvas\": {\"width\": 1024,\"height\": 576,\"initialScene\": \"scene1\"},\"images\": {\"image1\": {\"x\": 0,\"y\": 0,\"w\": 1024,\"h\": 576}},\"actions\": {\"openHomepage\": {\"type\": \"web\",\"text\": \"Open link1.\",\"params\": {\"linkUri\": \"http://dongri.github.io/\"}},\"showItem\": {\"type\": \"web\",\"text\": \"Open link2.\",\"params\": {\"linkUri\": \"https://dongri.github.io/post/2016-02-22-the-programmer-hierarchy/\"}}},\"scenes\": {\"scene1\": {\"draws\": [{\"image\": \"image1\",\"x\": 0,\"y\": 0,\"w\": 1024,\"h\": 576}],\"listeners\": [{\"type\": \"touch\",\"params\": [0, 0, 1024, 250],\"action\": \"openHomepage\"}, {\"type\": \"touch\",\"params\": [0, 250, 1024, 326],\"action\": \"showItem\"}]}}}"
}

downloadURL := "https://farm1.staticflickr.com/715/22658699705_7591e8d0a6_b.jpg"
altText := "リスト画面に表示される文字列"
markupJSON := getMarkupJSON()
botClient.SendRichMessage([]string{from}, downloadURL, altText, markupJSON)
```

#### Send MultipleMessage

```go
messageNotified := 0
var contents []line.Content
textContent := new(line.Content)
textContent.Text = "hoge"
contents = append(contents, *textContent)

imageContent := new(line.Content)
imageContent.OriginalContentURL = "https://farm1.staticflickr.com/715/22658699705_7591e8d0a6_b.jpg"
contents = append(contents, *imageContent)
botClient.SendMultipleMessage([]string{from}, messageNotified, contents)
```
