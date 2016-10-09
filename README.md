# LINE Bot Messaging API SDK for Go (Golang)

## Start using it

* Download and install it:

```go
$ go get github.com/dongri/line-bot-sdk-go
```

* Import it in your code:

```go
$ import "github.com/dongri/line-bot-sdk-go/linebot"
```

## Examples

```go
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dongri/line-bot-sdk-go/linebot"
)

var botClient *linebot.Client

func main() {
	channelAccessToken := os.Getenv("LINE_CHANNEL_ACCESSTOKEN")
	channelSecret := os.Getenv("LINE_CHANNEL_SECRET")

	botClient = linebot.NewClient(channelAccessToken)
	botClient.SetChannelSecret(channelSecret)

	// EventHandler
	var myEvent linebot.EventHandler = NewEventHandler()
	botClient.SetEventHandler(myEvent)

	http.Handle("/callback", linebot.Middleware(http.HandlerFunc(callbackHandler)))
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("=== callback ===")
}

// BotEventHandler ...
type BotEventHandler struct{}

// NewEventHandler ...
func NewEventHandler() *BotEventHandler {
	return &BotEventHandler{}
}

// OnFollowEvent ...
func (be *BotEventHandler) OnFollowEvent(replyToken string) {
	message := linebot.NewTextMessage("Hello!")
	botClient.ReplyMessage(replyToken, message)
}

// OnUnFollowEvent ...
func (be *BotEventHandler) OnUnFollowEvent() {
	log.Print("=== ブロックされた ===")
}

// OnJoinEvent ...
func (be *BotEventHandler) OnJoinEvent(replyToken string) {
	message := linebot.NewTextMessage("Room, Group 招待ありがとう!")
	botClient.ReplyMessage(replyToken, message)
}

// OnLeaveEvent ...
func (be *BotEventHandler) OnLeaveEvent() {
	log.Print("=== Groupから蹴られた ===")
}

// OnPostbackEvent ...
func (be *BotEventHandler) OnPostbackEvent(replyToken, postbackData string) {
	message := linebot.NewTextMessage("「" + postbackData + "」を選択したね！")
	botClient.ReplyMessage(replyToken, message)
}

// OnBeaconEvent ...
func (be *BotEventHandler) OnBeaconEvent(replyToken, beaconHwid, beaconYype string) {
	log.Print("=== Beacon Event ===")
}

// OnTextMessage ...
func (be *BotEventHandler) OnTextMessage(replyToken, text string) {
	message := linebot.NewTextMessage(text)
	botClient.ReplyMessage(replyToken, message)
}

// OnImageMessage ...
func (be *BotEventHandler) OnImageMessage(replyToken, id string) {
	originalContentURL := "https://dl.dropboxusercontent.com/u/358152/linebot/resource/gohper.jpg"
	previewImageURL := "https://dl.dropboxusercontent.com/u/358152/linebot/resource/gohper.jpg"
	message := linebot.NewImageMessage(originalContentURL, previewImageURL)
	botClient.ReplyMessage(replyToken, message)
}

// OnVideoMessage ...
func (be *BotEventHandler) OnVideoMessage(replyToken, id string) {
	originalContentURL := "https://dl.dropboxusercontent.com/u/358152/linebot/resource/video-original.mp4"
	previewImageURL := "https://dl.dropboxusercontent.com/u/358152/linebot/resource/video-preview.png"
	message := linebot.NewVideoMessage(originalContentURL, previewImageURL)
	botClient.ReplyMessage(replyToken, message)
}

// OnAudioMessage ...
func (be *BotEventHandler) OnAudioMessage(replyToken, id string) {
	originalContentURL := "https://dl.dropboxusercontent.com/u/358152/linebot/resource/ok.m4a"
	duration := 1000
	message := linebot.NewAudioMessage(originalContentURL, duration)
	botClient.ReplyMessage(replyToken, message)
}

// OnLocationMessage ...
func (be *BotEventHandler) OnLocationMessage(replyToken string, latitude, longitude float64) {
	title := "Disney Resort"
	address := "〒279-0031 千葉県浦安市舞浜１−１"
	lat := 35.632211
	lon := 139.881234
	message := linebot.NewLocationMessage(title, address, lat, lon)
	botClient.ReplyMessage(replyToken, message)
}

// OnStickerMessage ...
func (be *BotEventHandler) OnStickerMessage(replyToken, stickerID string) {
	message := linebot.NewStickerMessage("1", "1")
	botClient.ReplyMessage(replyToken, message)
}
```

## Demo

<img src="https://github.com/dongri/line-bot-sdk-go/blob/master/examples/QR.png" width="185">
