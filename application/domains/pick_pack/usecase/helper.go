package usecase

import (
	"butler/application/domains/pick_pack/models"
	"butler/constants"
	"context"
	"encoding/json"
	"errors"

	"bitbucket.org/hasaki-tech/zeus/package/hrequest"
)

const NO_RETRY = 0
const NO_DELAY = 0
const TIMEOUT_20S = 20

func apiRetryCondition(status int, body []byte, errRequest error) bool {
	return false
}

func (u *usecase) loginWms(ctx context.Context, email, password string) (*models.LoginWmsResponse, error) {
	headers := map[string]string{
		"Content-Type": "application/json",
	}

	endpoint := u.cfg.ApiExternal.Wms.Url
	url := endpoint + constants.WMS_LOGIN

	_, body, err := hrequest.MakeRequest(headers, constants.METHOD_POST, url, &models.LoginWmsRequest{
		EmailWms:    email,
		PasswordWms: password,
	}, 10, NO_RETRY, NO_DELAY, apiRetryCondition)
	if err != nil {
		return nil, err
	}
	response := &models.LoginWmsResponse{}
	if err := json.Unmarshal(body, response); err != nil {
		return nil, err
	}
	if response.Message != "" {
		return nil, errors.New(response.Message)
	}
	return response, nil
}

func (u *usecase) loginDiscord(ctx context.Context, email, password string) (*models.LoginDiscordResponse, error) {
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	endpoint := u.cfg.ApiExternal.Discord.Url
	url := endpoint + constants.DISCORD_LOGIN

	_, body, err := hrequest.MakeRequest(headers, constants.METHOD_POST, url, &models.LoginDiscordRequest{
		LoginDiscord:    email,
		PasswordDiscord: password,
		Undelete:        false,
	}, 10, NO_RETRY, NO_DELAY, apiRetryCondition)
	if err != nil {
		return nil, err
	}
	response := &models.LoginDiscordResponse{}
	if err := json.Unmarshal(body, response); err != nil {
		return nil, err
	}
	if response.Message != "" {
		return nil, errors.New(response.Message)
	}

	return response, nil

}
