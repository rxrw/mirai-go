package adapters

import (
	"errors"

	"github.com/rxrw/mirai-go/dos"

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

func (h GeneralAdapter) FriendMessage(target int64, quote int64, messageChains []interface{}) (int64, error) {
	uri := "sendFriendMessage"
	message := dos.FriendMessageRequest{
		GeneralMessage: dos.GeneralMessage{
			Target:       target,
			Quote:        quote,
			MessageChain: messageChains,
		},
		QQ: target,
	}
	message.SessionKey = h.sessionKey

	result, err := h.Send("POST", uri, message)

	messageBody, ok := result.(map[string]interface{})["messageId"]

	if !ok {
		return 0, err
	}

	return int64(messageBody.(float64)), err
}

func (h GeneralAdapter) GroupMessage(target int64, quote int64, messageChains []interface{}) (int64, error) {
	uri := "sendGroupMessage"
	message := dos.GroupMessageRequest{
		GeneralMessage: dos.GeneralMessage{
			Target:       target,
			Quote:        quote,
			MessageChain: messageChains,
		},
		Group: target,
	}
	message.SessionKey = h.sessionKey

	result, err := h.Send("POST", uri, message)

	messageBody, ok := result.(map[string]interface{})["messageId"]

	if !ok {
		return 0, err
	}

	return int64(messageBody.(float64)), err
}

func (h GeneralAdapter) TempMessage(group int64, qq int64, quote int64, messageChain []interface{}) (int64, error) {
	uri := "sendTempMessage"
	message := dos.TempMessageRequest{
		QQ:           qq,
		Group:        group,
		MessageChain: messageChain,
		Quote:        quote,
	}

	message.SessionKey = h.sessionKey

	result, err := h.Send("POST", uri, message)

	messageBody, ok := result.(map[string]interface{})["messageId"]

	if !ok {
		return 0, err
	}

	return int64(messageBody.(float64)), err
}

func (h GeneralAdapter) ReplyMessage(origin *dos.Message, addQuote bool, message []interface{}) (int64, error) {
	var quote int64
	if origin == nil {
		return 0, errors.New("origin is nil")
	}
	if !addQuote {
		quote = 0
	} else {
		quote = origin.GetMessageId()
	}
	if origin.IsType(dos.FRIEND) {
		return h.FriendMessage(int64(origin.Sender["id"].(float64)), quote, message)
	} else if origin.IsType(dos.GROUP) {
		return h.GroupMessage(int64(origin.Sender["group"].(map[string]interface{})["id"].(float64)), quote, message)
	} else if origin.IsType(dos.TEMP) {
		return h.TempMessage(int64(origin.Sender["group"].(map[string]interface{})["id"].(float64)), int64(origin.Sender["id"].(float64)), quote, message)
	}
	return 0, errors.New("no suitable message type")
}

func (h GeneralAdapter) Nudge(target int64, subject int64, kind string) error {
	uri := "sendNudge"
	message := &dos.NudgeMessageRequest{
		Target:  target,
		Subject: subject,
		Kind:    kind,
	}
	message.SessionKey = h.sessionKey
	_, err := h.Send("POST", uri, message)

	return err
}

func (h GeneralAdapter) Recall(target int64) error {
	uri := "recall"
	message := &dos.GeneralMessageRequest{
		Target: target,
	}
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

func (h GeneralAdapter) DeleteFriend(target int64) error {
	uri := "deleteFriend"
	req := dos.GeneralGroupRequest{
		Target: target,
	}
	req.SessionKey = h.sessionKey
	_, err := h.Send("POST", uri, req)

	return err
}

func (h GeneralAdapter) MuteGroupMember(target int64, memberId int64, duration int) error {
	uri := "mute"
	req := dos.MuteMemberRequest{
		GeneralGroupRequest: dos.GeneralGroupRequest{
			Target: target,
		},
		MemberID: memberId,
		Time:     duration,
	}
	req.SessionKey = h.sessionKey
	_, err := h.Send("POST", uri, req)

	return err
}

func (h GeneralAdapter) UnmuteGroupMember(target int64, memberId int64) error {
	uri := "unmute"
	req := dos.UnmuteMemberRequest{
		GeneralGroupRequest: dos.GeneralGroupRequest{
			Target: target,
		},
		MemberID: memberId,
	}
	req.SessionKey = h.sessionKey
	_, err := h.Send("POST", uri, req)

	return err
}

func (h GeneralAdapter) KickGroupMember(target int64, memberId int64) error {
	uri := "kick"
	req := dos.KickMemberRequest{
		GeneralGroupRequest: dos.GeneralGroupRequest{
			Target: target,
		},
		MemberID: memberId,
	}
	req.SessionKey = h.sessionKey
	_, err := h.Send("POST", uri, req)

	return err
}

func (h GeneralAdapter) QuitGroup(target int64) error {
	uri := "quit"
	req := dos.GeneralGroupRequest{
		Target: target,
	}
	req.SessionKey = h.sessionKey
	_, err := h.Send("POST", uri, req)

	return err
}

func (h GeneralAdapter) MuteGroup(target int64) error {
	uri := "muteAll"
	req := dos.GeneralGroupRequest{
		Target: target,
	}
	req.SessionKey = h.sessionKey
	_, err := h.Send("POST", uri, req)

	return err
}

func (h GeneralAdapter) UnmuteGroup(target int64) error {
	uri := "unmuteAll"
	req := dos.GeneralGroupRequest{
		Target: target,
	}
	req.SessionKey = h.sessionKey
	_, err := h.Send("POST", uri, req)

	return err
}

// EssenceGroup 设置群精华消息
func (h GeneralAdapter) EssenceGroup(target int64) error {
	uri := "setEssence"
	req := dos.GeneralGroupRequest{
		Target: target,
	}
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

func (h GeneralAdapter) GroupMemberSetting(target int64, memberID int64) (dos.UserInfo, error) {
	uri := "memberInfo"
	params := make(map[string]int64)
	params["target"] = target
	params["memberId"] = memberID
	var res dos.UserInfo
	result, err := h.Send("GET", uri, params)

	mapstructure.Decode(result, &res)
	return res, err
}

func (h GeneralAdapter) SetGroupMemberSetting(target int64, memberID int64, nickname string, specialTitle string) error {
	uri := "groupConfig"
	req := dos.SetMemberInfo{
		GeneralGroupRequest: dos.GeneralGroupRequest{Target: target},
		MemberID:            memberID,
		Info: dos.SetMemberInfoEntity{
			Name:         nickname,
			SpecialTitle: specialTitle,
		},
	}
	req.SessionKey = h.sessionKey
	_, err := h.Send("UPDATE", uri, req)

	return err
}

func (h GeneralAdapter) MemberAdmin(target int64, memberID int64, assign bool) error {
	uri := "memberAdmin"
	req := dos.SetGroupAdminRequest{
		GeneralGroupRequest: dos.GeneralGroupRequest{
			Target: target,
		},
		MemberID: memberID,
		Assign:   assign,
	}
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
