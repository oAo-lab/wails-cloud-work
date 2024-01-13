package util

import "sync"

// MessageQueue 定义消息队列结构体
type MessageQueue struct {
	messages []*Message
	mutex    sync.Mutex
}

type Message struct {
	Type    string `json:"type"`
	Payload any    `json:"payload"`
}

// NewMessageQueue 初始化消息队列
func NewMessageQueue() *MessageQueue {
	return &MessageQueue{
		messages: make([]*Message, 0),
	}
}

// IsEmpty 判断消息队列是否为空
func (mq *MessageQueue) IsEmpty() bool {
	mq.mutex.Lock()
	defer mq.mutex.Unlock()
	return len(mq.messages) == 0
}

// Push 向消息队列中加入消息
func (mq *MessageQueue) Push(message *Message) {
	mq.mutex.Lock()
	defer mq.mutex.Unlock()
	mq.messages = append(mq.messages, message)
}

// Pop 从消息队列中取出消息
func (mq *MessageQueue) Pop() *Message {
	mq.mutex.Lock()
	defer mq.mutex.Unlock()

	if len(mq.messages) == 0 {
		return nil
	}

	message := mq.messages[0]
	mq.messages = mq.messages[1:]
	return message
}

var MsgQueue = &MessageQueue{}
var TaskQueue = &MessageQueue{}