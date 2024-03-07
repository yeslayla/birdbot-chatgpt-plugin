package integration

import (
	"math/rand"

	"github.com/yeslayla/birdbot-common/common"
)

type Chat struct {
	module *Module

	locked chan bool

	chance float64
	bird   common.ExternalChatManager
}

func (chat *Chat) Initialize(bird common.ExternalChatManager) {
	chat.bird = bird

	// Unlocked by default
	chat.locked = make(chan bool, 1)
	chat.locked <- true

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
		// Wait for channel to be unlocked
		<-chat.locked

		content := chat.module.ChatGPT.Ask(user, message, "chat")
		if content == "" {
			chat.bird.SendMessage("BirdBot", "I'm sorry, having trouble understanding you right now.")
			return
		}
		chat.bird.SendMessage("BirdBot", content)

		// Unlock the channel by sending a value
		chat.locked <- true
	}()

}

func (m *Module) NewChat() *Chat {
	return &Chat{
		module: m,
	}
}
