package linebot

// LINE Server Endpoint
const (
	EndPoint = "https://api.line.me"
)

// API URLs
const (
	// URLSendMessage    = "/v1/events"
	// URLUserProfile    = "/v1/profiles"
	// URLMessageContent = "/v1/bot/message"

	PathPushMessage       = "/v2/bot/message/push"
	PathReplyMessage      = "/v2/bot/message/reply"
	PathGetMessageContent = "/v2/bot/message/%s/content"
	PathLeaveGroup        = "/v2/bot/group/%s/leave"
	PathLeaveRoom         = "/v2/bot/room/%s/leave"
	PathGetProfile        = "/v2/bot/profile/%s"
)

// ContentType
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

// OpType
const (
	OpTypeAdded   = 4
	OpTypeBlocked = 8
)

// Fixed
const (
	FixedToChannel         = 1383378250
	FixedEventTypeSingle   = "138311608800106203"
	FixedEventTypeMultiple = "140177271400161403"
)
