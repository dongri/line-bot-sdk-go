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
			log.Print("Added as friend")
		}

		if result.Content.OpType == line.OpTypeBlocked {
			log.Print("Blocked account")
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
