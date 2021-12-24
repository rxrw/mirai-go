package adapters

import (
	"reprover/mirai-go/dos"
)

type Method interface {
	About() (result *dos.PluginInfo, err error)

	MessageFromID(message string) (result *dos.Message, err error)

	FriendList() (result []*dos.User, err error)

	GroupList() (result []*dos.Group, err error)

	GroupMemberList(target int64) (result []*dos.GroupMember, err error)

	BotProfile() (result *dos.UserInfo, err error)

	FriendInfo(target int64) (result *dos.UserInfo, err error)

	GroupMemberInfo(target int64, memberID int64) (result *dos.UserInfo, err error)

	FriendMessage(message dos.FriendMessageRequest) (result int64, err error)

	GroupMessage(message dos.GroupMessageRequest) (result int64, err error)

	TempMessage(message dos.TempMessageRequest) (result int64, err error)

	Nudge(message dos.NudgeMessageRequest) error

	Recall(message dos.GeneralMessageRequest) error

	FileList(req dos.FileListRequest) (result []*dos.FileInfo, err error)

	File(req dos.FileRequest) (result *dos.FileInfo, err error)

	Mkdir(req dos.CreateFolderRequest) (result *dos.FileInfo, err error)

	DeleteFile(req dos.GeneralFileRequest) error

	MoveFile(req dos.MoveFileRequest) error

	RenameFile(req dos.RenameFileRequest) error

	DeleteFriend(req dos.GeneralGroupRequest) error

	MuteGroupMember(req dos.MuteMemberRequest) error

	UnmuteGroupMember(req dos.UnmuteMemberRequest) error

	KickGroupMember(req dos.KickMemberRequest) error

	QuitGroup(req dos.GeneralGroupRequest) error

	MuteGroup(req dos.GeneralGroupRequest) error

	UnmuteGroup(req dos.GeneralGroupRequest) error

	EssenceGroup(req dos.GeneralGroupRequest) error

	GroupSetting(target int64) (*dos.GroupSetting, error)

	SetGroupSetting(req dos.SetGroupInfoRequest) error

	GroupMemberSetting(target int64, memberID int64) (result *dos.MemberInfo, err error)

	SetGroupMemberSetting(req *dos.SetMemberInfo) error

	MemberAdmin(req dos.SetGroupAdminRequest) error

	DealNewFriendEvent(req *dos.EventRequest) error

	DealNewGroupMemberEvent(req *dos.EventRequest) error

	DealInvitedGroupEvent(req *dos.EventRequest) error
}

type Sender interface {
	Send(method string, uri string, data interface{}, result interface{}) error
	Connect(ws chan Sender) error
}
