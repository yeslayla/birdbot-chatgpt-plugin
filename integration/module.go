package integration

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/yeslayla/birdbot-chatgpt-plugin/chatgpt"
	"github.com/yeslayla/birdbot-common/common"
)

type Module struct {
	chat    *Chat
	ChatGPT *chatgpt.ChatGPT
	Config  *Config
}

func (m *Module) Initialize(birdbot common.ModuleManager) error {

	rand.Seed(time.Now().UnixNano())

	configFile := birdbot.GetConfigPath("chatgpt.yaml")
	log.Printf("Using config: %s", configFile)
	m.Config = &Config{}

	_, err := os.Stat(configFile)
	if errors.Is(err, os.ErrNotExist) {
		log.Printf("Config file not found: '%s'", configFile)
		err := cleanenv.ReadEnv(&m.Config)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err := cleanenv.ReadConfig(configFile, m.Config)
		if err != nil {
			log.Fatal(err)
		}
	}

	m.ChatGPT = chatgpt.NewChatGPT(m.Config.OpenAIKey, m.Config.Prompts)

	m.chat = m.NewChat()

	if m.Config.EnableChat.IsEnabled() {
		birdbot.RegisterExternalChat("chatgpt", m.chat)
	}

	if m.Config.EnableCommand.IsEnabledByDefault() {
		birdbot.RegisterCommand("ask", common.ChatCommandConfiguration{
			Description: "Asks the bot a question with ChatGPT",
			Options: map[string]common.ChatCommandOption{
				"message": {
					Description: "Message to be sent to bot",
					Type:        common.CommandTypeString,
					Required:    true,
				},
			},
			EphemeralResponse: false,
		}, func(u common.User, args map[string]any) string {
			return m.ChatGPT.Ask(u, fmt.Sprint(args["message"]), u.ID)
		})
	}

	return nil
}

func NewModule() common.Module {
	return &Module{}
}
