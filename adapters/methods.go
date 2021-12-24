package adapters

import (
	"reprover/mirai-go/dos"
)

type GeneralAdapter struct {
	Sender
	sessionKey string
}

func (h GeneralAdapter) About() (result *dos.PluginInfo, err error) {
	uri := "about"
	err = h.Send("GET", uri, nil, &result)
	return
}

func (h GeneralAdapter) MessageFromID(messageID string) (result *dos.Message, err error) {
	uri := "messageFromId"
	params := make(map[string]string)
	params["messageId"] = messageID
	err = h.Send("GET", uri, params, &result)
	return
}

func (h GeneralAdapter) FriendList() (result []*dos.User, err error) {
	uri := "friendList"
	err = h.Send("GET", uri, nil, &result)
	return
}

func (h GeneralAdapter) GroupList() (result []*dos.Group, err error) {
	uri := "groupList"
	err = h.Send("GET", uri, nil, &result)
	return
}

func (h GeneralAdapter) GroupMemberList(target int64) (result []*dos.GroupMember, err error) {
	uri := "memberList"
	params := make(map[string]int64)
	params["target"] = target
	err = h.Send("GET", uri, params, &result)
	return
}

func (h GeneralAdapter) BotProfile() (result *dos.UserInfo, err error) {
	uri := "botProfile"
	err = h.Send("GET", uri, nil, &result)
	return
}

func (h GeneralAdapter) FriendInfo(target int64) (result *dos.UserInfo, err error) {
	uri := "friendProfile"
	params := make(map[string]int64)
	params["target"] = target
	err = h.Send("GET", uri, params, &result)
	return
}

func (h GeneralAdapter) GroupMemberInfo(target int64, memberID int64) (result *dos.UserInfo, err error) {
	uri := "memberProfile"
	params := make(map[string]int64)
	params["target"] = target
	params["memberId"] = memberID
	err = h.Send("GET", uri, params, &result)
	return
}

func (h GeneralAdapter) FriendMessage(message dos.FriendMessageRequest) (result int64, err error) {
	uri := "sendFriendMessage"
	message.SessionKey = h.sessionKey
	err = h.Send("POST", uri, message, &result)

	return
}

func (h GeneralAdapter) GroupMessage(message dos.GroupMessageRequest) (result int64, err error) {
	uri := "sendGroupMessage"
	message.SessionKey = h.sessionKey
	err = h.Send("POST", uri, message, &result)

	return
}

func (h GeneralAdapter) TempMessage(message dos.TempMessageRequest) (result int64, err error) {
	uri := "sendTempMessage"
	message.SessionKey = h.sessionKey
	err = h.Send("POST", uri, message, &result)

	return
}

func (h GeneralAdapter) Nudge(message dos.NudgeMessageRequest) error {
	uri := "sendNudge"
	message.SessionKey = h.sessionKey
	err := h.Send("POST", uri, message, nil)

	return err
}

func (h GeneralAdapter) Recall(message dos.GeneralMessageRequest) error {
	uri := "recall"
	message.SessionKey = h.sessionKey
	err := h.Send("POST", uri, message, nil)

	return err
}

// URL 参数
func (h GeneralAdapter) FileList(req dos.FileListRequest) (result []*dos.FileInfo, err error) {
	uri := "file/list"
	req.SessionKey = h.sessionKey
	err = h.Send("GET", uri, req, &result)
	return
}

// URL 参数
func (h GeneralAdapter) File(req dos.FileRequest) (result *dos.FileInfo, err error) {
	uri := "file/info"
	req.SessionKey = h.sessionKey
	err = h.Send("GET", uri, req, &result)
	return
}

func (h GeneralAdapter) Mkdir(req dos.CreateFolderRequest) (result *dos.FileInfo, err error) {
	uri := "file/mkdir"
	req.SessionKey = h.sessionKey
	err = h.Send("POST", uri, req, &result)
	return
}

func (h GeneralAdapter) DeleteFile(req dos.GeneralFileRequest) error {
	uri := "file/delete"
	req.SessionKey = h.sessionKey
	err := h.Send("POST", uri, req, nil)

	return err
}

func (h GeneralAdapter) MoveFile(req dos.MoveFileRequest) error {
	uri := "file/move"
	req.SessionKey = h.sessionKey
	err := h.Send("POST", uri, req, nil)

	return err
}

func (h GeneralAdapter) RenameFile(req dos.RenameFileRequest) error {
	uri := "file/rename"
	req.SessionKey = h.sessionKey
	err := h.Send("POST", uri, req, nil)

	return err
}

func (h GeneralAdapter) DeleteFriend(req dos.GeneralGroupRequest) error {
	uri := "deleteFriend"
	req.SessionKey = h.sessionKey
	err := h.Send("POST", uri, req, nil)

	return err
}

func (h GeneralAdapter) MuteGroupMember(req dos.MuteMemberRequest) error {
	uri := "mute"
	req.SessionKey = h.sessionKey
	err := h.Send("POST", uri, req, nil)

	return err
}

func (h GeneralAdapter) UnmuteGroupMember(req dos.UnmuteMemberRequest) error {
	uri := "unmute"
	req.SessionKey = h.sessionKey
	err := h.Send("POST", uri, req, nil)

	return err
}

func (h GeneralAdapter) KickGroupMember(req dos.KickMemberRequest) error {
	uri := "kick"
	req.SessionKey = h.sessionKey
	err := h.Send("POST", uri, req, nil)

	return err
}

func (h GeneralAdapter) QuitGroup(req dos.GeneralGroupRequest) error {
	uri := "quit"
	req.SessionKey = h.sessionKey
	err := h.Send("POST", uri, req, nil)

	return err
}

func (h GeneralAdapter) MuteGroup(req dos.GeneralGroupRequest) error {
	uri := "muteAll"
	req.SessionKey = h.sessionKey
	err := h.Send("POST", uri, req, nil)

	return err
}

func (h GeneralAdapter) UnmuteGroup(req dos.GeneralGroupRequest) error {
	uri := "unmuteAll"
	req.SessionKey = h.sessionKey
	err := h.Send("POST", uri, req, nil)

	return err
}

func (h GeneralAdapter) EssenceGroup(req dos.GeneralGroupRequest) error {
	uri := "setEssence"
	req.SessionKey = h.sessionKey
	err := h.Send("POST", uri, req, nil)

	return err
}

func (h GeneralAdapter) GroupSetting(target int64) (result *dos.GroupSetting, err error) {
	uri := "groupConfig"
	params := make(map[string]int64)
	params["target"] = target
	err = h.Send("GET", uri, params, &result)
	return
}

func (h GeneralAdapter) SetGroupSetting(req dos.SetGroupInfoRequest) error {
	uri := "groupConfig"
	req.SessionKey = h.sessionKey
	err := h.Send("UPDATE", uri, req, nil)

	return err
}

func (h GeneralAdapter) GroupMemberSetting(target int64, memberID int64) (result *dos.MemberInfo, err error) {
	uri := "memberInfo"
	params := make(map[string]int64)
	params["target"] = target
	params["memberId"] = memberID
	err = h.Send("GET", uri, params, &result)
	return
}

func (h GeneralAdapter) SetGroupMemberSetting(req *dos.SetMemberInfo) error {
	uri := "groupConfig"
	req.SessionKey = h.sessionKey
	err := h.Send("UPDATE", uri, req, nil)

	return err
}

func (h GeneralAdapter) MemberAdmin(req dos.SetGroupAdminRequest) error {
	uri := "memberAdmin"
	req.SessionKey = h.sessionKey
	err := h.Send("POST", uri, req, nil)

	return err
}

func (h GeneralAdapter) DealNewFriendEvent(req *dos.EventRequest) error {
	uri := "resp/newFriendRequestEvent"
	req.SessionKey = h.sessionKey
	err := h.Send("POST", uri, req, nil)

	return err
}

func (h GeneralAdapter) DealNewGroupMemberEvent(req *dos.EventRequest) error {
	uri := "resp/memberJoinRequestEvent"
	req.SessionKey = h.sessionKey
	err := h.Send("POST", uri, req, nil)
	return err
}

func (h GeneralAdapter) DealInvitedGroupEvent(req *dos.EventRequest) error {
	uri := "resp_botInvitedJoinGroupRequestEvent"
	req.SessionKey = h.sessionKey
	err := h.Send("POST", uri, req, nil)

	return err
}
