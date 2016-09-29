package linebot

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

// Client ...
type Client struct {
	endPoint           string
	channelAccessToken string
	channelSecret      string
}

var eventHandler EventHandler

// NewClient ...
func NewClient(channelSecret, channelAccessToken string) *Client {
	return &Client{
		channelSecret:      channelSecret,
		channelAccessToken: channelAccessToken,
		endPoint:           EndPoint,
	}
}

// SetEventHandler ...
func (c *Client) SetEventHandler(evnt EventHandler) {
	eventHandler = evnt
}

func (c *Client) setHeader(req *http.Request) *http.Request {
	req.Header.Add("Content-Type", "application/json; charset=UTF-8")
	req.Header.Add("X-LINE-ChannelToken", c.channelAccessToken)
	req.Header.Add("Authorization", "Bearer "+c.channelAccessToken)
	return req
}

func (c *Client) do(req *http.Request) ([]byte, error) {
	client := &http.Client{
		Timeout: time.Duration(30 * time.Second),
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	return body, err
}

// SendText ...
func (c *Client) SendText(to []string, text string) (*SentResult, error) {
	content := new(Content)
	content.ContentType = ContentTypeText
	content.ToType = ToTypeUser
	content.Text = text
	return c.SendSingleMessage(to, *content)
}

// SendImage ...
func (c *Client) SendImage(to []string, originalContentURL, previewImageURL string) (*SentResult, error) {
	content := new(Content)
	content.ContentType = ContentTypeImage
	content.ToType = ToTypeUser
	content.OriginalContentURL = originalContentURL
	content.PreviewImageURL = previewImageURL
	return c.SendSingleMessage(to, *content)
}

// SendVideo ...
func (c *Client) SendVideo(to []string, originalContentURL, previewImageURL string) (*SentResult, error) {
	content := new(Content)
	content.ContentType = ContentTypeVideo
	content.ToType = ToTypeUser
	content.OriginalContentURL = originalContentURL
	content.PreviewImageURL = previewImageURL
	return c.SendSingleMessage(to, *content)
}

// SendAudio ...
func (c *Client) SendAudio(to []string, originalContentURL, audlen string) (*SentResult, error) {
	content := new(Content)
	content.ContentType = ContentTypeVideo
	content.ToType = ToTypeUser
	content.OriginalContentURL = originalContentURL
	metadata := new(ContentMetadata)
	metadata.AUDLEN = audlen
	content.ContentMetadata = *metadata
	return c.SendSingleMessage(to, *content)
}

// SendLocation ...
func (c *Client) SendLocation(to []string, address string, latitude, longitude float64, title string) (*SentResult, error) {
	content := new(Content)
	content.ContentType = ContentTypeLocation
	content.ToType = ToTypeUser
	content.Location.Address = address
	content.Location.Latitude = latitude
	content.Location.Longitude = longitude
	content.Location.Title = title
	return c.SendSingleMessage(to, *content)
}

// SendSticker ...
func (c *Client) SendSticker(to []string, stkID, stkpkgID, stkVer, stkTxt string) (*SentResult, error) {
	content := new(Content)
	content.ContentType = ContentTypeSticker
	content.ToType = ToTypeUser
	metadata := new(ContentMetadata)
	metadata.STKID = stkID
	metadata.STKPKGID = stkpkgID
	metadata.STKVER = stkVer
	metadata.STKTXT = stkTxt
	content.ContentMetadata = *metadata
	return c.SendSingleMessage(to, *content)
}

// SendContact ...
func (c *Client) SendContact(to []string, mid, displayName string) (*SentResult, error) {
	content := new(Content)
	content.ContentType = ContentTypeSticker
	content.ToType = ToTypeUser
	metadata := new(ContentMetadata)
	metadata.MID = mid
	metadata.DisplayName = displayName
	content.ContentMetadata = *metadata
	return c.SendSingleMessage(to, *content)
}

// SendRichMessage ...
func (c *Client) SendRichMessage(to []string, downloadURL string, altText string, markupJSON string) (*SentResult, error) {
	content := new(Content)
	content.ContentType = ContentTypeRich
	content.ToType = ToTypeUser
	metadata := new(ContentMetadata)
	metadata.DOWNLOADURL = downloadURL
	metadata.SPECREV = "1" //Fixed
	metadata.ALTTEXT = altText
	metadata.MARKUPJSON = markupJSON
	content.ContentMetadata = *metadata
	return c.SendSingleMessage(to, *content)
}

// SendMultipleMessage ...
func (c *Client) SendMultipleMessage(to []string, messageNotified int, content []Content) (*SentResult, error) {
	multipleMessage := new(MultipleMessage)
	multipleMessage.To = to
	multipleMessage.ToChannel = FixedToChannel
	multipleMessage.EventType = FixedEventTypeMultiple
	multipleContent := new(MultipleContent)
	multipleContent.MessageNotified = messageNotified
	multipleContent.Messages = content
	apiURL := c.endPoint + PathReplyMessage
	return c.sendMessage(apiURL, multipleMessage)
}

// SendSingleMessage ...
func (c *Client) SendSingleMessage(to []string, content Content) (*SentResult, error) {
	singleMessage := new(SingleMessage)
	singleMessage.To = to
	singleMessage.ToChannel = FixedToChannel
	singleMessage.EventType = FixedEventTypeSingle
	singleMessage.Content = content
	apiURL := c.endPoint + PathReplyMessage
	return c.sendMessage(apiURL, singleMessage)
}

// GetMessageContent ...
func (c *Client) GetMessageContent(messageID string) ([]byte, error) {
	apiURL := c.endPoint + fmt.Sprintf(PathGetMessageContent, messageID)
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, err
	}
	req = c.setHeader(req)
	body, err := c.do(req)
	return body, err
}

// GetUserProfiles ... mids is String (comma-separated)
func (c *Client) GetUserProfiles(mids string) (*UserProfiles, error) {
	apiURL := c.endPoint + fmt.Sprintf(PathGetProfile, mids)
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, err
	}
	values := url.Values{}
	values.Add("mids", mids)
	req.URL.RawQuery = values.Encode()
	req = c.setHeader(req)
	body, err := c.do(req)
	if err != nil {
		return nil, err
	}
	var profiles UserProfiles
	if err = json.Unmarshal(body, &profiles); err != nil {
		return nil, err
	}
	return &profiles, nil
}

// ReceiveMessageAndOperation ...
func (c *Client) ReceiveMessageAndOperation(body io.Reader) (*ReceivedMessage, error) {
	var receivedMessage ReceivedMessage
	b, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, &receivedMessage)
	if err != nil {
		return nil, err
	}
	return &receivedMessage, nil
}

func (c *Client) sendMessage(apiURL string, message interface{}) (*SentResult, error) {
	b, err := json.Marshal(message)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	req = c.setHeader(req)
	body, err := c.do(req)
	if err != nil {
		return nil, err
	}
	var result SentResult
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) validateSignature(signature string, body []byte) bool {
	decoded, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return false
	}
	hash := hmac.New(sha256.New, []byte(c.channelSecret))
	hash.Write(body)
	return hmac.Equal(decoded, hash.Sum(nil))
}
