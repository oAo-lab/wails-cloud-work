package bot

import (
	"context"
	"strings"

	"github.com/tencent-connect/botgo/dto"
	"github.com/tencent-connect/botgo/dto/message"
)

// ProcessMessage 消息处理流
func (p Processor) ProcessMessage(input string, data *dto.WSATMessageData) error {
	ctx := context.Background()
	cmd := message.ParseCommand(input)
	toCreate := &dto.MessageToCreate{
		Content: "默认回复" + message.Emoji(307),
		MessageReference: &dto.MessageReference{
			// 引用这条消息
			MessageID:             data.ID,
			IgnoreGetMessageError: true,
		},
	}

	log.Infof("执行命令: %s \n 频道ID: %s", cmd, data.ID)

	switch cmd.Cmd {

	case "/dm":
		p.dmHandler(data)

	case "/chat":
		p.sendChatReply(ctx, cmd.Content, data.ChannelID, toCreate)

	case "/test":
		switch cmd.Content {
		case "hi":
			p.sendReply(ctx, data.ChannelID, toCreate)
		case "time":
			toCreate.Content = genReplyContent(data)
			p.sendReply(ctx, data.ChannelID, toCreate)
		case "ark":
			toCreate.Ark = genReplyArk(data)
			p.sendReply(ctx, data.ChannelID, toCreate)
		case "公告":
			p.setAnnounces(ctx, data)
		case "pin":
			if data.MessageReference != nil {
				p.setPins(ctx, data.ChannelID, data.MessageReference.MessageID)
			}
		case "emoji":
			if data.MessageReference != nil {
				p.setEmoji(ctx, data.ChannelID, data.MessageReference.MessageID)
			}
		default:
			toCreate.Image = "https://c-ssl.duitang.com/uploads/blog/202207/09/20220709150824_97667.jpg"
			p.sendReply(ctx, data.ChannelID, toCreate)
		}
	}
	return nil
}

// PublicMessage 公开信息
func (p Processor) PublicMessage(cmd *message.CMD, data *dto.WSMessageData) error {

	toCreate := &dto.MessageToCreate{
		MessageReference: &dto.MessageReference{
			MessageID:             data.ID,
			IgnoreGetMessageError: true,
		},
	}

	p.sendChatReply(
		context.Background(),
		cmd.Content,
		data.ChannelID,
		toCreate,
	)

	return nil
}

// PasserMessage 解析命令
func PasserMessage(content string) *message.CMD {
	return message.ParseCommand(
		strings.ToLower(message.ETLInput(content)),
	)
}
