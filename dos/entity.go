package dos

import (
	"errors"
	"time"
)

// 插件信息
type PluginInfo struct {
	Version string `json:"version"`
}

// 群简单
type Group struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Permission string `json:"permission"`
}

type AbstractUser struct {
	ID int `json:"id"`
}

// 用户简单 Sender
type User struct {
	AbstractUser
	Nickname string `json:"nickname"`
	Remark   string `json:"remark"`
}

// 群成员信息 Sender
type GroupMember struct {
	AbstractUser
	MemberName         string `json:"memberName"`
	SpecialTitle       string `json:"specialTitle"`
	Permission         string `json:"permission"`
	JoinTimestamp      int64  `json:"joinTimestamp"`
	LastSpeakTimestamp int64  `json:"lastSpeakTimestamp"`
	MuteTimeRemaining  int64  `json:"muteTimeRemaining"`
	Group              Group  `json:"group"`
}

// 不同客户端 Sender
type Client struct {
	AbstractUser
	Platform string `json:"platform"`
}

// 用户信息
type UserInfo struct {
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Age      int64  `json:"age"`
	Level    int64  `json:"level"`
	Sign     string `json:"sign"`
	Sex      string `json:"sex"`
}

// 文件信息
type FileInfo struct {
	Name         string        `json:"name"`
	ID           string        `json:"id"`
	Path         string        `json:"path"`
	Parent       *FileInfo     `json:"parent"`
	Contact      interface{}   `json:"contact"`
	IsFile       bool          `json:"isFile"`
	IsDirectory  bool          `json:"isDirectory"`
	DownloadInfo *DownloadInfo `json:"downloadInfo"`
}

// 文件下载信息
type DownloadInfo struct {
	Sha1           string `json:"sha1"`
	Md5            string `json:"md5"`
	DownloadTimes  int64  `json:"downloadTimes"`
	UploaderID     int64  `json:"uploaderId"`
	UploadTime     int64  `json:"uploadTime"`
	LastModifyTime int64  `json:"lastModifyTime"`
	URL            string `json:"url"`
}

// 群设置
type GroupSetting struct {
	Name              string `json:"name"`
	Announcement      string `json:"announcement"`
	ConfessTalk       bool   `json:"confessTalk"`
	AllowMemberInvite bool   `json:"allowMemberInvite"`
	AutoApprove       bool   `json:"autoApprove"`
	AnonymousChat     bool   `json:"anonymousChat"`
}

type Message struct {
	Type         string                   `json:"type"`
	MessageChain []map[string]interface{} `json:"messageChain"`
	Sender       map[string]interface{}   `json:"sender"` // User / GroupMember / Platform
}

// GetSenderQQ 发送者的QQ
func (m Message) GetSenderQQ() int64 {
	sender := m.Sender
	return int64(sender["id"].(float64))
}

// GetAt 获取所有被艾特的QQ
func (m Message) GetAt() ([]int64, error) {
	at := make([]int64, 0)
	if m.Type != GROUP {
		return at, errors.New("invalid message type")
	}
	for _, c := range m.MessageChain {
		if c["type"].(string) == At {
			at = append(at, int64(c["target"].(float64)))
		}
	}
	return at, nil
}

func (m Message) IsType(kind string) bool {
	return m.Type == kind
}

// IsAt 是否艾特了指定成员
func (m Message) IsAt(qq int64) bool {
	if m.Type != GROUP {
		return false
	}
	for _, c := range m.MessageChain {
		if c["type"].(string) == At {
			if int64(c["target"].(float64)) == qq {
				return true
			}
		}
	}
	return false
}

// IsAtAll 是否是全体消息
func (m Message) IsAtAll() bool {
	if m.Type != GROUP {
		return false
	}
	for _, c := range m.MessageChain {
		if c["type"].(string) == AtAll {
			return true
		}
	}
	return false
}

// GetPlainMessage 获取所有文本
func (m Message) GetPlainMessage() string {
	content := ""
	for _, c := range m.MessageChain {
		if c["type"].(string) == Plain {
			content += c["message"].(string)
		}
	}
	return content
}

func (m Message) GetMessageSentAt() time.Time {
	for _, c := range m.MessageChain {
		if c["type"].(string) == Source {
			sendAt := int64(c["time"].(float64))
			return time.Unix(sendAt, 0).Local()
		}
	}
	return time.Now()
}

func (m Message) GetMessageId() int64 {
	for _, c := range m.MessageChain {
		if c["type"].(string) == Source {
			return int64(c["id"].(float64))
		}
	}
	return 0
}

func (m Message) GetQuoteId() int64 {
	for _, c := range m.MessageChain {
		if c["type"].(string) == Quote {
			return int64(c["id"].(float64))
		}
	}
	return 0
}
