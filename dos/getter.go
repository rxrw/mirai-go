package dos

import "encoding/json"

func WhichMessageChain(messageChan map[string]interface{}) interface{} {
	chanMap := make(map[string]interface{})
	chanMap["Source"] = SourceMessageChain{}
	chanMap["Quote"] = QuoteMessageChain{}
	chanMap["At"] = AtMessageChain{}
	chanMap["AtAll"] = AtAllMessageChain{}
	chanMap["Face"] = FaceMessageChain{}
	chanMap["Plain"] = PlainMessageChain{}
	chanMap["Image"] = ImageMessageChain{}
	chanMap["FlashImage"] = FlashImageMessageChain{}
	chanMap["Voice"] = VoiceMessageChain{}
	chanMap["Xml"] = XmlMessageChain{}
	chanMap["Json"] = JsonMessageChain{}
	chanMap["App"] = AppMessageChain{}
	chanMap["Poke"] = PokeMessageChain{}
	chanMap["Dice"] = DiceMessageChain{}
	chanMap["MusicShare"] = MusicShareMessageChain{}
	chanMap["Forward"] = ForwardMessageChain{}
	chanMap["File"] = FileMessageChain{}
	chanMap["MiraiCode"] = MiraiCodeMessageChain{}

	message, ok := chanMap[messageChan["type"].(string)]
	if !ok {
		// 事件
		return nil
	}

	str, _ := json.Marshal(messageChan)

	json.Unmarshal(str, &message)
	return message
}
