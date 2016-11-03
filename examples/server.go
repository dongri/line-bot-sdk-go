package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

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
	fmt.Fprintf(w, "LINE BOT SDK GO")
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
	if text == "Buttons" {
		templateLabel := "Go"
		templateText := "Hello, Golang!"
		thumbnailImageURL := "https://dl.dropboxusercontent.com/u/358152/linebot/resource/gopher.png"
		actionLabel := "Go to golang.org"
		actionURI := "https://golang.org"
		template := linebot.NewButtonsTemplate(
			thumbnailImageURL, templateLabel, templateText,
			linebot.NewTemplateURIAction(actionLabel, actionURI),
			linebot.NewTemplatePostbackAction("Go大好き", "Go大好き(Postback)", ""),
		)
		altText := "Go template"
		message := linebot.NewTemplateMessage(altText, template)
		result, err := botClient.ReplyMessage(replyToken, message)
		fmt.Println(result)
		fmt.Println(err)
	} else if text == "Confirm" {
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
	} else if text == "Audio" {
		originalContentURL := "https://dl.dropboxusercontent.com/u/358152/linebot/resource/ok.m4a"
		duration := 1000
		message := linebot.NewAudioMessage(originalContentURL, duration)
		result, err := botClient.ReplyMessage(replyToken, message)
		fmt.Println(result)
		fmt.Println(err)
	} else if text == "Carousel" {
		var columns []*linebot.CarouselColumn
		for i := 0; i < 5; i++ {
			originalContentURL := GetImageFromWeb()
			originalContentURL = strings.Replace(originalContentURL, "http://", "https://", -1)
			column := linebot.NewCarouselColumn(
				originalContentURL, "", "",
				linebot.NewTemplatePostbackAction("好き！", "好き！", "好き！"),
				linebot.NewTemplateMessageAction("普通", "普通"),
			)
			columns = append(columns, column)
		}
		template := linebot.NewCarouselTemplate(columns...)

		message := linebot.NewTemplateMessage("Sexy Girl", template)
		result, err := botClient.ReplyMessage(replyToken, message)
		fmt.Println(result)
		fmt.Println(err)
	} else {
		//message := linebot.NewTextMessage(text + "じゃねぇよ！")
		//result, err := botClient.ReplyMessage(replyToken, message)
		//fmt.Println(result)
		//fmt.Println(err)
	}
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

// OnEvent ...
func (be *BotEventHandler) OnEvent(event linebot.Event) {
}
