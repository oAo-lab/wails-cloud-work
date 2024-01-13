package bot

import (
	"context"
	"could-work/backend/core/define"
	"could-work/backend/util"
	"time"

	"github.com/tencent-connect/botgo"
	"github.com/tencent-connect/botgo/openapi"
	"github.com/tencent-connect/botgo/token"
	"github.com/tencent-connect/botgo/websocket"
)

var (
	log     = define.Log
	api     openapi.OpenAPI
	dispose = &Processor{}
)

func RegisterBot(config *util.QBot, env string) {

	intent := websocket.RegisterHandlers(
		// at 机器人事件，目前是在这个事件处理中有逻辑，会回消息，其他的回调处理都只把数据打印出来，不做任何处理
		ATMessageEventHandler(),
		// 如果想要捕获到连接成功的事件，可以实现这个回调
		ReadyHandler(),
		// 连接关闭回调
		ErrorNotifyHandler(),
		// 频道事件
		GuildEventHandler(),
		// 成员事件
		MemberEventHandler(),
		// 子频道事件
		ChannelEventHandler(),
		// 私信，目前只有私域才能够收到这个，如果你的机器人不是私域机器人，会导致连接报错，那么启动 example 就需要注释掉这个回调
		DirectMessageHandler(),
		// 频道消息，只有私域才能够收到这个，如果你的机器人不是私域机器人，会导致连接报错，那么启动 example 就需要注释掉这个回调
		CreateMessageHandler(),
		// 互动事件
		InteractionHandler(),
		// 发帖事件
		ThreadEventHandler(),
	)

	token := token.BotToken(config.AppID, config.Token)

	// 初始化机器人
	switch env {
	case "release":
		api = botgo.NewOpenAPI(token)
	default:
		api = botgo.NewSandboxOpenAPI(token)
	}

	// 连接超时
	api.WithTimeout(3 * time.Second)
	dispose.API = api

	// Websocket 连接
	wsInfo, err := api.WS(context.Background(), nil, "")
	if err != nil {
		log.Errorf("WebSocket connection error %s", err)
	}

	// 指定需要启动的分片数为 2 的话可以手动修改 wsInfo
	if err := botgo.NewSessionManager().Start(wsInfo, token, &intent); err != nil {
		log.Errorf("Start bot failed :> %s", err)
	}
}
