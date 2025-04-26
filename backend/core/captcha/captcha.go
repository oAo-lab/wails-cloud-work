package captcha

import (
	"bytes"
	"could-work/backend/core/define"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"

	"github.com/Fromsko/gouitls/knet"
	"github.com/tidwall/gjson"
)

type VerifyRequest struct {
	File     *bytes.Buffer
	Token    string
	Code     string
	TokenUrl string
	CodeUrl  string
	Auth     Auth
}

type Auth struct {
	UserName string
	PassWord string
}

func (r *VerifyRequest) GetToken() {
	sendRequest := knet.SendRequest{
		Method:   "POST",
		FetchURL: r.TokenUrl,
		Data: strings.NewReader((url.Values{
			"username": {r.Auth.UserName},
			"password": {r.Auth.PassWord},
		}).Encode()),
		Headers: map[string]string{
			"Content-Type": "application/x-www-form-urlencoded",
		},
	}
	sendRequest.Send(func(resp []byte, cookies []*http.Cookie, err error) {
		if gjson.Get(string(resp), "status").Int() != 200 {
			msg := gjson.Get(string(resp), "err").String()
			define.Log.Errorf("获取Token失败: %v", msg)
		}
		r.Token = gjson.Get(string(resp), "token").String()
	})
}

func (r *VerifyRequest) Recognize(content []byte) {
	writer := multipart.NewWriter(r.File)
	if part, err := writer.CreateFormFile("file", "code.png"); err == nil {
		_, _ = part.Write(content)
		_ = writer.Close()
	}

	sendRequest := knet.SendRequest{
		Name:     "Recognize",
		Method:   "POST",
		FetchURL: r.CodeUrl,
		Data:     r.File,
		Headers: map[string]string{
			"accept":        "application/json",
			"Authorization": "Bearer " + r.Token,
			"Content-Type":  writer.FormDataContentType(),
		},
	}

	sendRequest.Send(func(resp []byte, cookies []*http.Cookie, err error) {
		if gjson.Get(string(resp), "status").Int() != 200 {
			msg := gjson.Get(string(resp), "err").String()
			define.Log.Errorf("识别失败: %s", msg)
		}
		r.Code = gjson.Get(string(resp), "msg").String()
		define.Log.Infof("识别到验证码: %s", r.Code)
	})
}
