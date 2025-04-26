package spider

import (
	"could-work/backend/core/define"

	rod "github.com/Fromsko/rodPro"
	"github.com/Fromsko/rodPro/lib/launcher"
)

var log = define.Log

type (
	// Web 自动化浏览器
	Web struct {
		Browser *rod.Browser
		Page    *rod.Page
	}
	// DiaryHandler 日志回调函数
	DiaryHandler map[string]*rod.Element
	// CallBack 回调函数
	CallBack func(web *Web, ele *rod.Element) error
)

// InitWeb 初始化浏览器
func InitWeb(URL string) *Web {
	u := launcher.New().Leakless(false).MustLaunch()
	w := &Web{
		Browser: rod.New().ControlURL(u).MustConnect(),
		Page:    nil,
	}
	w.Page = w.Browser.MustPage(URL)
	return w
}
