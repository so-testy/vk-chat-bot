package vk

import (
	"github.com/SevereCloud/vksdk/v2/api"
	"vk-chat-bot/internal/connection"
)

type ConnectionImpl struct {
	vk *api.VK
}

// NewConnection - функция создания новго соединения с API VK
func NewConnection(vk *api.VK) connection.Connection {
	return &ConnectionImpl{vk: vk}
}
