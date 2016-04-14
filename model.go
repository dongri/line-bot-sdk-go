package line

// ReceivedMessage ...
type ReceivedMessage struct {
	Result []Result `json:"result"`
}

// Result ...
type Result struct {
	ID          string   `json:"id"`
	From        string   `json:"from"`
	FromChannel int      `json:"fromChannel"`
	To          []string `json:"to"`
	ToChannel   int      `json:"toChannel"`
	EventType   string   `json:"eventType"`
	Content     Content  `json:"content"`
}

// SingleMessage ..
type SingleMessage struct {
	To        []string `json:"to"`
	ToChannel int      `json:"toChannel"`
	EventType string   `json:"eventType"`
	Content   Content  `json:"content"`
}

// Content ...
type Content struct {
	ID                 string          `json:"id"`
	ContentType        int             `json:"contentType"`
	From               string          `json:"from"`
	CreatedTime        int             `json:"createdTime"`
	To                 []string        `json:"to"`
	ToType             int             `json:"toType"`
	Text               string          `json:"text"`
	Location           Location        `json:"location"`
	ContentMetadata    ContentMetadata `json:"contentMetadata"`
	OriginalContentURL string          `json:"originalContentUrl"`
	PreviewImageURL    string          `json:"previewImageUrl"`
}

// Location ...
type Location struct {
	Title     string  `json:"title"`
	Address   string  `json:"address"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// ContentMetadata ...
type ContentMetadata struct {
	STKID       string `json:"STKID"`
	STKPKGID    string `json:"STKPKGID"`
	STKVER      string `json:"STKVER"`
	STKTXT      string `json:"STKTXT"`
	MID         string `json:"mid"`
	DisplayName string `json:"displayName"`
	AUDLEN      string `json:"AUDLEN"`
	DOWNLOADURL string `json:"DOWNLOAD_URL"`
	SPECREV     string `json:"SPEC_REV"`
	ALTTEXT     string `json:"ALT_TEXT"`
	MARKUPJSON  string `json:"MARKUP_JSON"`
}

// UserProfiles ...
type UserProfiles struct {
	Contacts []Profile `json:"contacts"`
	Count    int       `json:"coudnt"`
	Display  int       `json:"display"`
	Start    int       `json:"start"`
	Total    int       `json:"total"`
}

// Profile ...
type Profile struct {
	DisplayName   string `json:"displayName"`
	MID           string `json:"mid"`
	PictureURL    string `json:"pictureUrl"`
	StatusMessage string `json:"statusMessage"`
}

// MultipleMessage ...
type MultipleMessage struct {
	To        []string        `json:"to"`
	ToChannel int             `json:"toChannel"`
	EventType string          `json:"eventType"`
	Content   MultipleContent `json:"content"`
}

// MultipleContent ...
type MultipleContent struct {
	MessageNotified int       `json:"messageNotified"`
	Messages        []Content `json:"messages"`
}

// SentResult ...
type SentResult struct {
	Failed    []interface{}
	MessageID string
	Timestamp int64
	Version   int
}

// MessageContent ...
type MessageContent struct {
	Content     Content  `json:"content"`
	From        string   `json:"from"`
	FromChannel int      `json:"fromChannel"`
	To          []string `json:"to"`
	ToChannel   int      `json:"toChannel"`
	EventType   string   `json:"eventType"`
	CreatedTime int      `json:"createdTime"`
	ID          string   `json:"id"`
}
