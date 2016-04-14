package line

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// Client ...
type Client struct {
	endpoint      string
	channelID     string
	channelSecret string
	mID           string
	proxyURL      *url.URL
}

// Fixed ...
const (
	EndPoint               = "https://trialbot-api.line.me"
	FixedToChannel         = 1383378250
	FixedEventTypeSingle   = "138311608800106203"
	FixedEventTypeMultiple = "140177271400161403"
)

// ContentType ....
const (
	ContentTypeText int = iota + 1
	ContentTypeImage
	ContentTypeVideo
	ContentTypeAudio
	ContentTypeLocation = 7
	ContentTypeSticker  = 8
	ContentTypeContact  = 10
	ContentTypeRich     = 12
)

// ToType
const (
	ToTypeUser int = iota + 1
)

// API URLs
const (
	URLSendMessage = "/v1/events"
	URLUserProfile = "/v1/profiles"
)

// NewClient ...
func NewClient(endpoint, channelID, channelSecret, mid string, proxyURL *url.URL) *Client {
	return &Client{
		endpoint:      endpoint,
		channelID:     channelID,
		channelSecret: channelSecret,
		mID:           mid,
		proxyURL:      proxyURL,
	}
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
	apiURL := c.endpoint + URLSendMessage
	return c.sendMessage(apiURL, multipleMessage)
}

// SendSingleMessage ...
func (c *Client) SendSingleMessage(to []string, content Content) (*SentResult, error) {
	singleMessage := new(SingleMessage)
	singleMessage.To = to
	singleMessage.ToChannel = FixedToChannel
	singleMessage.EventType = FixedEventTypeSingle
	singleMessage.Content = content
	apiURL := c.endpoint + URLSendMessage
	return c.sendMessage(apiURL, singleMessage)
}

// GetUserProfiles ... mids is String (comma-separated)
func (c *Client) GetUserProfiles(mids string) (*UserProfiles, error) {
	apiURL := c.endpoint + URLUserProfile
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, err
	}
	values := url.Values{}
	values.Add("mids", mids)
	req.URL.RawQuery = values.Encode()
	req = c.setHeader(req)
	body, err := DoRequest(req, c.proxyURL)
	if err != nil {
		return nil, err
	}
	var profiles UserProfiles
	if err = json.Unmarshal(body, &profiles); err != nil {
		return nil, err
	}
	return &profiles, nil
}

// ReceiveMessage ...
func (c *Client) ReceiveMessage(body io.Reader) (*ReceivedMessage, error) {
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

func (c *Client) setHeader(req *http.Request) *http.Request {
	req.Header.Add("Content-Type", "application/json; charset=UTF-8")
	req.Header.Add("X-Line-ChannelID", c.channelID)
	req.Header.Add("X-Line-ChannelSecret", c.channelSecret)
	req.Header.Add("X-Line-Trusted-User-With-ACL", c.mID)
	return req
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
	body, err := DoRequest(req, c.proxyURL)
	if err != nil {
		return nil, err
	}
	var result SentResult
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	log.Print(result)
	return &result, nil
}
