package vk

import "vk-chat-bot/internal/connection"

type ConnectionImpl struct {
}

// NewConnection - функция создания новго соединения с API VK
func NewConnection() connection.Connection {
	return &ConnectionImpl{}
}
