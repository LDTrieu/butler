package handler

import (
	"butler/application/domains/pick_pack/models"
	"butler/application/domains/pick_pack/usecase"
	initServices "butler/application/domains/services/init"
	"butler/application/lib"
	"butler/config"
	"context"
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	lib     *lib.Lib
	usecase usecase.IUseCase
}

func InitHandler(lib *lib.Lib, cfg *config.Config, services *initServices.Services) Handler {
	usecase := usecase.InitUseCase(lib, cfg, services)
	return Handler{
		lib:     lib,
		usecase: usecase,
	}
}

func (h Handler) ReadyPickPack(s *discordgo.Session, m *discordgo.MessageCreate) error {
	// Dùng regex để bắt các số trong dấu []
	re := regexp.MustCompile(`\[(\d+)\]`)
	matches := re.FindAllStringSubmatch(m.Content, -1)

	if len(matches) < 2 {
		return fmt.Errorf("Thiếu tham số. Vui lòng sử dụng: !runpickpack [mã đơn][mã vận chuyển]")
	}

	// matches[i][1] chứa số trong dấu [] (group đầu tiên của regex)
	shipmentNumber := matches[0][1]
	shippingUnitId, err := strconv.ParseInt(matches[1][1], 10, 64)
	if err != nil {
		return fmt.Errorf("Mã vận chuyển không hợp lệ: %v", err)
	}

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(30*time.Second))
	defer cancel()
	requestParams := &models.AutoPickPackRequest{
		LoginRequest: models.LoginRequest{
			LoginDiscordRequest: models.LoginDiscordRequest{
				Login:    "sonplh1@hasaki.vn",
				Password: "12345a@A",
				Undelete: false,
			},
		},
		SalesOrderNumber: shipmentNumber,
		ShippingUnitId:   shippingUnitId,
	}

	result, err := h.usecase.AutoPickPack(ctx, *requestParams)
	if err != nil {
		logrus.Errorf("Failed ready pickpack: %v", err)
		return err
	}

	// Chia nhỏ kết quả thành nhiều phần nếu quá dài
	const maxLength = 1990 // Để lại một chút dư cho các ký tự định dạng
	for i := 0; i < len(result); i += maxLength {
		end := i + maxLength
		if end > len(result) {
			end = len(result)
		}
		chunk := result[i:end]
		_, err := s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("```%s```", chunk))
		if err != nil {
			logrus.Errorf("Failed to send message: %v", err)
			return err
		}
	}
	return nil
}
