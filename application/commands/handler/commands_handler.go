package delivery

import (
	"butler/application/commands/helper"
	"butler/application/commands/usecase"
	"butler/constants"
	"strings"

	accountHandler "butler/application/domains/account/delivery/discord/handler"
	attendanceHandler "butler/application/domains/attendance/delivery/discord/handler"
	makersuiteHandler "butler/application/domains/external/makersuite/handler"

	"github.com/bwmarrin/discordgo"
)

type Handler interface {
	GetCommandsHandler(*discordgo.Session, *discordgo.MessageCreate)
}

type commandHandler struct {
	discord           *discordgo.Session
	accHandler        accountHandler.DiscordHandler
	attendanceHandler attendanceHandler.Handler
	makersuiteHandler makersuiteHandler.Handler
}

func NewCommandsDelivery(
	discord *discordgo.Session,
	accHandler accountHandler.DiscordHandler,
	attendanceHandler attendanceHandler.Handler,
	makersuiteHandler makersuiteHandler.Handler,
) Handler {
	return &commandHandler{
		discord:           discord,
		accHandler:        accHandler,
		attendanceHandler: attendanceHandler,
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
		_ = usecase.HandleHelpCommand(s, m)
	case helper.CheckMention(m, s.State.User):
		_ = c.makersuiteHandler.Ask(s, m)
	case helper.CheckPrefixCommand(m.Content, constants.COMMAND_SAVE_USER):
		_ = c.accHandler.CreateAccount(s, m)
	case helper.CheckPrefixCommand(m.Content, constants.COMMAND_ROLL_CALL):
		_ = c.attendanceHandler.RollCall(s, m)
	case helper.CheckPrefixCommand(m.Content, "img"):
		_ = usecase.HandleSendImage(s, m)
	}
}
