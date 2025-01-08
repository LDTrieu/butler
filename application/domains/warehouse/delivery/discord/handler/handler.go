package handler

import (
	initServices "butler/application/domains/services/init"
	"butler/application/domains/warehouse/models"
	"butler/application/domains/warehouse/usecase"
	"butler/application/lib"
	"butler/constants"
	"context"
	"fmt"
	"regexp"
	"strconv"
	"strings"
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

func (h Handler) ShowConfigWarehouse(s *discordgo.Session, m *discordgo.MessageCreate) error {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(30*time.Second))
	defer cancel()

	reg := regexp.MustCompile(`[0-9]+`)
	warehouseId := reg.FindString(m.Content)
	warehouseIdInt, err := strconv.ParseInt(warehouseId, 10, 64)
	if err != nil {
		logrus.Errorf("Failed to parse warehouse id: %v", err)
		return err
	}

	wh, err := h.usecase.GetWarehouseById(ctx, warehouseIdInt)
	if err != nil {
		logrus.Errorf("Failed to get warehouse by id: %v", err)
		return err
	}

	operation := "add"
	if strings.Contains(m.Content, "sub") {
		operation = "sub"
	}

	description := "React vào emoji để thêm config"
	if operation == "sub" {
		description = "React vào emoji để bỏ config"
	}

	response, err := s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:       wh.WarehouseName,
		Description: description,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:  constants.EMOJI_CHAR_A + ": config location",
				Value: "\n",
			},
			{
				Name:  constants.EMOJI_CHAR_B + ": config abcxyz",
				Value: "\n",
			},
			{
				Name:  constants.EMOJI_CHAR_C + ": TODO: thêm config",
				Value: "\n",
			},
		},
	})
	if err != nil {
		logrus.Errorf("Error handle help command %v", err)
		return err
	}
	s.MessageReactionAdd(m.ChannelID, response.ID, constants.EMOJI_CHAR_A)
	s.MessageReactionAdd(m.ChannelID, response.ID, constants.EMOJI_CHAR_B)
	s.MessageReactionAdd(m.ChannelID, response.ID, constants.EMOJI_CHAR_C)

	h.lib.Cache.Set(response.ID+"::"+constants.CACHE_KEY_WH_CONFIG, &models.UpdateConfigWarehouseRequest{
		WarehouseId: wh.WarehouseId,
		Operation:   operation,
	})

	return nil
}

func (h Handler) UpdateConfigWarehouse(ctx context.Context, request *models.UpdateConfigWarehouseRequest) error {
	switch request.Operation {
	case "sub":
		return h.usecase.RemoveConfigWarehouse(ctx, request)
	case "add":
		return h.usecase.AddConfigWarehouse(ctx, request)
	}

	return nil
}
