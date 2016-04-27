package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/dongri/line-bot-sdk-go/linebot"
)

var botClient *linebot.Client

func main() {
	channelID := os.Getenv("LINE_CHANNEL_ID")
	channelSecret := os.Getenv("LINE_CHANNEL_SECRET")
	mid := os.Getenv("LINE_MID")
	proxyURL := getProxyURL() // can set nil if not need

	botClient = linebot.NewClient(channelID, channelSecret, mid, proxyURL)

	// EventHandler
	var myEvent linebot.EventHandler = NewEventHandler()
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
	log.Print("=== Received Text ===")
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
