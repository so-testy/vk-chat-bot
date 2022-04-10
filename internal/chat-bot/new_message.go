package chat_bot

import (
	"context"
	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/SevereCloud/vksdk/v2/events"
	"log"
	"math/rand"
	"time"
)

func (s ServiceImpl) NewMessage(ctx context.Context, obj events.MessageNewObject) {
	rand.Seed(time.Now().UTC().UnixNano())

	switch obj.Message.Text {
	case "Привет":
		messageBuilder := params.NewMessagesSendBuilder()
		messageBuilder.RandomID(rand.Int())
		messageBuilder.PeerID(obj.Message.PeerID)
		messageBuilder.Message("Привет вездекодерам!")

		// TODO: Проверить, не будет ли админ сам себе получать ответ на "Привет"

		if err := s.connection.SendByParams(messageBuilder.Params); err != nil {
			log.Println(err)
		}
	}
}
