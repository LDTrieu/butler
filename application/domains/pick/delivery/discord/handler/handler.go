package handler

import (
	"butler/application/domains/pick/models"
	"butler/application/domains/pick/usecase"
	initServices "butler/application/domains/services/init"
	"butler/application/lib"
	"butler/constants"
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	lib     *lib.Lib
	usecase usecase.IUseCase
}

func InitHandler(lib *lib.Lib, services *initServices.Services) Handler {
	usecase := usecase.InitUseCase(lib, services)
	return Handler{
		lib:     lib,
		usecase: usecase,
	}
}

func (h Handler) ReadyPickOutbound(s *discordgo.Session, m *discordgo.MessageCreate) error {
	reg := regexp.MustCompile(`[0-9]+`)
	saleOrderNumber := reg.FindString(m.Content)
	logrus.Infof("Prepare for outbound [%s] to ready pick", saleOrderNumber)
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(30*time.Second))
	defer cancel()

	if err := h.usecase.ReadyPickOutbound(ctx, &models.ReadyPickOutboundRequest{
		SalesOrderNumber: saleOrderNumber,
	}); err != nil {
		logrus.Errorf("Failed ready outbound: %v", err)
		return err
	}
	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Outbound [%s] is ready to be picked!", saleOrderNumber))
	return nil
}

func (h Handler) PreparePick(s *discordgo.Session, m *discordgo.MessageCreate) error {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(30*time.Second))
	defer cancel()

	response, err := s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:       "Chọn kho đi pick",
		Description: "React vào emoji để chọn kho đi pick",
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
		return err
	}
	s.MessageReactionAdd(m.ChannelID, response.ID, constants.EMOJI_NUMBER_ONE)
	s.MessageReactionAdd(m.ChannelID, response.ID, constants.EMOJI_NUMBER_TWO)

	if err := h.lib.Rdb.Set(ctx, response.ID, "waiting select warehouse", 5*time.Minute).Err(); err != nil {
		return err
	}

	return nil
}

// func (h Handler) Pick(s *discordgo.Session, m *discordgo.MessageCreate) error {
// 	reg := regexp.MustCompile(`[0-9]+`)
// 	saleOrderNumber := reg.FindString(m.Content)
// 	logrus.Infof("Prepare for outbound [%s] to pick", saleOrderNumber)
// 	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(30*time.Second))
// 	defer cancel()

// 	if err := h.usecase.Pick(ctx, &models.PickRequest{
// 		SalesOrderNumber: saleOrderNumber,
// 	}); err != nil {
// 		logrus.Errorf("Failed pick outbound: %v", err)
// 		return err
// 	}
// 	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Outbound [%s] is picked!", saleOrderNumber))
// 	return nil
// }
