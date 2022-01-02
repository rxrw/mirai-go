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

// 事件类型
const (
	BOT_ONLINE_EVENT          = "BotOnlineEvent"
	BOT_OFFLINE_EVENT_ACTIVE  = "BotOfflineEventActive"
	BOT_OFFLINE_EVENT_FORCE   = "BotOfflineEventForce"
	BOT_OFFLINE_EVENT_DROPPED = "BotOfflineEventDropped"
	BOT_RELOGIN_EVENT         = "BotReloginEvent"

	FRIEND_INPUT_STATUS_CHANGED_EVENT = "FriendInputStatusChangedEvent"
	FRIEND_NICK_CHANGED_EVENT         = "FriendNickChangedEvent"

	BOT_GROUP_PERMISSION_CHANGE_EVENT        = "BotGroupPermissionChangeEvent"
	BOT_MUTE_EVENT                           = "BotMuteEvent"
	BOT_UNMUTE_EVENT_BOT_JOIN_GROUP_EVENT    = "BotUnmuteEventBotJoinGroupEvent"
	BOT_LEAVE_EVENT_ACTIVE                   = "BotLeaveEventActive"
	BOT_LEAVE_EVENT_KICK                     = "BotLeaveEventKick"
	GROUP_RECALL_EVENT                       = "GroupRecallEvent"
	FRIEND_RECALL_EVENT                      = "FriendRecallEvent"
	NUDGE_EVENT                              = "NudgeEvent"
	GROUP_NAME_CHANGE_EVENT                  = "GroupNameChangeEvent"
	GROUP_ENTRANCE_ANNOUNCEMENT_CHANGE_EVENT = "GroupEntranceAnnouncementChangeEvent"
	GROUP_MUTE_ALL_EVENT                     = "GroupMuteAllEvent"
	GROUP_ALLOW_ANONYMOUS_CHAT_EVENT         = "GroupAllowAnonymousChatEvent"
	GROUP_ALLOW_CONFESS_TALK_EVENT           = "GroupAllowConfessTalkEvent"
	GROUP_ALLOW_MEMBER_INVITE_EVENT          = "GroupAllowMemberInviteEvent"
	MEMBER_JOIN_EVENT                        = "MemberJoinEvent"
	MEMBER_LEAVE_EVENT_KICK                  = "MemberLeaveEventKick"
	MEMBER_LEAVE_EVENT_QUIT                  = "MemberLeaveEventQuit"
	MEMBER_CARD_CHANGE_EVENT                 = "MemberCardChangeEvent"
	MEMBER_SPECIAL_TITLE_CHANGE_EVENT        = "MemberSpecialTitleChangeEvent"
	MEMBER_PERMISSION_CHANGE_EVENT           = "MemberPermissionChangeEvent"
	MEMBER_MUTE_EVENT                        = "MemberMuteEvent"
	MEMBER_UNMUTE_EVENT                      = "MemberUnmuteEvent"
	MEMBER_HONOR_CHANGE_EVENT                = "MemberHonorChangeEvent"

	NEW_FRIEND_REQUEST_EVENT             = "NewFriendRequestEvent"
	MEMBER_JOIN_REQUEST_EVENT            = "MemberJoinRequestEvent"
	BOT_INVITED_JOIN_GROUP_REQUEST_EVENT = "BotInvitedJoinGroupRequestEvent"
	OTHER_CLIENT_ONLINE_EVENT            = "OtherClientOnlineEvent"
	OTHER_CLIENT_OFFLINE_EVENT           = "OtherClientOfflineEvent"
	COMMAND_EXECUTED_EVENT               = "CommandExecutedEvent"
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
