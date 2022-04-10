package chat_bot

import (
	"github.com/SevereCloud/vksdk/v2/callback"
	chatbot "vk-chat-bot/internal/chat-bot"
)

type Controller struct {
	callback *callback.Callback
	service  chatbot.Service
}

// NewController - функция создания контроллера chat-bot сервиса
func NewController(callback *callback.Callback, service chatbot.Service) Controller {
	return Controller{
		callback: callback,
		service:  service,
	}
}

// Run - метод запуска обработки приходящих событий
func (c Controller) Run() {

	//Обработка калбека нового сообщения
	c.callback.MessageNew(c.service.NewMessage)
}
