package chat_bot

import (
	chatbot "vk-chat-bot/internal/chat-bot"
)

type Controller struct {
	service chatbot.Service
}

// NewController - функция создания контроллера chat-bot сервиса
func NewController(service chatbot.Service) Controller {
	return Controller{service: service}
}
