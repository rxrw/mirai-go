package dos

const (
	NORMAL = iota
	ERROR_VERIFY_KEY
	BOT_NOT_EXISTS
	SESSION_INVALID
	SESSION_UNAUTHORIZED
	TARGET_NOT_FOUND
	FILE_NOT_FOUND
	PERMISSION_DENIED = 10
	BOT_HAS_BEEN_MUTE = 20
	MESSAGE_TOO_LONG  = 30
	BAD_REQUEST       = 40
)

const (
	GROUP        = "GroupMessage"
	FRIEND       = "FriendMessage"
	STRANGER     = "StrangerMessage"
	TEMP         = "TempMessage"
	OTHER_CLIENT = "OtherClientMessage"
)

const (
	Source         = "Source"
	Quote          = "Quote"
	At             = "At"
	AtAll          = "AtAll"
	Face           = "Face"
	Plain          = "Plain"
	Image          = "Image"
	Voice          = "Voice"
	FlashImage     = "FlashImage"
	Xml            = "Xml"
	Json           = "Json"
	App            = "App"
	Poke           = "Poke"
	Dice           = "Dice"
	MusicShare     = "MusicShare"
	ForwardMessage = "ForwardMessage"
	File           = "File"
	MiraiCode      = "MiraiCode"

	NUDGE = "Nudge"
)
