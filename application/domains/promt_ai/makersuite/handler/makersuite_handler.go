package handler

import (
	initServices "butler/application/domains/services/init"
	promtAiSv "butler/application/domains/services/promt_ai/service"
	"butler/config"
	"butler/constants"
	"regexp"

	"github.com/bwmarrin/discordgo"
)

type Handler struct {
	cfg       *config.Config
	promtAiSv promtAiSv.IService
}

func InitHandler(
	cfg *config.Config,
	services *initServices.Services,
) Handler {
	return Handler{
		promtAiSv: services.PromtAiSv,
		cfg:       cfg,
	}
}

func (h Handler) Ask(s *discordgo.Session, m *discordgo.MessageCreate) error {
	message := preprocessingMessage(m.Content)
	res, err := h.promtAiSv.Ask(message)
	if err != nil {
		return err
	}
	s.ChannelMessageSend(m.ChannelID, res) // nolint
	return nil
}

func preprocessingMessage(message string) string {
	reg := regexp.MustCompile(constants.USER_ID_REGEX)
	message = reg.ReplaceAllString(message, "")
	return message
}
