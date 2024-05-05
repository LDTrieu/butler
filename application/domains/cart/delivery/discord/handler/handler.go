package handler

import (
	"butler/application/domains/cart/models"
	"butler/application/domains/cart/usecase"
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

func (h Handler) ResetCart(s *discordgo.Session, m *discordgo.MessageCreate) error {
	// find cart code in message
	reg := regexp.MustCompile(`[0-9]+`)
	cartCode := reg.FindString(m.Content)

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(30*time.Second))
	defer cancel()

	if err := h.usecase.ResetCart(ctx, &models.ResetCartRequest{
		CartCode: cartCode,
	}); err != nil {
		logrus.Errorf("Failed to reset cart: %v", err)
		return err
	}
	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Cart [%s] has been reset!", cartCode))
	return nil
}
