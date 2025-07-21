package handler

import (
	"butler/application/domains/pick/models"
	"butler/application/domains/pick/usecase"
	initServices "butler/application/domains/services/init"
	"butler/application/lib"
	"butler/constants"
	"context"
	"fmt"
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

func (h Handler) ReadyPickOutbound(s *discordgo.Session, m *discordgo.MessageCreate) error {
	parts := strings.Fields(m.Content)
	if len(parts) < 2 {
		return fmt.Errorf("Invalid command format. \nExample: !readypick <outbound_order_number_1>, <outbound_order_number_2>")
	}

	orderInput := strings.Join(parts[1:], " ")
	salesOrderNumbers := []string{}
	for _, order := range strings.Split(orderInput, ",") {
		order = strings.TrimSpace(order)
		if order != "" {
			salesOrderNumbers = append(salesOrderNumbers, order)
		}
	}

	if len(salesOrderNumbers) == 0 {
		return fmt.Errorf("Not found sales_order_number")
	}

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(60*time.Second))
	defer cancel()

	if len(salesOrderNumbers) == 1 {
		saleOrderNumber := salesOrderNumbers[0]
		logrus.Infof("Prepare for outbound [%s] to ready pick", saleOrderNumber)

		if err := h.usecase.ReadyPickOutbound(ctx, &models.ReadyPickOutboundRequest{
			SalesOrderNumber: saleOrderNumber,
		}); err != nil {
			logrus.Errorf("Failed ready outbound: %v", err)
			return err
		}

		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Đơn hàng [%s] đã sẵn sàng để pick!", saleOrderNumber))
		return nil
	}

	var successOrders []string
	var failedOrders []struct {
		OrderNumber string
		Error       string
	}

	for _, saleOrderNumber := range salesOrderNumbers {
		logrus.Infof("Prepare for outbound [%s] to ready pick", saleOrderNumber)

		err := h.usecase.ReadyPickOutbound(ctx, &models.ReadyPickOutboundRequest{
			SalesOrderNumber: saleOrderNumber,
		})

		if err != nil {
			logrus.Errorf("Failed ready outbound [%s]: %v", saleOrderNumber, err)
			failedOrders = append(failedOrders, struct {
				OrderNumber string
				Error       string
			}{
				OrderNumber: saleOrderNumber,
				Error:       err.Error(),
			})
		} else {
			successOrders = append(successOrders, saleOrderNumber)
		}
	}

	var message string
	if len(successOrders) > 0 {
		message += fmt.Sprintf("Đơn sẵn sàng để pick: %s\n", strings.Join(successOrders, ", "))
	}

	if len(failedOrders) > 0 {
		message += "Đơn xử lý thất bại:\n"
		for _, failed := range failedOrders {
			message += fmt.Sprintf("- %s\n", failed.Error)
		}
	}

	s.ChannelMessageSend(m.ChannelID, message)
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
