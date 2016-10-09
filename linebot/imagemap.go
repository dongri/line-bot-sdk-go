package linebot

// ImagemapActionType ...
type ImagemapActionType string

// ImagemapActionType ...
const (
	ImagemapActionTypeURI     ImagemapActionType = "uri"
	ImagemapActionTypeMessage ImagemapActionType = "message"
)

// ImagemapAction ...
type ImagemapAction interface {
	imagemapAction()
}

// ImagemapMessage ...
type ImagemapMessage struct {
	Type     MessageType      `json:"type"`
	BaseURL  string           `json:"baseUrl"`
	AltText  string           `json:"altText"`
	BaseSize ImagemapBaseSize `json:"baseSize"`
	Actions  []ImagemapAction `json:"actions"`
}

// ImagemapBaseSize ...
type ImagemapBaseSize struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

// ImagemapArea ...
type ImagemapArea struct {
	X      int `json:"x"`
	Y      int `json:"y"`
	Width  int `json:"width"`
	Height int `json:"height"`
}

// ImagemapURIAction ...
type ImagemapURIAction struct {
	Type    ImagemapActionType `json:"type"`
	LinkURI string             `json:"linkUri"`
	Area    ImagemapArea       `json:"area"`
}

// ImagemapMessageAction ...
type ImagemapMessageAction struct {
	Type ImagemapActionType `json:"type"`
	Text string             `json:"text"`
	Area ImagemapArea       `json:"area"`
}

func (*ImagemapURIAction) imagemapAction()     {}
func (*ImagemapMessageAction) imagemapAction() {}

// NewImagemapMessage ...
func NewImagemapMessage(baseURL, altText string, baseSizeWidth, baseSizeHeght int, actions ...ImagemapAction) *ImagemapMessage {
	m := new(ImagemapMessage)
	m.Type = MessageTypeImagemap
	m.BaseURL = baseURL
	m.AltText = altText
	size := new(ImagemapBaseSize)
	size.Width = baseSizeWidth
	size.Height = baseSizeHeght
	m.BaseSize = *size
	m.Actions = actions
	return m
}

// NewImagemapURIAction ...
func NewImagemapURIAction(linkURI string, area ImagemapArea) *ImagemapURIAction {
	a := new(ImagemapURIAction)
	a.Type = ImagemapActionTypeURI
	a.LinkURI = linkURI
	a.Area = area
	return a
}

// NewImagemapMessageAction ...
func NewImagemapMessageAction(text string, area ImagemapArea) *ImagemapMessageAction {
	a := new(ImagemapMessageAction)
	a.Type = ImagemapActionTypeMessage
	a.Text = text
	a.Area = area
	return a
}

// NewImagemapArea ...
func NewImagemapArea(x, y, width, height int) *ImagemapArea {
	a := new(ImagemapArea)
	a.X = x
	a.Y = y
	a.Width = width
	a.Height = height
	return a
}
