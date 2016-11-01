package linebot

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

// Webhook ...
type Webhook struct {
	Events []Event `json:"events"`
}

// Event ...
type Event struct {
	ReplyToken string      `json:"replyToken"`
	Type       EventType   `json:"type"`
	Timestamp  int64       `json:"timestamp"`
	Source     EventSource `json:"source"`
	Message    struct {
		ID        string      `json:"id"`
		Type      MessageType `json:"type"`
		Text      string      `json:"text"`
		Duration  int         `json:"duration"`
		Title     string      `json:"title"`
		Address   string      `json:"address"`
		Latitude  float64     `json:"latitude"`
		Longitude float64     `json:"longitude"`
		PackageID string      `json:"packageId"`
		StickerID string      `json:"stickerId"`
	} `json:"message"`
	Postback `json:"postback"`
	Beacon   `json:"beacon"`
}

// EventType ...
type EventType string

// EentTypes
const (
	EventTypeMessage  EventType = "message"
	EventTypeFollow   EventType = "follow"
	EventTypeUnfollow EventType = "unfollow"
	EventTypeJoin     EventType = "join"
	EventTypeLeave    EventType = "leave"
	EventTypePostback EventType = "postback"
	EventTypeBeacon   EventType = "beacon"
)

// EventSourceType ...
type EventSourceType string

// EventSourceType ....
const (
	EventSourceTypeUser  EventSourceType = "user"
	EventSourceTypeGroup EventSourceType = "group"
	EventSourceTypeRoom  EventSourceType = "room"
)

// EventSource ...
type EventSource struct {
	Type    EventSourceType `json:"type"`
	UserID  string          `json:"userId"`
	GroupID string          `json:"groupId"`
	RoomID  string          `json:"roomId"`
}

// Postback ...
type Postback struct {
	Data string `json:"data"`
}

// Beacon ...
type Beacon struct {
	Hwid string `json:"hwid"`
	Type string `json:"type"`
}

// Middleware ...
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		if !validateSignature(r.Header.Get("X-LINE-Signature"), body) {
			log.Fatal(errors.New("invalid signature"))
		}
		webhook := new(Webhook)
		if err = json.Unmarshal(body, webhook); err != nil {
			log.Fatal(err)
		}
		for _, event := range webhook.Events {
			eventHandler.OnEvent(event)
			switch event.Type {
			case EventTypeMessage:
				switch event.Message.Type {
				case MessageTypeText:
					eventHandler.OnTextMessage(event.Source, event.ReplyToken, event.Message.Text)
				case MessageTypeImage:
					eventHandler.OnImageMessage(event.Source, event.ReplyToken, event.Message.ID)
				case MessageTypeVideo:
					eventHandler.OnVideoMessage(event.Source, event.ReplyToken, event.Message.ID)
				case MessageTypeAudio:
					eventHandler.OnAudioMessage(event.Source, event.ReplyToken, event.Message.ID)
				case MessageTypeLocation:
					eventHandler.OnLocationMessage(event.Source, event.ReplyToken, event.Message.Title, event.Message.Address, event.Message.Latitude, event.Message.Longitude)
				case MessageTypeSticker:
					eventHandler.OnStickerMessage(event.Source, event.ReplyToken, event.Message.PackageID, event.Message.StickerID)
				}
			case EventTypeFollow:
				eventHandler.OnFollowEvent(event.Source, event.ReplyToken)
			case EventTypeUnfollow:
				eventHandler.OnUnFollowEvent(event.Source)
			case EventTypeJoin:
				eventHandler.OnJoinEvent(event.Source, event.ReplyToken)
			case EventTypeLeave:
				eventHandler.OnLeaveEvent(event.Source)
			case EventTypePostback:
				eventHandler.OnPostbackEvent(event.Source, event.ReplyToken, event.Postback.Data)
			case EventTypeBeacon:
				eventHandler.OnBeaconEvent(event.Source, event.ReplyToken, event.Beacon.Hwid, event.Beacon.Type)
			}
		}
		next.ServeHTTP(w, r)
	})
}

func validateSignature(signature string, body []byte) bool {
	decoded, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return false
	}
	hash := hmac.New(sha256.New, []byte(channelSecret))
	hash.Write(body)
	return hmac.Equal(decoded, hash.Sum(nil))
}

// EventHandler ...
type EventHandler interface {
	OnEvent(event Event)
	OnFollowEvent(source EventSource, replyToken string)
	OnUnFollowEvent(source EventSource)
	OnJoinEvent(source EventSource, replyToken string)
	OnLeaveEvent(source EventSource)
	OnPostbackEvent(source EventSource, replyToken, postbackData string)
	OnBeaconEvent(source EventSource, replyToken, beaconHwid string, beaconType string)
	OnTextMessage(source EventSource, replyToken, text string)
	OnImageMessage(source EventSource, replyToken, id string)
	OnVideoMessage(source EventSource, replyToken, id string)
	OnAudioMessage(source EventSource, replyToken, id string)
	OnLocationMessage(source EventSource, replyToken string, title, address string, latitude, longitude float64)
	OnStickerMessage(source EventSource, replyToken, packageID, stickerID string)
}
