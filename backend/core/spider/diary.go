package spider

import (
	"could-work/backend/core/define"
	"could-work/backend/util"
	"fmt"
	"strconv"
	"time"

	rod "github.com/Fromsko/rodPro"
)

// SearchLatestDiary 实习日志列表
func (w *Web) SearchLatestDiary(ele *rod.Element) *rod.Element {
	log.Info("实习列表")
	return ele.MustElement("#diaryListOne").MustWaitLoad()
}

func (w *Web) SearchDiaryList(ele *rod.Element) {
	log.Info("遍历列表")
	diaryList := ele.MustWaitStable().MustElements("div.item.clearfix")

	if !diaryList.Empty() {
		log.Info("找到了")
		for index, diary := range diaryList {
			diary.MustWaitStable().MustScreenshot(
				fmt.Sprintf(define.InternshipDiaryList, strconv.Itoa(index)),
			)
			log.Info(index, diary)
		}
	}
}

func (w *Web) DiaryHandle(ele *rod.Element) DiaryHandler {
	log.Info("准备处理日志")

	// 实习日志数量记录
	quantity := ele.MustWaitLoad().MustElement("div.grid-toolbar.clearfix")
	quantity.MustWaitStable().MustScreenshot(define.InternshipJournal)
	quantity.MustElement("#diary_submit").MustClick()

	dialogPage := quantity.MustElementX(`/html/body/div[5]/div/table/tbody`)
	iframe := dialogPage.MustElementX("tr[2]/td/div/iframe").MustFrame().MustWaitLoad()

	log.Info(iframe)

	return DiaryHandler{
		"地点": iframe.MustElement("#address"),
		"内容": iframe.MustElement("#content"),
		"文件": iframe.MustElement("[type=file]"),
		"提交": iframe.MustElement("#btn_info_submit"),
		"取消": dialogPage.MustElementX("tr[1]/td/button"),
		"截图": dialogPage,
	}
}

func (diary DiaryHandler) WaitUserReceive(callBack func(DiaryHandler, *util.Message)) {
	log.Info("处理日志", "等待用户输入")

	timer := time.NewTimer(time.Second * 60)
	defer timer.Stop()

	for {
		select {
		case <-timer.C:
			return
		default:
			if !util.TaskQueue.IsEmpty() {
				message := util.TaskQueue.Pop()

				switch message.Type {
				case "client":
					log.Info("正在对话|重置时间")
					timer.Reset(time.Second * 60)
				case "receive":
					callBack(diary, message)
				default:
					log.Warnf("未知消息类型:> %s", message.Type)
				}
			}
		}
	}
}
