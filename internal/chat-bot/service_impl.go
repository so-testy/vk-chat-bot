package chat_bot

import (
	"vk-chat-bot/internal/connection"
	"vk-chat-bot/internal/repository"
)

type ServiceImpl struct {
	connection connection.Connection
	repository repository.Repository
}

// NewService - функция создания сервиса ChatBot
func NewService(conn connection.Connection, repo repository.Repository) Service {
	return &ServiceImpl{
		connection: conn,
		repository: repo,
	}
}
