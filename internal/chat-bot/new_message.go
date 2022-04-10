package chat_bot

import (
	"context"
	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/SevereCloud/vksdk/v2/events"
	"log"
)

func (s ServiceImpl) NewMessage(ctx context.Context, obj events.MessageNewObject) {
	switch obj.Message.Text {
	case "Привет":
		messageBuilder := params.NewMessagesSendBuilder()
		messageBuilder.Message("Привет вездекодерам!")

		// TODO: Тут понять как пользователю сообщества отправлять сообщение

		if err := s.connection.SendByParams(messageBuilder.Params); err != nil {
			log.Println(err)
		}
	}
}
