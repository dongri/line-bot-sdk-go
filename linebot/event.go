package linebot

// EventHandler ...
type EventHandler interface {
	OnAddedAsFriendOperation(MIDs []string)
	OnBlockedAccountOperation(MIDs []string)
	OnTextMessage(from, text string)
	OnImageMessage(from string)
	OnVideoMessage(from string)
	OnAudioMessage(from string)
	OnLocationMessage(from, title, address string, latitude, longitude float64)
	OnStickerMessage(from, stickerPackageID, stickerID, stickerVersion, stickerText string)
	OnContactMessage(from, MID, displayName string)
}
