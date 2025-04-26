package captcha

import (
	"bytes"
	"could-work/backend/core/define"
	"log"
	"os"
)

func IdentifyCode() string {
	req := VerifyRequest{
		File:     &bytes.Buffer{},
		TokenUrl: define.TokenUrl,
		CodeUrl:  define.CodeUrl,
		Auth: Auth{
			UserName: "admin",
			PassWord: "admin",
		},
	}
	req.GetToken()
	if content, err := os.ReadFile(define.ValidImg); err != nil {
		log.Fatalf("验证码读取错误: %v", err)
	} else {
		req.Recognize(content)
	}
	return req.Code
}
