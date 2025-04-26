package event

import (
	"could-work/backend/core/chat"
	"could-work/backend/core/define"
	"could-work/backend/core/spider"
	"could-work/backend/util"
	"fmt"
	rod "github.com/Fromsko/rodPro"
	"strings"
	"time"
)

var log = define.Log

func handleChat(msg *WebSocketMessage) {
	define.Log.Error("聊天:> ", msg)

	cb := chat.NewChatBot()

	reply, err := cb.Send(msg.Payload)
	if err != nil {
		reply = err.Error()
		define.Log.Errorf("Chat bot err: %s", err)
	} else {
		define.Log.Infof("Chat bot err: %s", reply)
	}

	util.MsgQueue.Push(&util.Message{
		Type:    "chat",
		Payload: reply,
	})
}

// RegisterWorkTask 注册工作任务
func RegisterWorkTask() {
	Web := spider.InitWeb(define.BaseUrl)

	if !Web.LoginPage() {
		log.Info("登录失败")
	} else {
		log.Info("登录成功")
	}

	Web.IndexPage(func(web *spider.Web, ele *rod.Element) error {
		web.SearchDiaryList(web.SearchLatestDiary(ele))
		handle := web.DiaryHandle(ele)

		go util.MsgQueue.Push(&util.Message{
			Type:    "chat",
			Payload: "请输入需要填写的日志, [格式: /task 内容|xxx哈哈哈]",
		})

		handle.WaitUserReceive(func(diaryHandler spider.DiaryHandler, message *util.Message) {

			task := message.Payload.(util.H)["task"]
			content := message.Payload.(util.H)["info"]
			log.Info("正在执行任务: ", task, content)

			go util.MsgQueue.Push(&util.Message{
				Type:    "chat",
				Payload: fmt.Sprintf("正在执行任务: %s ", task),
			})

			switch task {
			case "地点":
				diaryHandler["地点"].MustInput(content)
			case "内容":
				diaryHandler["内容"].MustInput(content)
			case "文件":
				diaryHandler["文件"].MustSetFiles(define.ValidImg)
			case "取消":
				diaryHandler["取消"].MustClick()
			case "提交":
			// dh["提交"].MustClick() // TODO: 会提交日志
			case "截图":
				go diaryHandler["截图"].MustScreenshot("实时截图")
			}
			time.Sleep(time.Second * 5)
		})

		return nil
	})

	// time.Sleep(11 * time.Hour) // TODO: 调试

	defer func() {
		if r := recover(); r != nil {
			define.Log.Errorf("Recover %s", r)
			Web.Browser.MustClose()
		} else {
			define.Log.Info("关闭成功")
			Web.LogOutPage()
			Web.Browser.MustClose()
		}

		defer func() {
			util.Toask(util.H{
				"data":  "运行结束",
				"title": define.Title,
				"logo":  define.NotifyImg,
			})
		}()
	}()
}

func ParseTaskArgs(message any) *util.H {
	if taskMeta := strings.Split(message.(string), "|"); len(taskMeta) != 2 {
		return nil
	} else {
		t := &util.H{
			"task": taskMeta[0],
			"info": taskMeta[1],
		}
		return t
	}
}

func handleTask(msg *WebSocketMessage) {

	log.Infof("接收到参数: %s %s", msg.Payload, msg.Type)

	if msg.Payload == "启动任务" {
		go RegisterWorkTask()
		util.MsgQueue.Push(&util.Message{
			Type:    "chat",
			Payload: "任务已经成功启动!",
		})
	} else {
		log.Info("开始解析参数: ", msg.Payload)
		if payload := ParseTaskArgs(msg.Payload); payload != nil {
			util.TaskQueue.Push(&util.Message{
				Type:    "receive",
				Payload: *payload,
			})
			log.Info(payload)
		} else {
			util.TaskQueue.Push(&util.Message{
				Type: "client",
			})
		}
	}
}

func handlePing(wb *WebSocketMessage) {
	// define.Log.Info("WebSocket write error:", wb)
}
