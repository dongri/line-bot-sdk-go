package linebot

// MessageType ...
type MessageType string

// MessageType ...
const (
	MessageTypeText     MessageType = "text"
	MessageTypeImage    MessageType = "image"
	MessageTypeVideo    MessageType = "video"
	MessageTypeAudio    MessageType = "audio"
	MessageTypeLocation MessageType = "location"
	MessageTypeSticker  MessageType = "sticker"
	MessageTypeTemplate MessageType = "template"
	MessageTypeImagemap MessageType = "imagemap"
)

// Message ...
type Message interface {
	message()
}

// TextMessage ...
type TextMessage struct {
	Type MessageType `json:"type"`
	Text string      `json:"text"`
}

// ImageMessage ...
type ImageMessage struct {
	Type               MessageType `json:"type"`
	OriginalContentURL string      `json:"originalContentUrl"`
	PreviewImageURL    string      `json:"previewImageUrl"`
}

// VideoMessage ...
type VideoMessage struct {
	Type               MessageType `json:"type"`
	OriginalContentURL string      `json:"originalContentUrl"`
	PreviewImageURL    string      `json:"previewImageUrl"`
}

// AudioMessage ...
type AudioMessage struct {
	Type               MessageType `json:"type"`
	OriginalContentURL string      `json:"originalContentUrl"`
	Duration           int         `json:"duration"`
}

// LocationMessage ...
type LocationMessage struct {
	Type      MessageType `json:"type"`
	Title     string      `json:"title"`
	Address   string      `json:"address"`
	Latitude  float64     `json:"latitude"`
	Longitude float64     `json:"longitude"`
}

// StickerMessage ...
type StickerMessage struct {
	Type      MessageType `json:"type"`
	PackageID string      `json:"packageId"`
	StickerID string      `json:"stickerId"`
}

// Message interface
func (*TextMessage) message()     {}
func (*ImageMessage) message()    {}
func (*VideoMessage) message()    {}
func (*AudioMessage) message()    {}
func (*LocationMessage) message() {}
func (*StickerMessage) message()  {}
func (*TemplateMessage) message() {}
func (*ImagemapMessage) message() {}

// NewTextMessage ...
func NewTextMessage(text string) *TextMessage {
	m := new(TextMessage)
	m.Type = MessageTypeText
	m.Text = text
	return m
}

// NewImageMessage ...
func NewImageMessage(originalContentURL, previewImageURL string) *ImageMessage {
	m := new(ImageMessage)
	m.Type = MessageTypeImage
	m.OriginalContentURL = originalContentURL
	m.PreviewImageURL = previewImageURL
	return m
}

// NewVideoMessage ...
func NewVideoMessage(originalContentURL, previewImageURL string) *VideoMessage {
	m := new(VideoMessage)
	m.Type = MessageTypeVideo
	m.OriginalContentURL = originalContentURL
	m.PreviewImageURL = previewImageURL
	return m
}

// NewAudioMessage ...
func NewAudioMessage(originalContentURL string, duration int) *AudioMessage {
	m := new(AudioMessage)
	m.Type = MessageTypeAudio
	m.OriginalContentURL = originalContentURL
	m.Duration = duration
	return m
}

// NewLocationMessage ...
func NewLocationMessage(title, address string, latitude, longitude float64) *LocationMessage {
	m := new(LocationMessage)
	m.Type = MessageTypeLocation
	m.Title = title
	m.Address = address
	m.Latitude = latitude
	m.Longitude = longitude
	return m
}

// NewStickerMessage ...
func NewStickerMessage(packageID, stickerID string) *StickerMessage {
	m := new(StickerMessage)
	m.Type = MessageTypeSticker
	m.PackageID = packageID
	m.StickerID = stickerID
	return m
}
