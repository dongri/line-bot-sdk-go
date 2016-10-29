# LINE Bot Messaging API SDK for Go (Golang)

## Start using it

* Download and install it:

```go
$ go get github.com/dongri/line-bot-sdk-go
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
func (be *BotEventHandler) OnFollowEvent(source linebot.EventSource, replyToken string) {
	log.Print(source.UserID + "=== フォローされた ===")
	// source.UserID と Token を保存してnotifyで使える
	message := linebot.NewTextMessage("Hello!")
	result, err := botClient.ReplyMessage(replyToken, message)
	fmt.Println(result)
	fmt.Println(err)
}

// OnUnFollowEvent ...
func (be *BotEventHandler) OnUnFollowEvent(source linebot.EventSource) {
	log.Print(source.UserID + "=== ブロックされた ===")
}

// OnJoinEvent ...
func (be *BotEventHandler) OnJoinEvent(source linebot.EventSource, replyToken string) {
	message := linebot.NewTextMessage("Room, Group 招待ありがとう!")
	result, err := botClient.ReplyMessage(replyToken, message)
	fmt.Println(result)
	fmt.Println(err)
}

// OnLeaveEvent ...
func (be *BotEventHandler) OnLeaveEvent(source linebot.EventSource) {
	log.Print("=== Groupから蹴られた ===")
}

// OnPostbackEvent ...
func (be *BotEventHandler) OnPostbackEvent(source linebot.EventSource, replyToken, postbackData string) {
	message := linebot.NewTextMessage("「" + postbackData + "」を選択したね！")
	result, err := botClient.ReplyMessage(replyToken, message)
	fmt.Println(result)
	fmt.Println(err)
}

// OnBeaconEvent ...
func (be *BotEventHandler) OnBeaconEvent(source linebot.EventSource, replyToken, beaconHwid, beaconYype string) {
	log.Print("=== Beacon Event ===")
}

// OnTextMessage ...
func (be *BotEventHandler) OnTextMessage(source linebot.EventSource, replyToken, text string) {
	message := linebot.NewTextMessage(text + "じゃねぇよ！")
	result, err := botClient.ReplyMessage(replyToken, message)
	fmt.Println(result)
	fmt.Println(err)
}

// OnImageMessage ...
func (be *BotEventHandler) OnImageMessage(source linebot.EventSource, replyToken, id string) {
	originalContentURL := "https://dl.dropboxusercontent.com/u/358152/linebot/resource/gohper.jpg"
	previewImageURL := "https://dl.dropboxusercontent.com/u/358152/linebot/resource/gohper.jpg"
	message := linebot.NewImageMessage(originalContentURL, previewImageURL)
	result, err := botClient.ReplyMessage(replyToken, message)
	fmt.Println(result)
	fmt.Println(err)
}

// OnVideoMessage ...
func (be *BotEventHandler) OnVideoMessage(source linebot.EventSource, replyToken, id string) {
	originalContentURL := "https://dl.dropboxusercontent.com/u/358152/linebot/resource/video-original.mp4"
	previewImageURL := "https://dl.dropboxusercontent.com/u/358152/linebot/resource/video-preview.png"
	message := linebot.NewVideoMessage(originalContentURL, previewImageURL)
	result, err := botClient.ReplyMessage(replyToken, message)
	fmt.Println(result)
	fmt.Println(err)
}

// OnAudioMessage ...
func (be *BotEventHandler) OnAudioMessage(source linebot.EventSource, replyToken, id string) {
	originalContentURL := "https://dl.dropboxusercontent.com/u/358152/linebot/resource/ok.m4a"
	duration := 1000
	message := linebot.NewAudioMessage(originalContentURL, duration)
	result, err := botClient.ReplyMessage(replyToken, message)
	fmt.Println(result)
	fmt.Println(err)
}

// OnLocationMessage ...
func (be *BotEventHandler) OnLocationMessage(source linebot.EventSource, replyToken string, title, address string, latitude, longitude float64) {
	title = "Disney Resort"
	address = "〒279-0031 千葉県浦安市舞浜１−１"
	lat := 35.632211
	lon := 139.881234
	message := linebot.NewLocationMessage(title, address, lat, lon)
	result, err := botClient.ReplyMessage(replyToken, message)
	fmt.Println(result)
	fmt.Println(err)
}

// OnStickerMessage ...
func (be *BotEventHandler) OnStickerMessage(source linebot.EventSource, replyToken, packageID, stickerID string) {
	message := linebot.NewStickerMessage("1", "1")
	result, err := botClient.ReplyMessage(replyToken, message)
	fmt.Println(result)
	fmt.Println(err)
}
```

## Demo

<img src="https://github.com/dongri/line-bot-sdk-go/blob/master/examples/QR.png" width="185">
