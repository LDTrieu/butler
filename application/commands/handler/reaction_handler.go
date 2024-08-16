package delivery

import (
	"butler/constants"
	"context"
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
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

	process, err := c.lib.Rdb.Get(ctx, getKeyPreparePick(r.MessageID)).Result()
	if err != nil {
		logrus.Errorf("Failed get prepare pick key: %v", err)
		return
	}
	if process == "" {
		return
	}

	switch r.Emoji.Name {
	case constants.EMOJI_NUMBER_ONE:
		s.ChannelMessageSend(r.ChannelID, "you choose warehouse 555 3 THANG 2")
	case constants.EMOJI_NUMBER_TWO:
		s.ChannelMessageSend(r.ChannelID, "you choose warehouse 29 HOANG VIET")
	default:
		// s.ChannelMessageSend(r.ChannelID, "you choose wrong warehouse")
		return
	}

	key := getKeyPreparePick(r.UserID)
	if err := c.lib.Rdb.Set(ctx, key, fmt.Sprintf("%v:", constants.MAP_EMOJI_WAREHOUSE[r.Emoji.Name]), 5*time.Minute).Err(); err != nil {
		logrus.Errorf("Failed set prepare pick key: %v", err)
		return
	}

	nextRequest, err := s.ChannelMessageSendEmbed(r.ChannelID, &discordgo.MessageEmbed{
		Title:       "Chọn đơn vị vận chuyển",
		Description: "Gõ số đơn để đi pick",
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:  constants.EMOJI_NUMBER_ONE + ": SHOP - 555 3 THANG 2",
				Value: "\n",
			},
			{
				Name:  constants.EMOJI_NUMBER_TWO + ": SHOP-29 HOANG VIET",
				Value: "\n",
			},
		},
	})
	if err != nil {
		logrus.Errorf("Error handle help command %v", err)
		return
	}

	if err := c.lib.Rdb.Set(ctx, getKeyPreparePick(r.MessageID), nextRequest.ID, 5*time.Minute).Err(); err != nil {
		logrus.Errorf("Failed set prepare pick key: %v", err)
		return
	}
}

func getKeyPreparePick(userId string) string {
	return fmt.Sprintf("prepare_pick:%s", userId)
}
