package global

import (
	"context"
	"fmt"
	"github.com/JathamJ/SagaFlow/config"
	"github.com/JathamJ/SagaFlow/utils"
	"github.com/cloudwego/eino-ext/components/model/qwen"
	"github.com/cloudwego/eino/components/model"
	"net/http"
)

type ServiceContext struct {
	ChatModel model.BaseChatModel
}

var svc *ServiceContext

func Init(cfg *config.Config) error {
	svc = new(ServiceContext)
	ctx := context.Background()
	var err error

	svc.ChatModel, err = qwen.NewChatModel(ctx, &qwen.ChatModelConfig{
		APIKey:           cfg.ChatModel.ApiKey,
		Timeout:          30,
		HTTPClient:       &http.Client{},
		BaseURL:          cfg.ChatModel.BaseUrl,
		Model:            cfg.ChatModel.Model,
		MaxTokens:        utils.Of(2048),
		Temperature:      nil,
		TopP:             nil,
		Stop:             nil,
		PresencePenalty:  nil,
		ResponseFormat:   nil,
		Seed:             nil,
		FrequencyPenalty: nil,
		LogitBias:        nil,
		User:             nil,
		EnableThinking:   nil,
		Modalities:       nil,
		Audio:            nil,
	})
	if err != nil {
		return fmt.Errorf("init chatModel failed, err: %v", err)
	}

	return nil
}

func Ctx() *ServiceContext {
	return svc
}
