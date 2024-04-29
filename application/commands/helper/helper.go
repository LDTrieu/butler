package helper

import (
	"butler/constants"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func CheckPrefixCommand(message string, command string) bool {
	return strings.HasPrefix(message, constants.BOT_COMMAND_PREFIX+command)
}

func CheckMention(m *discordgo.MessageCreate, user *discordgo.User) bool {
	for _, item := range m.Mentions {
		if item.ID == user.ID {
			return true
		}
	}
	return false
}
