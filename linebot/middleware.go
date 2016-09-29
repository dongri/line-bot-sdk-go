package linebot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// EventType type
type EventType string

// Types
const (
	EventTypeMessage  EventType = "message"
	EventTypeFollow   EventType = "follow"
	EventTypeUnfollow EventType = "unfollow"
	EventTypeJoin     EventType = "join"
	EventTypeLeave    EventType = "leave"
	EventTypePostback EventType = "postback"
	EventTypeBeacon   EventType = "beacon"
)

// EventSourceType type
type EventSourceType string

// EventSourceType constants
const (
	EventSourceTypeUser  EventSourceType = "user"
	EventSourceTypeGroup EventSourceType = "group"
	EventSourceTypeRoom  EventSourceType = "room"
)

// EventSource type
type EventSource struct {
	Type    EventSourceType `json:"type"`
	UserID  string          `json:"userId"`
	GroupID string          `json:"groupId"`
	RoomID  string          `json:"roomId"`
}

// Postback type
type Postback struct {
	Data string `json:"data"`
}

// BeaconEventType type
type BeaconEventType string

// BeaconEventType constants
const (
	BeaconEventTypeEnter BeaconEventType = "enter"
)

// Beacon type
type Beacon struct {
	Hwid string          `json:"hwid"`
	Type BeaconEventType `json:"type"`
}

// Event type
type Event struct {
	ReplyToken string
	Type       EventType
	Timestamp  time.Time
	Source     *EventSource
	Message    Message
	Postback   *Postback
	Beacon     *Beacon
}

// Message inteface
type Message interface {
	json.Marshaler
	message()
}

// TextMessage type
type TextMessage struct {
	ID   string
	Text string
}

// Middleware ...
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		defer r.Body.Close()
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		// if !validateSignature(r.Header.Get("X-LINE-Signature"), body) {
		// 	log.Fatal(errors.New("invalid signature"))
		// }
		result := &struct {
			Events []*Event `json:"events"`
		}{}
		if err = json.Unmarshal(body, result); err != nil {
			log.Fatal(err)
		}

		for _, event := range result.Events {
			switch event.Type {
			case EventTypeMessage:
				fmt.Println("message")
			case EventTypeFollow:
				fmt.Println("follow")
			case EventTypeUnfollow:
				fmt.Println("unfollow")
			case EventTypeJoin:
				fmt.Println("join")
			case EventTypeLeave:
				fmt.Println("leave")
			case EventTypePostback:
				fmt.Println("postback")
			case EventTypeBeacon:
				fmt.Println("beacon")
			}
		}

		// b, err := ioutil.ReadAll(r.Body)
		// if err != nil {
		// 	log.Print(err)
		// }
		// var received ReceivedMessage
		// err = json.Unmarshal(b, &received)
		// if err != nil {
		// 	log.Print(err)
		// }
		// r.Body = ioutil.NopCloser(bytes.NewBuffer(b))

		// for _, result := range received.Result {
		// 	content := result.Content
		// 	from := content.From
		// 	if content.OpType == OpTypeAdded {
		// 		MIDs := content.Params
		// 		eventHandler.OnAddedAsFriendOperation(MIDs)
		// 	}
		// 	if content.OpType == OpTypeBlocked {
		// 		MIDs := content.Params
		// 		eventHandler.OnBlockedAccountOperation(MIDs)
		// 	}
		// 	if content.ContentType == ContentTypeText {
		// 		from := content.From
		// 		text := content.Text
		// 		eventHandler.OnTextMessage(from, text)
		// 	}
		// 	if content.ContentType == ContentTypeImage {
		// 		eventHandler.OnImageMessage(from)
		// 	}
		// 	if content.ContentType == ContentTypeVideo {
		// 		eventHandler.OnVideoMessage(from)
		// 	}
		// 	if content.ContentType == ContentTypeAudio {
		// 		eventHandler.OnAudioMessage(from)
		// 	}
		// 	if content.ContentType == ContentTypeLocation {
		// 		location := content.Location
		// 		title := location.Title
		// 		address := location.Address
		// 		latitude := location.Latitude
		// 		longitude := location.Longitude
		// 		eventHandler.OnLocationMessage(from, title, address, latitude, longitude)
		// 	}
		// 	if content.ContentType == ContentTypeSticker {
		// 		contentMetadata := result.Content.ContentMetadata
		// 		stickerID := contentMetadata.STKID
		// 		stickerPackageID := contentMetadata.STKID
		// 		stickerVersion := contentMetadata.STKVER
		// 		stickerText := contentMetadata.STKTXT
		// 		eventHandler.OnStickerMessage(from, stickerID, stickerPackageID, stickerVersion, stickerText)
		// 	}
		// 	if content.ContentType == ContentTypeContact {
		// 		contentMetadata := result.Content.ContentMetadata
		// 		MID := contentMetadata.MID
		// 		displayName := contentMetadata.DisplayName
		// 		eventHandler.OnContactMessage(from, MID, displayName)
		// 	}
		// }
		next.ServeHTTP(w, r)
	})
}
