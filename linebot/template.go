package linebot

// TemplateType ...
type TemplateType string

// TemplateType ...
const (
	TemplateTypeButtons  TemplateType = "buttons"
	TemplateTypeConfirm  TemplateType = "confirm"
	TemplateTypeCarousel TemplateType = "carousel"
)

// TemplateActionType ...
type TemplateActionType string

// TemplateActionType ...
const (
	TemplateActionTypeURI      TemplateActionType = "uri"
	TemplateActionTypeMessage  TemplateActionType = "message"
	TemplateActionTypePostback TemplateActionType = "postback"
)

// Template ...
type Template interface {
	template()
}

// TemplateAction ...
type TemplateAction interface {
	templateAction()
}

// TemplateMessage ...
type TemplateMessage struct {
	Type     MessageType `json:"type"`
	AltText  string      `json:"altText"`
	Template Template    `json:"template"`
}

// ButtonsTemplate ...
type ButtonsTemplate struct {
	Type              TemplateType     `json:"type"`
	ThumbnailImageURL string           `json:"thumbnailImageUrl,omitempty"`
	Title             string           `json:"title,omitempty"`
	Text              string           `json:"text"`
	Actions           []TemplateAction `json:"actions"`
}

// ConfirmTemplate ...
type ConfirmTemplate struct {
	Type    TemplateType     `json:"type"`
	Text    string           `json:"text"`
	Actions []TemplateAction `json:"actions"`
}

// CarouselTemplate ...
type CarouselTemplate struct {
	Type    TemplateType      `json:"type"`
	Columns []*CarouselColumn `json:"columns"`
}

// TemplateURIAction ...
type TemplateURIAction struct {
	Type  TemplateActionType `json:"type"`
	Label string             `json:"label"`
	URI   string             `json:"uri"`
}

// TemplateMessageAction ...
type TemplateMessageAction struct {
	Type  TemplateActionType `json:"type"`
	Label string             `json:"label"`
	Text  string             `json:"text"`
}

// TemplatePostbackAction ...
type TemplatePostbackAction struct {
	Type  TemplateActionType `json:"type"`
	Label string             `json:"label"`
	Data  string             `json:"data"`
	Text  string             `json:"text,omitempty"`
}

// CarouselColumn ...
type CarouselColumn struct {
	ThumbnailImageURL string           `json:"thumbnailImageUrl,omitempty"`
	Title             string           `json:"title,omitempty"`
	Text              string           `json:"text"`
	Actions           []TemplateAction `json:"actions"`
}

// Template interface
func (*ConfirmTemplate) template()  {}
func (*ButtonsTemplate) template()  {}
func (*CarouselTemplate) template() {}

// TemplateAction interface
func (*TemplateURIAction) templateAction()      {}
func (*TemplateMessageAction) templateAction()  {}
func (*TemplatePostbackAction) templateAction() {}

// NewTemplateMessage ...
func NewTemplateMessage(altText string, template Template) *TemplateMessage {
	m := new(TemplateMessage)
	m.Type = MessageTypeTemplate
	m.AltText = altText
	m.Template = template
	return m
}

// NewConfirmTemplate ...
func NewConfirmTemplate(text string, actions ...TemplateAction) *ConfirmTemplate {
	t := new(ConfirmTemplate)
	t.Type = TemplateTypeConfirm
	t.Text = text
	t.Actions = actions
	return t
}

// NewButtonsTemplate ...
func NewButtonsTemplate(thumbnailImageURL, title, text string, actions ...TemplateAction) *ButtonsTemplate {
	t := new(ButtonsTemplate)
	t.Type = TemplateTypeButtons
	t.ThumbnailImageURL = thumbnailImageURL
	t.Title = title
	t.Text = text
	t.Actions = actions
	return t
}

// NewCarouselTemplate ...
func NewCarouselTemplate(columns ...*CarouselColumn) *CarouselTemplate {
	t := new(CarouselTemplate)
	t.Type = TemplateTypeCarousel
	t.Columns = columns
	return t
}

// NewCarouselColumn ...
func NewCarouselColumn(thumbnailImageURL, title, text string, actions ...TemplateAction) *CarouselColumn {
	c := new(CarouselColumn)
	c.ThumbnailImageURL = thumbnailImageURL
	c.Title = title
	c.Text = text
	c.Actions = actions
	return c
}

// NewTemplateURIAction ...
func NewTemplateURIAction(label, uri string) *TemplateURIAction {
	a := new(TemplateURIAction)
	a.Type = TemplateActionTypeURI
	a.Label = label
	a.URI = uri
	return a
}

// NewTemplateMessageAction ...
func NewTemplateMessageAction(label, text string) *TemplateMessageAction {
	a := new(TemplateMessageAction)
	a.Type = TemplateActionTypeMessage
	a.Label = label
	a.Text = text
	return a
}

// NewTemplatePostbackAction ...
func NewTemplatePostbackAction(label, data, text string) *TemplatePostbackAction {
	a := new(TemplatePostbackAction)
	a.Type = TemplateActionTypePostback
	a.Label = label
	a.Data = data
	a.Text = text
	return a
}
