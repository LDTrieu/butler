package delivery

import (
	"butler/constants"
	"context"
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"

	whModel "butler/application/domains/warehouse/models"
)

func (c *commandHandler) GetReactionHandler(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	defer func() {
		if err := recover(); err != nil {
			logrus.Errorf("runtime error: %v", err)
			s.ChannelMessageSend(r.ChannelID, fmt.Sprintf("Something went wrong: %v", err))
		}
	}()
	if s.State.User.ID == r.MessageReaction.UserID {
		return
	}
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(30*time.Second))
	defer cancel()

	if value := c.lib.Cache.Get(getKeyWhConfig(r.MessageID)); value != nil {
		// TODO: write separate function to handle this
		c.lib.Cache.Delete(getKeyWhConfig(r.MessageID))
		request, ok := value.(*whModel.UpdateConfigWarehouseRequest)
		if !ok {
			s.ChannelMessageSend(r.ChannelID, fmt.Sprintf("cache error: failed to parse request: %#v", value))
			return
		}
		if request == nil {
			s.ChannelMessageSend(r.ChannelID, "Update config warehouse failed, please try again")
			return
		}
		switch r.Emoji.Name {
		case constants.EMOJI_CHAR_A:
			request.Config = constants.WAREHOUSE_CONFIG_ENABLE_LOCATION
		default:
			return
		}

		if err := c.whHandler.UpdateConfigWarehouse(ctx, request); err != nil {
			s.ChannelMessageSend(r.ChannelID, err.Error())
			return
		} else {
			s.ChannelMessageSend(r.ChannelID, "thêm config thành công")
		}
	}
}

func getKeyWhConfig(messageId string) string {
	return messageId + "::" + constants.CACHE_KEY_WH_CONFIG
}
