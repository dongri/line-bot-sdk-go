package LineBot

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// Middleware ...
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Print(err)
		}
		var received ReceivedMessage
		err = json.Unmarshal(b, &received)
		if err != nil {
			log.Print(err)
		}
		r.Body = ioutil.NopCloser(bytes.NewBuffer(b))
		for _, result := range received.Result {
			content := result.Content
			from := content.From
			if content.OpType == OpTypeAdded {
				MIDs := content.Params
				eventHandler.OnAddedAsFriendOperation(MIDs)
			}
			if content.OpType == OpTypeBlocked {
				MIDs := content.Params
				eventHandler.OnBlockedAccountOperation(MIDs)
			}
			if content.ContentType == ContentTypeText {
				from := content.From
				text := content.Text
				eventHandler.OnTextMessage(from, text)
			}
			if content.ContentType == ContentTypeImage {
				eventHandler.OnImageMessage(from)
			}
			if content.ContentType == ContentTypeVideo {
				eventHandler.OnVideoMessage(from)
			}
			if content.ContentType == ContentTypeAudio {
				eventHandler.OnAudioMessage(from)
			}
			if content.ContentType == ContentTypeLocation {
				location := content.Location
				title := location.Title
				address := location.Address
				latitude := location.Latitude
				longitude := location.Longitude
				eventHandler.OnLocationMessage(from, title, address, latitude, longitude)
			}
			if content.ContentType == ContentTypeSticker {
				contentMetadata := result.Content.ContentMetadata
				stickerID := contentMetadata.STKID
				stickerPackageID := contentMetadata.STKID
				stickerVersion := contentMetadata.STKVER
				stickerText := contentMetadata.STKTXT
				eventHandler.OnStickerMessage(from, stickerID, stickerPackageID, stickerVersion, stickerText)
			}
			if content.ContentType == ContentTypeContact {
				contentMetadata := result.Content.ContentMetadata
				MID := contentMetadata.MID
				displayName := contentMetadata.DisplayName
				eventHandler.OnContactMessage(from, MID, displayName)
			}
		}
		next.ServeHTTP(w, r)
	})
}
