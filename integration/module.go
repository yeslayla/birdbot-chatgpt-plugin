package integration

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/yeslayla/birdbot-chatgpt-plugin/chatgpt"
	"github.com/yeslayla/birdbot-chatgpt-plugin/dalle"
	"github.com/yeslayla/birdbot-common/common"
)

type Module struct {
	chat    *Chat
	ChatGPT *chatgpt.ChatGPT
	DALLE   *dalle.DALLE
	Config  *Config
}

func (m *Module) Initialize(birdbot common.ModuleManager) error {

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

	m.ChatGPT = chatgpt.New(m.Config.OpenAIKey, m.Config.Prompts)
	m.DALLE = dalle.New(m.Config.OpenAIKey)

	if m.Config.EnableImageGeneration.IsEnabledByDefault() {
		dalle.RegisterDalleWithChatGPT(m.ChatGPT, m.DALLE)

		birdbot.RegisterCommand("generate", common.ChatCommandConfiguration{
			Description: "Generates an image from a prompt",
			Options: map[string]common.ChatCommandOption{
				"prompt": {
					Description: "The prompt to generate the image from",
					Type:        common.CommandTypeString,
					Required:    true,
				},
			},
			EphemeralResponse: false,
		}, func(u common.User, args map[string]any) string {
			return m.DALLE.Ask(u, fmt.Sprint(args["message"]))
		})
	}

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
