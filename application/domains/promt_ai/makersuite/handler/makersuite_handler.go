package handler

import (
	services "butler/application/domains/services/init"
	"butler/config"

	"github.com/bwmarrin/discordgo"
)

type Handler struct {
	cfg     *config.Config
}

func InitHandler(
	cfg *config.Config,
	services makersuite.IService,
) Handler {
	return Handler{
		service: service,
		cfg:     cfg,
	}
}

func (h Handler) Ask(s *discordgo.Session, m *discordgo.MessageCreate) error {
	res, err := h.service.AskAnything(s, m)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, err.Error()) // nolint
		return err
	}
	s.ChannelMessageSend(m.ChannelID, res) // nolint
	return nil
}
