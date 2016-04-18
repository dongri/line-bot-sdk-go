package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/dongri/LineBot"
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
	// SendImage
	originalContentURL := "http://weknowyourdreamz.com/image.php?pic=/images/robot/robot-03.jpg"
	previewImageURL := "http://weknowyourdreamz.com/image.php?pic=/images/robot/robot-03.jpg"
	botClient.SendImage([]string{from}, originalContentURL, previewImageURL)
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
	// Send Random Sticker
	r := rand.Intn(10) + 1
	stkID := strconv.Itoa(r)
	stkpkgID := "1"
	stkVer := "100"
	botClient.SendSticker([]string{from}, stkID, stkpkgID, stkVer, "hoge")
}

// OnContactMessage ...
func (be *BotEventHandler) OnContactMessage(from, MID, displayName string) {
	log.Print("=== Received Contact ===")
}
