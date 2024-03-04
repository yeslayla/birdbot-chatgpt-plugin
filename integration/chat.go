package integration

import (
	"math/rand"

	"github.com/yeslayla/birdbot-common/common"
)

type Chat struct {
	module *Module

	chance float64
	bird   common.ExternalChatManager
}

func (chat *Chat) Initialize(bird common.ExternalChatManager) {
	chat.bird = bird

	if chat.module.Config.ResponseChance == 0 {
		chat.chance = 1
	} else {
		chat.chance = chat.module.Config.ResponseChance / 100.0
	}
}

func (chat *Chat) RecieveMessage(user common.User, message string) {
	if rand.Float64() <= chat.chance {
		return
	}

	go func() {
		content := chat.module.ChatGPT.Ask(user, message, "chat")
		chat.bird.SendMessage("BirdBot", content)
	}()

}

func (m *Module) NewChat() *Chat {
	return &Chat{
		module: m,
	}
}
