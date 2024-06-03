package handler

import (
	initServices "butler/application/domains/services/init"
	"butler/application/domains/warehouse/models"
	"butler/application/domains/warehouse/usecase"
	"context"
	"fmt"
	"strings"
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

func (h Handler) ShowWarehouse(s *discordgo.Session, m *discordgo.MessageCreate) error {
	warehouseName := strings.ReplaceAll(m.Content, "!showwarehouse ", "")
	logrus.Infof("show warehouse: %s", warehouseName)
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(30*time.Second))
	defer cancel()

	if err := h.usecase.ShowWarehouse(ctx, &models.ShowWarehouseRequest{
		WarehouseName: warehouseName,
	}); err != nil {
		logrus.Errorf("Failed to show warehouse: %v", err)
		return err
	}
	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Warehouse [%s] is ready!", warehouseName))
	return nil
}

func (h Handler) ResetShowWarehouse(s *discordgo.Session, m *discordgo.MessageCreate) error {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(30*time.Second))
	defer cancel()

	if err := h.usecase.ResetShowWarehouse(ctx); err != nil {
		logrus.Errorf("Failed to reset show warehouse: %v", err)
		return err
	}
	s.ChannelMessageSend(m.ChannelID, "reset show warehouse success!")
	return nil
}
