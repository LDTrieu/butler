package handler

import (
	"butler/application/domains/pick/models"
	"butler/application/domains/pick/usecase"
	initServices "butler/application/domains/services/init"
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	usecase usecase.IUseCase
}

func InitHandler(services *initServices.Services) Handler {
	usecase := usecase.InitUseCase(services)
	return Handler{
		usecase,
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
