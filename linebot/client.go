package linebot

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

// API URLs
const (
	EndPoint          = "https://api.line.me"
	PushMessage       = "/v2/bot/message/push"
	ReplyMessage      = "/v2/bot/message/reply"
	GetMessageContent = "/v2/bot/message/%s/content"
	LeaveGroup        = "/v2/bot/group/%s/leave"
	LeaveRoom         = "/v2/bot/room/%s/leave"
	GetProfile        = "/v2/bot/profile/%s"
)

// APISendResult ...
type APISendResult struct {
	Message string `json:"message"`
}

// BasicResponse ...
type BasicResponse struct {
}

// MessageContentResponse ...
type MessageContentResponse struct {
	Content       []byte
	ContentLength int64
	ContentType   string
}

// UserProfileResponse ...
type UserProfileResponse struct {
	UserID        string `json:"userId"`
	DisplayName   string `json:"displayName"`
	PictureURL    string `json:"pictureUrl"`
	StatusMessage string `json:"statusMessage"`
}

// Client ...
type Client struct {
	endPoint           string
	channelAccessToken string
}

var eventHandler EventHandler
var channelSecret string

// NewClient ...
func NewClient(channelAccessToken string) *Client {
	return &Client{
		channelAccessToken: channelAccessToken,
		endPoint:           EndPoint,
	}
}

// SetEventHandler ...
func (c *Client) SetEventHandler(event EventHandler) {
	eventHandler = event
}

// SetChannelSecret ...
func (c *Client) SetChannelSecret(secret string) {
	channelSecret = secret
}

func (c *Client) setHeader(req *http.Request) *http.Request {
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("X-LINE-ChannelToken", c.channelAccessToken)
	req.Header.Set("Authorization", "Bearer "+c.channelAccessToken)
	req.Header.Set("User-Agent", "dongri/line-bot-sdk-go")
	return req
}

func (c *Client) do(req *http.Request) (*http.Response, []byte, error) {
	req = c.setHeader(req)
	client := &http.Client{
		Timeout: time.Duration(30 * time.Second),
	}
	res, err := client.Do(req)
	if err != nil {
		return res, nil, err
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode >= 400 {
		body, readErr := io.ReadAll(res.Body)
		if readErr != nil {
			return res, nil, readErr
		}
		var result APISendResult
		if unmarshalErr := json.Unmarshal(body, &result); unmarshalErr != nil {
			return res, nil, unmarshalErr
		}
		fmt.Println(result)
		return res, nil, errors.New("server error status code: " + strconv.Itoa(res.StatusCode))
	}
	body, err := io.ReadAll(res.Body)
	return res, body, err
}

// PushMessage ...
func (c *Client) PushMessage(to string, messages ...Message) (*APISendResult, error) {
	pushMessage := struct {
		To       string    `json:"to"`
		Messages []Message `json:"messages"`
	}{
		To:       to,
		Messages: messages,
	}
	b, err := json.Marshal(pushMessage)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", EndPoint+PushMessage, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	_, body, err := c.do(req)
	if err != nil {
		return nil, err
	}
	var result APISendResult
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ReplyMessage ...
func (c *Client) ReplyMessage(replyToken string, messages ...Message) (*APISendResult, error) {
	replyMessage := struct {
		ReplyToken string    `json:"replyToken"`
		Messages   []Message `json:"messages"`
	}{
		ReplyToken: replyToken,
		Messages:   messages,
	}
	b, err := json.Marshal(replyMessage)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", EndPoint+ReplyMessage, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	_, body, err := c.do(req)
	if err != nil {
		return nil, err
	}
	var result APISendResult
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetMessageContent ...
func (c *Client) GetMessageContent(messageID string) (*MessageContentResponse, error) {
	endpoint := fmt.Sprintf(EndPoint+GetMessageContent, messageID)
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	res, body, err := c.do(req)
	if err != nil {
		return nil, err
	}
	result := MessageContentResponse{
		Content:       body,
		ContentType:   res.Header.Get("Content-Type"),
		ContentLength: res.ContentLength,
	}
	return &result, nil
}

// GetProfile ...
func (c *Client) GetProfile(userID string) (*UserProfileResponse, error) {
	endpoint := fmt.Sprintf(EndPoint+GetProfile, userID)
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	_, body, err := c.do(req)
	if err != nil {
		return nil, err
	}
	result := UserProfileResponse{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// LeaveGroup ...
func (c *Client) LeaveGroup(groupID string) (*BasicResponse, error) {
	endpoint := fmt.Sprintf(EndPoint+LeaveGroup, groupID)
	req, err := http.NewRequest("POST", endpoint, nil)
	if err != nil {
		return nil, err
	}
	_, body, err := c.do(req)
	if err != nil {
		return nil, err
	}
	result := BasicResponse{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// LeaveRoom ...
func (c *Client) LeaveRoom(roomID string) (*BasicResponse, error) {
	endpoint := fmt.Sprintf(EndPoint+LeaveRoom, roomID)
	req, err := http.NewRequest("POST", endpoint, nil)
	if err != nil {
		return nil, err
	}
	_, body, err := c.do(req)
	if err != nil {
		return nil, err
	}
	result := BasicResponse{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
