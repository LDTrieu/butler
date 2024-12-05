package handler

import (
	"butler/application/domains/cart/models"
	"butler/application/domains/cart/usecase"
	initServices "butler/application/domains/services/init"
	"context"
	"fmt"
	"regexp"
	"strconv"
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
	logrus.Infof("Reset cart code: %s", cartCode)
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

func (h Handler) ResetCartByUser(s *discordgo.Session, m *discordgo.MessageCreate) error {
	// find user_id in message
	reg := regexp.MustCompile(`[0-9]+`)
	userIdStr := reg.FindString(m.Content)
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		logrus.Errorf("Failed to parse user ID: %v", err)
		return err
	}

	logrus.Infof("Reset cart mapping : %s", userIdStr)
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(30*time.Second))
	defer cancel()

	cartCode, err := h.usecase.ResetCartByUser(ctx, &models.ResetCartByUserRequest{
		UserId: userId,
	})
	if err != nil {
		logrus.Errorf("Failed to reset cart mapping: %v", err)
		return err
	}
	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Cart [%s] - User [%s] has been reset!", cartCode, userIdStr))
	return nil
}
