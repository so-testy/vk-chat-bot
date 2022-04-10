package chat_bot

import (
	"context"
	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/SevereCloud/vksdk/v2/events"
	"log"
	"math/rand"
	"strings"
	"time"
)

func (s ServiceImpl) NewMessage(ctx context.Context, obj events.MessageNewObject) {
	rand.Seed(time.Now().UTC().UnixNano())

	text := strings.ToLower(obj.Message.Text)

	switch text {
	case "привет":
		messageBuilder := params.NewMessagesSendBuilder()
		messageBuilder.RandomID(rand.Int())
		messageBuilder.PeerID(obj.Message.PeerID)
		messageBuilder.Message("Привет вездекодерам!")

		// TODO: Проверить, не будет ли админ сам себе получать ответ на "Привет"

		if err := s.connection.SendByParams(messageBuilder.Params); err != nil {
			log.Println(err)
		}

	default:
		messageBuilder := params.NewMessagesSendBuilder()
		messageBuilder.RandomID(rand.Int())
		messageBuilder.PeerID(obj.Message.PeerID)
		messageBuilder.Message("Дороу бандит! Остальной функционал пока не завезли, а пока можешь онуть с мемема и лайкнуть его - https://vk.com/photo-197700721_457240802")

		// TODO: Проверить, не будет ли админ сам себе получать ответ на "Привет"

		if err := s.connection.SendByParams(messageBuilder.Params); err != nil {
			log.Println(err)
		}
	}
}
