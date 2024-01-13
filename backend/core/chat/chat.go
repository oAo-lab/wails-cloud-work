package chat

import (
	"context"
	"could-work/backend/core/define"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"

	openai "github.com/sashabaranov/go-openai"
)

type WorkBot struct {
	BaseKey   string         // 密钥
	BaseURL   string         // api地址
	Proxy     string         // 代理地址
	MaxTokens int            // token 数量
	Client    *openai.Client // 连接器
}

func NewWorkBot(opts ...func(*WorkBot)) *WorkBot {

	w := &WorkBot{}

	for _, opt := range opts {
		opt(w)
	}

	if w.Proxy != "" {
		NewProxyClient(w)
	}

	if w.BaseKey != "" {
		w.Client = NewDefaultClient(w)
	}

	return w
}

func defaultConfig(key, url string) *openai.ClientConfig {
	config := openai.DefaultConfig(key)

	if url != "" {
		config.BaseURL = url
	}

	return &config
}

func NewDefaultClient(wb *WorkBot) *openai.Client {
	return openai.NewClientWithConfig(
		*defaultConfig(wb.BaseKey, wb.BaseURL),
	)
}

func NewProxyClient(client *WorkBot) *openai.Client {

	proxyUrl, _ := url.Parse(client.Proxy)
	config := defaultConfig(client.BaseKey, client.BaseURL)

	config.HTTPClient = &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyUrl),
		},
	}

	return openai.NewClientWithConfig(*config)
}

func LoadPrompt() string {
	prompt, err := os.ReadFile(define.PromptPath)
	if err != nil {
		return ""
	}
	return string(prompt)
}

func WithBaseKey(key string) func(*WorkBot) {
	return func(bot *WorkBot) {
		bot.BaseKey = key
	}
}

func WithBaseUrl(url string) func(*WorkBot) {
	return func(bot *WorkBot) {
		bot.BaseURL = url
	}
}

func WithProxy(proxy string) func(*WorkBot) {
	return func(wb *WorkBot) {
		wb.Proxy = proxy
	}
}

func WithClient(client *WorkBot) func(*WorkBot) {
	return func(wb *WorkBot) {
		wb.Client = client.Client
	}
}

func WithMaxTokens(maxTokens int) func(*WorkBot) {
	return func(wb *WorkBot) {
		wb.MaxTokens = maxTokens
	}
}

func (wb *WorkBot) Send(prompt string) (string, error) {
	if wb.Client == nil {
		return "", fmt.Errorf("client is not initialized")
	}

	req := &openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    "system",
				Content: LoadPrompt(),
			},
			{
				Role:    "user",
				Content: prompt,
			},
		},
	}

	if wb.MaxTokens != 0 {
		req.MaxTokens = wb.MaxTokens
	} else {
		req.MaxTokens = 1024
	}

	reply, err := sendWithTimeout(wb.Client, req, time.Second*60)
	return reply, err
}

func sendWithTimeout(client *openai.Client, req *openai.ChatCompletionRequest, timeout time.Duration) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	resultCh := make(chan string)

	go func() {
		resp, err := client.CreateChatCompletion(ctx, *req)
		if err != nil {
			resultCh <- fmt.Sprintf("completion error: %v", err)
			return
		}
		resultCh <- resp.Choices[0].Message.Content
	}()

	select {
	case <-ctx.Done():
		fmt.Println("Send timeout")
		return "", ctx.Err()

	case result := <-resultCh:
		return result, nil
	}
}

func NewChatBot() *WorkBot {
	workBot := NewWorkBot(
		WithMaxTokens(2000),
		WithBaseKey(define.CONFIG.OpenAI.Key),
		WithBaseUrl(define.CONFIG.OpenAI.BaseUrl),
	)
	return workBot
}
