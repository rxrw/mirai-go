package dos

type AbstractMessageChain struct {
	Type string `json:"type"`
}

type SourceMessageChain struct {
	AbstractMessageChain
	Time int64 `json:"time"`
	ID   int64 `json:"id"`
}

type QuoteMessageChain struct {
	AbstractMessageChain
	ID       int64         `json:"id"` // 被引用回复的原消息的messageId
	GroupID  int64         `json:"groupId"`
	SenderID int64         `json:"senderId"`
	TargetID int64         `json:"targetId"`
	Origin   []interface{} `json:"origin"`
}

type AtMessageChain struct {
	AbstractMessageChain
	Target  int64  `json:"target"`
	Display string `json:"display"`
}

type AtAllMessageChain struct {
	AbstractMessageChain
}

type FaceMessageChain struct {
	AbstractMessageChain
	FaceID int    `json:"faceId"`
	Name   string `json:"name"`
}

type PlainMessageChain struct {
	AbstractMessageChain
	Text string `json:"text"`
}

type ImageMessageChain struct {
	AbstractMessageChain
	ImageID string      `json:"imageId"`
	URL     string      `json:"url"`
	Path    interface{} `json:"path"`
	Base64  interface{} `json:"base64"`
}

type FlashImageMessageChain struct {
	ImageMessageChain
}

type VoiceMessageChain struct {
	AbstractMessageChain
	VoiceID string      `json:"voiceId"`
	URL     string      `json:"url"`
	Path    interface{} `json:"path"`
	Base64  interface{} `json:"base64"`
	Length  int         `json:"length"`
}

type XmlMessageChain struct {
	AbstractMessageChain
	Xml string `json:"xml"`
}

type JsonMessageChain struct {
	AbstractMessageChain
	JSON string `json:"json"`
}

type AppMessageChain struct {
	AbstractMessageChain
	Content string `json:"content"`
}

// PokeMessageChain 戳一戳
type PokeMessageChain struct {
	AbstractMessageChain
	Name string `json:"name"`
}

// DiceMessageChain 点数
type DiceMessageChain struct {
	AbstractMessageChain
	Value int `json:"value"`
}

type MusicShareMessageChain struct {
	AbstractMessageChain
	Kind       string `json:"kind"`    // 类型
	Title      string `json:"title"`   // 标题
	Summary    string `json:"summary"` // 概括
	JumpURL    string `json:"jumpUrl"`
	PictureURL string `json:"pictureUrl"`
	MusicURL   string `json:"musicUrl"` // 音源
	Brief      string `json:"brief"`
}

// ForwardMessageChain 转发的消息
type ForwardMessageChain struct {
	AbstractMessageChain
	NodeList []NodeList `json:"nodeList"`
}

type NodeList struct {
	SenderID     int64         `json:"senderId"`
	Time         int64         `json:"time"`
	SenderName   string        `json:"senderName"`
	MessageChain []interface{} `json:"messageChain"`
	MessageID    int64         `json:"messageId"`
}

type FileMessageChain struct {
	AbstractMessageChain
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Size int64  `json:"size"`
}

// MiraiCodeMessageChain hello[mirai:at:1234567]
type MiraiCodeMessageChain struct {
	AbstractMessageChain
	Code string `json:"code"`
}

func NewAbstractMessageChain(Type string) AbstractMessageChain {
	return AbstractMessageChain{
		Type: Type,
	}
}

func NewPlainMessageChain(text string) PlainMessageChain {
	return PlainMessageChain{
		AbstractMessageChain: NewAbstractMessageChain("Plain"),
		Text:                 text,
	}
}

func NewAtMessageChain(target int64) AtMessageChain {
	return AtMessageChain{
		AbstractMessageChain: NewAbstractMessageChain("At"),
		Target:               target,
	}
}

func NewAtAllMessageChain() AtAllMessageChain {
	return AtAllMessageChain{
		AbstractMessageChain: NewAbstractMessageChain("AtAll"),
	}
}

func NewFaceMessageChain(faceId int, name string) FaceMessageChain {
	return FaceMessageChain{
		AbstractMessageChain: NewAbstractMessageChain("Face"),
		FaceID:               faceId,
		Name:                 name,
	}
}

func NewImageMessageChain(imageId string, url string, path string, base64 string) ImageMessageChain {
	return ImageMessageChain{
		AbstractMessageChain: NewAbstractMessageChain("Image"),
		ImageID:              imageId,
		URL:                  url,
		Path:                 path,
		Base64:               base64,
	}
}

func NewVocieMessageChain(vocieId string, url string, path string, base64 string) VoiceMessageChain {
	return VoiceMessageChain{
		AbstractMessageChain: NewAbstractMessageChain("Voice"),
		VoiceID:              vocieId,
		URL:                  url,
		Path:                 path,
		Base64:               base64,
	}
}

func NewFlashImageMessageChain(imageId string, url string, path string, base64 string) FlashImageMessageChain {
	return FlashImageMessageChain{
		ImageMessageChain: ImageMessageChain{
			AbstractMessageChain: NewAbstractMessageChain("FlashImage"),
			ImageID:              imageId,
			URL:                  url,
			Path:                 path,
			Base64:               base64},
	}
}
