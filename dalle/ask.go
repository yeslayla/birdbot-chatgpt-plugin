package dalle

import (
	"context"
	"log"

	"github.com/sashabaranov/go-openai"
	"github.com/yeslayla/birdbot-common/common"
)

func (dalle *DALLE) Ask(user common.User, message string) string {
	ctx := context.Background()

	res, err := dalle.client.CreateImage(ctx, openai.ImageRequest{
		Prompt:         message,
		User:           user.ID,
		Size:           openai.CreateImageSize1024x1024,
		ResponseFormat: openai.CreateImageResponseFormatURL,
		Model:          openai.CreateImageModelDallE3,
		N:              1,
	})
	if err != nil {
		log.Println("Failed to create image: ", err)
		return ""
	}

	return res.Data[0].URL
}
