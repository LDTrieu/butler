package delivery

import (
	"butler/application/commands/helper"
	"butler/constants"
	"strings"

	makersuiteHandler "butler/application/domains/promt_ai/makersuite/handler"

	"github.com/bwmarrin/discordgo"
)

type Handler interface {
	GetCommandsHandler(*discordgo.Session, *discordgo.MessageCreate)
}

type commandHandler struct {
	discord           *discordgo.Session
	makersuiteHandler makersuiteHandler.Handler
}

func NewCommandHandler(
	discord *discordgo.Session,
	makersuiteHandler makersuiteHandler.Handler,
) Handler {
	return &commandHandler{
		discord:           discord,
		makersuiteHandler: makersuiteHandler,
	}
}

func (c *commandHandler) GetCommandsHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	if !strings.HasPrefix(m.Content, constants.BOT_COMMAND_PREFIX) && !helper.CheckMention(m, s.State.User) {
		return
	}

	switch {
	case helper.CheckPrefixCommand(m.Content, constants.COMMAND_HELP):
		_ = helper.HandleHelpCommand(s, m)
	case helper.CheckMention(m, s.State.User):
		_ = c.makersuiteHandler.Ask(s, m)
	}
}
