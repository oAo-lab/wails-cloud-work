package spider

import (
	"could-work/backend/core/captcha"
	"could-work/backend/core/define"
)

// SearchParams 查找参数
func (w *Web) SearchParams(text string) bool {
	search, err := w.Page.Search(text)
	if err != nil {
		return false
	}
	return search.ResultCount != 0
}

// LoginPage 登陆页面
func (w *Web) LoginPage() bool {
	page := w.Page
	page.MustWaitLoad()
	page.MustElement("#user_name").MustInput(define.UserName)
	page.MustElement("#password").MustInput(define.PassWord)
	page.MustElement("#img_valid_code").MustScreenshot(define.ValidImg)
	page.MustElement("#valid_code").MustInput(captcha.IdentifyCode())
	page.MustElement("#privacy-agreement").MustClick()
	page.MustElement("#login_btn").MustClick()
	page.MustScreenshot(define.LoginImg)
	return !(w.SearchParams("请输入正确的验证码") || w.SearchParams("生源不存在"))
}

// LogOutPage 退出登录
func (w *Web) LogOutPage() {
	w.Page.MustWaitLoad().MustSearch("退出").MustClick()
}

// IndexPage 主页面
func (w *Web) IndexPage(callBack CallBack) {
	w.Page.MustWaitStable().MustScreenshot(define.IndexImg)
	w.Page.MustWaitLoad().MustSearch("实习管理").MustClick()
	elements := w.Page.MustWaitLoad().MustElements("#tab.nav.nav-tabs li")
	for index, element := range elements {
		//fmt.Println(index, element.MustText())
		switch index {
		case 1:
			// TODO: 实习通知
			log.Info("实习通知")
			element.MustClick()
			notice := w.Page.MustElement("#n_grid.grid").MustWaitLoad()
			notice.MustWaitStable().MustScreenshot(define.InternshipNotice)
		case 3:
			// TODO: 实习日志
			log.Info("实习日志")
			element.MustClick()
			journal := w.Page.MustElement("#tab_3").MustWaitLoad()
			_ = callBack(w, journal)
		}
	}
}
