package define

import (
	"could-work/backend/util"
	"os"

	"path/filepath"

	"github.com/Fromsko/gouitls/logs"
)

var (
	Log       = logs.InitLogger()
	CONFIG, _ = util.InitConfig()
	UserName  = CONFIG.Username         // "202127530334"
	PassWord  = CONFIG.Password         // "Lyhqyh99@"
	BaseUrl   = CONFIG.HtmlURL.BaseUrl  // "https://jy.hniu.cn/login"
	TokenUrl  = CONFIG.HtmlURL.TokenUrl // "http://localhost:20000/api/v1/login"
	CodeUrl   = CONFIG.HtmlURL.CodeUrl  // "http://localhost:20000/api/v1/verify-code"
)

var (
	WorkPath, _         = os.Getwd()
	PromptPath          = filepath.Join(WorkPath, "res", "prompt")
	ImgPath             = filepath.Join(WorkPath, "res", "img")
	NotifyImg           = filepath.Join(ImgPath, "气泡通知.png")
	LoginImg            = filepath.Join(ImgPath, "登录页.png")
	LogoutImg           = filepath.Join(ImgPath, "退出页")
	ValidImg            = filepath.Join(ImgPath, "验证码.png")
	IndexImg            = filepath.Join(ImgPath, "主页.png")
	InternshipNotice    = filepath.Join(ImgPath, "实习通知.png")
	InternshipJournal   = filepath.Join(ImgPath, "实习日志.png")
	InternshipDiaryList = filepath.Join(ImgPath, "%s实习日志.png")
)

var (
	Title   = "云就业平台"
	Version = "1.0"
)
