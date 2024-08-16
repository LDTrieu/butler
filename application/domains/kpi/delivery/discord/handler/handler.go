package handler

import (
	"butler/application/domains/kpi/usecase"
	initServices "butler/application/domains/services/init"
	"butler/application/lib"
	"context"
	"fmt"
	"time"

	"strings"

	"github.com/bwmarrin/discordgo"
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

func (h Handler) CountKpi(s *discordgo.Session, m *discordgo.MessageCreate) error {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(30*time.Second))
	defer cancel()
	if strings.Contains(m.Content, "!prod kpi ") {
		if !isUserHavePermission(s, m) {
			return fmt.Errorf("you don't have permission to do this action")
		}

		date := strings.ReplaceAll(m.Content, "!prod kpi ", "")
		return h.usecase.CountKpi(ctx, date, "prod")
	}
	if strings.Contains(m.Content, "!kpi ") {
		date := strings.ReplaceAll(m.Content, "!kpi ", "")
		return h.usecase.CountKpi(ctx, date, "qc")
	}

	return nil
}

func isUserHavePermission(s *discordgo.Session, m *discordgo.MessageCreate) bool {
	return false
}
