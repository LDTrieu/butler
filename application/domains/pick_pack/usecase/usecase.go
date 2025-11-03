package usecase

import (
	"butler/application/domains/pick_pack/models"
	initServices "butler/application/domains/services/init"
	outboundOrderSv "butler/application/domains/services/outbound_order/service"
	outboundOrderExtendSv "butler/application/domains/services/outbound_order_extend/service"
	"butler/application/lib"
	"butler/config"
	"butler/constants"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"
)

type usecase struct {
	lib                   *lib.Lib
	cfg                   *config.Config
	outboundOrderSv       outboundOrderSv.IService
	outboundOrderExtendSv outboundOrderExtendSv.IService
}

func InitUseCase(
	lib *lib.Lib,
	cfg *config.Config,
	services *initServices.Services,
) IUseCase {
	return &usecase{
		lib:                   lib,
		cfg:                   cfg,
		outboundOrderSv:       services.OutboundOrderService,
		outboundOrderExtendSv: services.OutboundOrderExtendService,
	}
}

func (u *usecase) Login(ctx context.Context, params *models.LoginRequest) (*models.LoginResponse, error) {
	var wms *models.LoginWmsResponse
	var discord *models.LoginDiscordResponse

	wmsDataStr, err := u.lib.Rdb.Get(ctx, fmt.Sprintf("%s:%s", constants.WMS_DATA, params.LoginWmsRequest.EmailWms)).Result()
	if err != nil && err.Error() != "redis: nil" {
		return nil, err
	}

	var wmsData map[string]string
	if wmsDataStr != "" {
		if err := json.Unmarshal([]byte(wmsDataStr), &wmsData); err != nil {
			return nil, err
		}
	}

	if wmsData == nil || wmsData["token"] == "" {
		wms, err = u.loginWms(ctx, params.LoginWmsRequest.EmailWms, params.LoginWmsRequest.PasswordWms)
		if err != nil {
			return nil, err
		}

		wmsData = map[string]string{
			"token":  wms.Token,
			"email":  wms.User.Email,
			"userId": strconv.FormatInt(int64(wms.User.UserId), 10),
		}

		wmsDataBytes, err := json.Marshal(wmsData)
		if err != nil {
			return nil, err
		}

		if err := u.lib.Rdb.Set(ctx, fmt.Sprintf("%s:%s", constants.WMS_DATA, wms.User.Email),
			string(wmsDataBytes), 8*time.Hour).Err(); err != nil {
			return nil, err
		}
	} else {
		userId, _ := strconv.Atoi(wmsData["userId"])
		wms = &models.LoginWmsResponse{
			Token: wmsData["token"],
			User: struct {
				UserId   int    `json:"user_id"`
				LastName string `json:"last_name"`
				Email    string `json:"email"`
				Status   string `json:"status"`
			}{
				Email:  wmsData["email"],
				UserId: userId,
			},
		}
	}

	tokenDiscord, err := u.lib.Rdb.Get(ctx, constants.TOKEN_DISCORD).Result()
	if err != nil && err.Error() != "redis: nil" {
		return nil, err
	}

	if tokenDiscord == "" {
		discord, err = u.loginDiscord(ctx, params.LoginDiscordRequest.LoginDiscord, params.LoginDiscordRequest.PasswordDiscord)
		if err != nil {
			return nil, err
		}
		tokenDiscord = discord.Token

		if err := u.lib.Rdb.Set(ctx, constants.TOKEN_DISCORD, tokenDiscord, 24*time.Hour).Err(); err != nil {
			return nil, err
		}
	} else {
		discord = &models.LoginDiscordResponse{
			Token: tokenDiscord,
		}
	}

	if wms.Token == "" || discord.Token == "" {
		return nil, errors.New("token wms or token discord is empty")
	}

	return &models.LoginResponse{
		LoginWmsResponse:     *wms,
		LoginDiscordResponse: *discord,
	}, nil
}
