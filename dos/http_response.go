package dos

// GeneralResponse 发送拍一拍 / 撤回消息 / 上传文件 / 移动文件 / 重命名文件 / 删除好友 / ...
type GeneralResponse struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}

// PluginInfoResponse 获取插件信息
type PluginInfoResponse struct {
	GeneralResponse
	Data *PluginInfo `json:"data"`
}

// MessageByIDResponse 通过messageId获取消息
type MessageByIDResponse struct {
	GeneralResponse
	Data *Message `json:"data"`
}

// FriendListResponse 获取好友列表
type FriendListResponse struct {
	GeneralResponse
	Data []*User `json:"data"`
}

// 获取群列表
type GroupListResponse struct {
	GeneralResponse
	Data []*Group `json:"data"`
}

// 获取群成员列表
type GroupMemberResponse struct {
	GeneralResponse
	Data []*GroupMember `json:"data"`
}

// 获取 QQ 资料
type MemberInfo struct {
	GeneralResponse
	Data *UserInfo `json:"data"`
}

// 发送正常消息
type SendMessageResult struct {
	GeneralResponse
	MessageID int64 `json:"messageId"`
}

// 获取文件列表
type FileListResult struct {
	GeneralResponse
	Data []*FileInfo `json:"data"`
}

// 文件信息、创建文件夹
type FileInfoResult struct {
	GeneralResponse
	Data *FileInfo `json:"data"`
}

// 获取群设置
type GroupSettingResponse struct {
	GroupSetting
}
