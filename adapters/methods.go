package adapters

import (
	"fmt"
	"reprover/mirai-go/dos"

	"github.com/goinggo/mapstructure"
)

type GeneralAdapter struct {
	Sender
	sessionKey string
}

func (h GeneralAdapter) About() (dos.PluginInfo, error) {
	uri := "about"
	var res dos.PluginInfo
	result, err := h.Send("GET", uri, nil)
	mapstructure.Decode(result, &res)
	return res, err
}

func (h GeneralAdapter) MessageFromID(messageID string) (dos.Message, error) {
	uri := "messageFromId"
	params := make(map[string]string)
	params["messageId"] = messageID
	var res dos.Message
	result, err := h.Send("GET", uri, params)

	mapstructure.Decode(result, &res)
	return res, err
}

func (h GeneralAdapter) FriendList() ([]dos.User, error) {
	uri := "friendList"
	var res []dos.User
	result, err := h.Send("GET", uri, nil)
	if err != nil {
		return nil, err
	}
	result = result.(map[string]interface{})["data"]

	mapstructure.Decode(result, &res)
	return res, err
}

func (h GeneralAdapter) GroupList() ([]dos.Group, error) {
	uri := "groupList"
	var res []dos.Group
	result, err := h.Send("GET", uri, nil)
	if err != nil {
		return nil, err
	}
	result = result.(map[string]interface{})["data"]
	err = mapstructure.Decode(result, &res)
	return res, err
}

func (h GeneralAdapter) GroupMemberList(target int64) ([]dos.GroupMember, error) {
	uri := "memberList"
	params := make(map[string]int64)
	params["target"] = target
	var res []dos.GroupMember
	result, err := h.Send("GET", uri, params)
	if err != nil {
		return nil, err
	}
	result = result.(map[string]interface{})["data"]

	mapstructure.Decode(result, &res)
	return res, err
}

func (h GeneralAdapter) BotProfile() (dos.UserInfo, error) {
	uri := "botProfile"
	var res dos.UserInfo
	result, err := h.Send("GET", uri, nil)

	mapstructure.Decode(result, &res)
	return res, err
}

func (h GeneralAdapter) FriendInfo(target int64) (dos.UserInfo, error) {
	uri := "friendProfile"
	var res dos.UserInfo
	params := make(map[string]int64)
	params["target"] = target
	result, err := h.Send("GET", uri, params)
	mapstructure.Decode(result, &res)

	return res, err
}

func (h GeneralAdapter) GroupMemberInfo(target int64, memberID int64) (dos.UserInfo, error) {
	uri := "memberProfile"
	params := make(map[string]int64)
	params["target"] = target
	params["memberId"] = memberID
	var res dos.UserInfo
	result, err := h.Send("GET", uri, params)

	mapstructure.Decode(result, &res)
	return res, err
}

func (h GeneralAdapter) FriendMessage(message dos.FriendMessageRequest) (int, error) {
	uri := "sendFriendMessage"
	message.SessionKey = h.sessionKey

	fmt.Println(h.Sender)

	result, err := h.Send("POST", uri, message)

	messageBody, ok := result.(map[string]interface{})["messageId"]

	if !ok {
		return 0, err
	}

	return int(messageBody.(float64)), err
}

func (h GeneralAdapter) GroupMessage(message dos.GroupMessageRequest) (int, error) {
	uri := "sendGroupMessage"
	message.SessionKey = h.sessionKey

	result, err := h.Send("POST", uri, message)

	messageBody, ok := result.(map[string]interface{})["messageId"]

	if !ok {
		return 0, err
	}

	return int(messageBody.(float64)), err
}

func (h GeneralAdapter) TempMessage(message dos.TempMessageRequest) (int, error) {
	uri := "sendTempMessage"
	message.SessionKey = h.sessionKey

	result, err := h.Send("POST", uri, message)

	messageBody, ok := result.(map[string]interface{})["messageId"]

	if !ok {
		return 0, err
	}

	return int(messageBody.(float64)), err
}

func (h GeneralAdapter) Nudge(message dos.NudgeMessageRequest) error {
	uri := "sendNudge"
	message.SessionKey = h.sessionKey
	_, err := h.Send("POST", uri, message)

	return err
}

func (h GeneralAdapter) Recall(message dos.GeneralMessageRequest) error {
	uri := "recall"
	message.SessionKey = h.sessionKey
	_, err := h.Send("POST", uri, message)

	return err
}

// URL 参数
func (h GeneralAdapter) FileList(req dos.FileListRequest) ([]dos.FileInfo, error) {
	uri := "file/list"
	req.SessionKey = h.sessionKey
	var res []dos.FileInfo
	result, err := h.Send("GET", uri, req)
	if err != nil {
		return nil, err
	}
	result = result.(map[string]interface{})["data"]

	mapstructure.Decode(result, &res)
	return res, err
}

// URL 参数
func (h GeneralAdapter) File(req dos.FileRequest) (dos.FileInfo, error) {
	uri := "file/info"
	req.SessionKey = h.sessionKey
	var res dos.FileInfo
	result, err := h.Send("GET", uri, req)

	mapstructure.Decode(result, &res)
	return res, err
}

func (h GeneralAdapter) Mkdir(req dos.CreateFolderRequest) (dos.FileInfo, error) {
	uri := "file/mkdir"
	req.SessionKey = h.sessionKey
	var res dos.FileInfo
	result, err := h.Send("POST", uri, req)

	mapstructure.Decode(result, &res)
	return res, err
}

func (h GeneralAdapter) DeleteFile(req dos.GeneralFileRequest) error {
	uri := "file/delete"
	req.SessionKey = h.sessionKey
	_, err := h.Send("POST", uri, req)

	return err
}

func (h GeneralAdapter) MoveFile(req dos.MoveFileRequest) error {
	uri := "file/move"
	req.SessionKey = h.sessionKey
	_, err := h.Send("POST", uri, req)

	return err
}

func (h GeneralAdapter) RenameFile(req dos.RenameFileRequest) error {
	uri := "file/rename"
	req.SessionKey = h.sessionKey
	_, err := h.Send("POST", uri, req)

	return err
}

func (h GeneralAdapter) DeleteFriend(req dos.GeneralGroupRequest) error {
	uri := "deleteFriend"
	req.SessionKey = h.sessionKey
	_, err := h.Send("POST", uri, req)

	return err
}

func (h GeneralAdapter) MuteGroupMember(req dos.MuteMemberRequest) error {
	uri := "mute"
	req.SessionKey = h.sessionKey
	_, err := h.Send("POST", uri, req)

	return err
}

func (h GeneralAdapter) UnmuteGroupMember(req dos.UnmuteMemberRequest) error {
	uri := "unmute"
	req.SessionKey = h.sessionKey
	_, err := h.Send("POST", uri, req)

	return err
}

func (h GeneralAdapter) KickGroupMember(req dos.KickMemberRequest) error {
	uri := "kick"
	req.SessionKey = h.sessionKey
	_, err := h.Send("POST", uri, req)

	return err
}

func (h GeneralAdapter) QuitGroup(req dos.GeneralGroupRequest) error {
	uri := "quit"
	req.SessionKey = h.sessionKey
	_, err := h.Send("POST", uri, req)

	return err
}

func (h GeneralAdapter) MuteGroup(req dos.GeneralGroupRequest) error {
	uri := "muteAll"
	req.SessionKey = h.sessionKey
	_, err := h.Send("POST", uri, req)

	return err
}

func (h GeneralAdapter) UnmuteGroup(req dos.GeneralGroupRequest) error {
	uri := "unmuteAll"
	req.SessionKey = h.sessionKey
	_, err := h.Send("POST", uri, req)

	return err
}

func (h GeneralAdapter) EssenceGroup(req dos.GeneralGroupRequest) error {
	uri := "setEssence"
	req.SessionKey = h.sessionKey
	_, err := h.Send("POST", uri, req)

	return err
}

func (h GeneralAdapter) GroupSetting(target int64) (dos.GroupSetting, error) {
	uri := "groupConfig"
	params := make(map[string]int64)
	params["target"] = target
	var res dos.GroupSetting
	result, err := h.Send("GET", uri, params)

	mapstructure.Decode(result, &res)
	return res, err
}

func (h GeneralAdapter) SetGroupSetting(req dos.SetGroupInfoRequest) error {
	uri := "groupConfig"
	req.SessionKey = h.sessionKey
	_, err := h.Send("UPDATE", uri, req)

	return err
}

func (h GeneralAdapter) GroupMemberSetting(target int64, memberID int64) (dos.MemberInfo, error) {
	uri := "memberInfo"
	params := make(map[string]int64)
	params["target"] = target
	params["memberId"] = memberID
	var res dos.MemberInfo
	result, err := h.Send("GET", uri, params)

	mapstructure.Decode(result, &res)
	return res, err
}

func (h GeneralAdapter) SetGroupMemberSetting(req dos.SetMemberInfo) error {
	uri := "groupConfig"
	req.SessionKey = h.sessionKey
	_, err := h.Send("UPDATE", uri, req)

	return err
}

func (h GeneralAdapter) MemberAdmin(req dos.SetGroupAdminRequest) error {
	uri := "memberAdmin"
	req.SessionKey = h.sessionKey
	_, err := h.Send("POST", uri, req)

	return err
}

func (h GeneralAdapter) DealNewFriendEvent(req dos.EventRequest) error {
	uri := "resp/newFriendRequestEvent"
	req.SessionKey = h.sessionKey
	_, err := h.Send("POST", uri, req)

	return err
}

func (h GeneralAdapter) DealNewGroupMemberEvent(req dos.EventRequest) error {
	uri := "resp/memberJoinRequestEvent"
	req.SessionKey = h.sessionKey
	_, err := h.Send("POST", uri, req)
	return err
}

func (h GeneralAdapter) DealInvitedGroupEvent(req dos.EventRequest) error {
	uri := "resp_botInvitedJoinGroupRequestEvent"
	req.SessionKey = h.sessionKey
	_, err := h.Send("POST", uri, req)

	return err
}
