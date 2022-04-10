package chat_bot

import (
	"context"
	"github.com/SevereCloud/vksdk/v2/events"
)

type Service interface {
	NewMessage(ctx context.Context, obj events.MessageNewObject)
}
