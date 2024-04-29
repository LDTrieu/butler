package service

import (
	"butler/config"
	"butler/constants"
	"context"
	"regexp"
	"time"

	"github.com/google/generative-ai-go/genai"
	"github.com/sirupsen/logrus"
)

type service struct {
	cfg        *config.Config
	genaiModel *genai.GenerativeModel
}

func InitService(
	cfg *config.Config,
	genaiClient *genai.Client,
) IService {
	model := genaiClient.GenerativeModel(cfg.Makersuite.Model)
	return &service{
		cfg:        cfg,
		genaiModel: model,
	}
}

func (sv *service) Ask(question string) (string, error) {
	ctx, cancelCtx := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancelCtx()

	// logrus.Infof("User [%s] needs help: %s", m.Author.Username, m.Content)
	// message := preprocessingMessage(m.Content)
	response, err := sv.genaiModel.GenerateContent(ctx, genai.Text(question))
	if err != nil {
		logrus.Errorf("get response failed with err=%v", err)
		return "", err
	}

	if len(response.Candidates) == 0 {
		return "I really don't know", nil
	}

	return getAllGenaiResponse(response), nil
}

func getAllGenaiResponse(resp *genai.GenerateContentResponse) string {
	response := ""
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				response += string(part.(genai.Text))
			}
		}
	}
	return response
}

func preprocessingMessage(message string) string {
	reg := regexp.MustCompile(constants.USER_ID_REGEX)
	message = reg.ReplaceAllString(message, "")
	return message
}
