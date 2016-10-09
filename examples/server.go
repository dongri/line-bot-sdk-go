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

	http.HandleFunc("/", indexHandler)
	http.Handle("/callback", linebot.Middleware(http.HandlerFunc(callbackHandler)))
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Bot")
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
	result, err := botClient.ReplyMessage(replyToken, message)
	fmt.Println(result)
	fmt.Println(err)
}

// OnUnFollowEvent ...
func (be *BotEventHandler) OnUnFollowEvent() {
	log.Print("=== ブロックされた ===")
}

// OnJoinEvent ...
func (be *BotEventHandler) OnJoinEvent(replyToken string) {
	message := linebot.NewTextMessage("Room, Group 招待ありがとう!")
	result, err := botClient.ReplyMessage(replyToken, message)
	fmt.Println(result)
	fmt.Println(err)
}

// OnLeaveEvent ...
func (be *BotEventHandler) OnLeaveEvent() {
	log.Print("=== Groupから蹴られた ===")
}

// OnPostbackEvent ...
func (be *BotEventHandler) OnPostbackEvent(replyToken, postbackData string) {
	message := linebot.NewTextMessage("「" + postbackData + "」を選択したね！")
	result, err := botClient.ReplyMessage(replyToken, message)
	fmt.Println(result)
	fmt.Println(err)
}

// OnBeaconEvent ...
func (be *BotEventHandler) OnBeaconEvent(replyToken, beaconHwid, beaconYype string) {
	log.Print("=== Beacon Event ===")
}

// OnTextMessage ...
func (be *BotEventHandler) OnTextMessage(replyToken, text string) {
	if text == "imagemap" {
		baseURL := "https://dl.dropboxusercontent.com/u/358152/linebot/resource"
		altText := "Golang"
		baseSizeWidth := 520
		baseSizeHeight := 520
		message := linebot.NewImagemapMessage(baseURL, altText, baseSizeWidth, baseSizeHeight,
			linebot.NewImagemapURIAction("https://golang.org", *linebot.NewImagemapArea(0, 0, 520, 520)),
		)
		result, err := botClient.ReplyMessage(replyToken, message)
		fmt.Println(result)
		fmt.Println(err)
	} else if text == "template" {
		templateLabel := "Go"
		templateText := "Hello, Golang!"
		thumbnailImageURL := "https://dl.dropboxusercontent.com/u/358152/linebot/resource/gopher.png"
		actionLabel := "Go to golang.org"
		actionURI := "https://golang.org"
		template := linebot.NewButtonsTemplate(
			thumbnailImageURL, templateLabel, templateText,
			linebot.NewTemplateURIAction(actionLabel, actionURI),
			linebot.NewTemplatePostbackAction("hello postback", "hello こんにちは", "必須じゃない？！"),
			linebot.NewTemplateMessageAction("hello message", "hello こんにちは2"),
		)
		altText := "Go template"
		message := linebot.NewTemplateMessage(altText, template)
		result, err := botClient.ReplyMessage(replyToken, message)
		fmt.Println(result)
		fmt.Println(err)
	} else if text == "confirm" {
		template := linebot.NewConfirmTemplate(
			"Do it?",
			linebot.NewTemplateMessageAction("Yes", "Yes!"),
			linebot.NewTemplateMessageAction("No", "No!"),
		)
		altText := "Confirm template"
		message := linebot.NewTemplateMessage(altText, template)
		result, err := botClient.ReplyMessage(replyToken, message)
		fmt.Println(result)
		fmt.Println(err)
	} else {
		message := linebot.NewTextMessage(text)
		result, err := botClient.ReplyMessage(replyToken, message)
		fmt.Println(result)
		fmt.Println(err)
	}
}

// OnImageMessage ...
func (be *BotEventHandler) OnImageMessage(replyToken, id string) {
	originalContentURL := "https://dl.dropboxusercontent.com/u/358152/linebot/resource/gohper.jpg"
	previewImageURL := "https://dl.dropboxusercontent.com/u/358152/linebot/resource/gohper.jpg"
	message := linebot.NewImageMessage(originalContentURL, previewImageURL)
	result, err := botClient.ReplyMessage(replyToken, message)
	fmt.Println(result)
	fmt.Println(err)
}

// OnVideoMessage ...
func (be *BotEventHandler) OnVideoMessage(replyToken, id string) {
	originalContentURL := "https://dl.dropboxusercontent.com/u/358152/linebot/resource/video-original.mp4"
	previewImageURL := "https://dl.dropboxusercontent.com/u/358152/linebot/resource/video-preview.png"
	message := linebot.NewVideoMessage(originalContentURL, previewImageURL)
	result, err := botClient.ReplyMessage(replyToken, message)
	fmt.Println(result)
	fmt.Println(err)
}

// OnAudioMessage ...
func (be *BotEventHandler) OnAudioMessage(replyToken, id string) {
	originalContentURL := "https://dl.dropboxusercontent.com/u/358152/linebot/resource/ok.m4a"
	duration := 1000
	message := linebot.NewAudioMessage(originalContentURL, duration)
	result, err := botClient.ReplyMessage(replyToken, message)
	fmt.Println(result)
	fmt.Println(err)
}

// OnLocationMessage ...
func (be *BotEventHandler) OnLocationMessage(replyToken string, latitude, longitude float64) {
	title := "Disney Resort"
	address := "〒279-0031 千葉県浦安市舞浜１−１"
	lat := 35.632211
	lon := 139.881234
	message := linebot.NewLocationMessage(title, address, lat, lon)
	result, err := botClient.ReplyMessage(replyToken, message)
	fmt.Println(result)
	fmt.Println(err)
}

// OnStickerMessage ...
func (be *BotEventHandler) OnStickerMessage(replyToken, stickerID string) {
	message := linebot.NewStickerMessage("1", "1")
	result, err := botClient.ReplyMessage(replyToken, message)
	fmt.Println(result)
	fmt.Println(err)
}
