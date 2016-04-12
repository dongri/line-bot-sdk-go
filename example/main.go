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
	m, err := botClient.ReceiveMessage(r.Body)
	if err != nil {
		log.Print(err)
	}
	for _, result := range m.Result {
		from := result.Content.From

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
			fmt.Fprintf(w, "Failed")
			return
		}

		// Send Sticker
		metadata := new(line.ContentMetadata)
		metadata.STKID = "2"
		metadata.STKPKGID = "1"
		metadata.STKVER = "100"
		sentResult, err = botClient.SendSticker([]string{from}, *metadata)
		if err != nil {
			log.Print(err)
		}
		if len(sentResult.Failed) == 0 {
			fmt.Fprintf(w, "Failed")
			return
		}

		// Send Image
		originalContentURL := "https://upload.wikimedia.org/wikipedia/commons/5/5e/Line_logo.png"
		previewImageURL := "http://i.imgur.com/Aaso4sY.png"
		sentResult, err = botClient.SendImage([]string{from}, originalContentURL, previewImageURL)
		if err != nil {
			log.Print(err)
		}
		if len(sentResult.Failed) == 0 {
			fmt.Fprintf(w, "Failed")
			return
		}

		// Send Video ....

	}
	fmt.Fprintf(w, "Success")
}

func getProxyURL() *url.URL {
	proxyURL, err := url.Parse(os.Getenv("PROXY_URL"))
	if err != nil {
		log.Fatal(err)
	}
	return proxyURL
}
