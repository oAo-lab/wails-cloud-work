package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/electricbubble/go-toast"
	"github.com/panjf2000/ants/v2"
)

type (
	H map[string]string

	Config struct { // NOTE: 配置
		Username string  `json:"username"`
		Password string  `json:"password"`
		Proxy    string  `json:"proxy"`
		Email    Email   `json:"email"`
		HtmlURL  HtmlURL `json:"html_url"`
		OpenAI   OpenAI  `json:"open-ai"`
		QBot     QBot    `json:"qq-bot"`
	}

	Email struct { // NOTE: 邮箱
		Host   string `json:"host"`
		Sender string `json:"sender"`
		Key    string `json:"key"`
	}

	HtmlURL struct { // NOTE: html
		BaseUrl  string `json:"base_url"`
		TokenUrl string `json:"token_url"`
		CodeUrl  string `json:"code_url"`
	}

	OpenAI struct { // NOTE: open-ai
		Key     string `json:"key"`
		BaseUrl string `json:"url"`
	}

	QBot struct { // NOTE: q-bot
		AppID uint64 `json:"appid"`
		Token string `json:"token"`
	}

	Proxy struct {
		HTTPProxy  string `json:"http_proxy"`
		HTTPSProxy string `json:"https_proxy"`
	}
)

func InitConfig() (*Config, error) {
	var (
		f        *os.File
		err      error
		conf     = &Config{}
		confName = "config.json"
	)

	if FileExists(confName) {

		f, err = os.Open(confName)
		if err != nil {
			return nil, fmt.Errorf("failed to open config file: %w", err)
		}

		err = json.NewDecoder(f).Decode(&conf)
		if err != nil {
			return nil, fmt.Errorf("failed to decode config file: %w", err)
		}

	} else {

		f, err = os.Create(confName)
		if err != nil {
			return nil, fmt.Errorf("failed to create config file: %w", err)
		}

		err = json.NewEncoder(f).Encode(&Config{})
		if err != nil {
			return nil, fmt.Errorf("failed to write config file: %w", err)
		} else {
			os.Exit(0)
		}

	}

	if conf.Proxy != "" {
		SetProxy(&Proxy{
			HTTPProxy:  conf.Proxy,
			HTTPSProxy: conf.Proxy,
		})

	}

	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	return conf, nil
}

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	if err == nil {
		return true
	}

	if errors.Is(err, os.ErrNotExist) {
		return false
	}
	return false
}

func TaskRunner(tasks ...func()) {

	pool, _ := ants.NewPool(len(tasks))
	for _, t := range tasks {
		_ = pool.Submit(t)
	}

	defer pool.Release()
}

func SetProxy(config *Proxy) {

	if config.HTTPProxy != "" {
		_ = os.Setenv("HTTP_PROXY", config.HTTPProxy)
	} else {
		_ = os.Unsetenv("HTTP_PROXY")
	}

	if config.HTTPSProxy != "" {
		_ = os.Setenv("HTTPS_PROXY", config.HTTPSProxy)
	} else {
		_ = os.Unsetenv("HTTPS_PROXY")
	}
}

func Toask(msg H) {
	_ = toast.Push(
		msg["data"],
		toast.WithIcon(msg["logo"]),
		toast.WithTitle(msg["title"]),
		toast.WithAppID(msg["app_id"]),
		toast.WithProtocolAction("🎉 Finished"),
	)
}
