package dos

type GeneralRequest struct {
	SessionKey string `json:"sessionKey"`
}

type GeneralMessage struct {
	GeneralRequest
	Target       int64         `json:"target"`
	Quote        int64         `json:"quote,omitempty"`
	MessageChain []interface{} `json:"messageChain"`
}

type FriendMessageRequest struct {
	GeneralMessage
	QQ int64 `json:"qq"`
}

type GroupMessageRequest struct {
	GeneralMessage
	Group int64 `json:"group"`
}

func (g *GeneralMessage) AddChain(chain interface{}) *GeneralMessage {
	g.MessageChain = append(g.MessageChain, chain)
	return g
}

type TempMessageRequest struct {
	GeneralRequest
	QQ           int64         `json:"qq"`
	Group        int64         `json:"group"`
	Quote        int64         `json:"quote"`
	MessageChain []interface{} `json:"messageChain"`
}

type NudgeMessageRequest struct {
	GeneralRequest
	Target  int64  `json:"target"`
	Subject int64  `json:"subject"`
	Kind    string `json:"kind"`
}

// 撤回消息/删除好友
type GeneralMessageRequest struct {
	GeneralRequest
	Target int64 `json:"target"`
}

// 通用文件请求/删除文件
type GeneralFileRequest struct {
	GeneralRequest
	ID     string `json:"id"`
	Path   string `json:"path"`
	Target int64  `json:"target"`
	Group  int64  `json:"group"`
	Qq     int64  `json:"qq"`
}

type FileRequest struct {
	GeneralFileRequest
	WithDownloadInfo bool `json:"withDownloadInfo"`
}

type FileListRequest struct {
	GeneralFileRequest
	Offset int64 `json:"offset"`
	Size   int64 `json:"size"`
}

type CreateFolderRequest struct {
	GeneralFileRequest
	DirectoryName string `json:"directoryName"`
}

type MoveFileRequest struct {
	GeneralFileRequest
	MoveTo     *string `json:"moveTo"`
	MoveToPath *string `json:"moveToPath"`
}

type RenameFileRequest struct {
	GeneralFileRequest
	RenameTo *string `json:"rennameTo"`
}

// 通用群相关请求 / 退出群聊 / 全体禁言/解禁 / 设置群精华
type GeneralGroupRequest struct {
	GeneralRequest
	Target int64 `json:"target"`
}

// 禁言
type MuteMemberRequest struct {
	GeneralGroupRequest
	MemberID int64 `json:"memberId"`
	Time     int   `json:"time"`
}

// 解除某人禁言
type UnmuteMemberRequest struct {
	GeneralGroupRequest
	MemberID int64 `json:"memberId"`
}

// 移除某人
type KickMemberRequest struct {
	GeneralGroupRequest
	MemberID int64  `json:"memberID"`
	Msg      string `json:"msg"`
}

type SetMemberInfoEntity struct {
	Name         string `json:"name"`
	SpecialTitle string `json:"specialTitle"`
}

type SetMemberInfo struct {
	GeneralGroupRequest
	MemberID int64               `json:"memberId"`
	Info     SetMemberInfoEntity `json:"info"`
}

type EventRequest struct {
	GeneralRequest
	EventID int64  `json:"eventId"`
	FromID  int64  `json:"fromId"`
	GroupID int64  `json:"groupId"`
	Operate int64  `json:"operate"`
	Message string `json:"message"`
}

type SetGroupInfoRequest struct {
	GeneralGroupRequest
	Config GroupSetting `json:"config"`
}

type SetGroupAdminRequest struct {
	GeneralGroupRequest
	MemberID int64 `json:"memberId"`
	Assign   bool  `json:"assign"`
}
