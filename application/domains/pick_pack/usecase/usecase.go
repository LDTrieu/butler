package usecase

import (
	"butler/application/domains/pick_pack/models"
	initServices "butler/application/domains/services/init"
	"butler/application/lib"
	"butler/config"
	"bytes"
	"context"
	"fmt"
	"log"
	"os/exec"
)

type usecase struct {
	lib *lib.Lib
	cfg *config.Config
}

func InitUseCase(
	lib *lib.Lib,
	cfg *config.Config,
	services *initServices.Services,
) IUseCase {
	return &usecase{
		lib: lib,

		cfg: cfg,
	}
}

func (u *usecase) AutoPickPack(ctx context.Context, params models.AutoPickPackRequest) (string, error) {
	// Login

	login, err := u.Login(ctx, &params.LoginRequest)
	if err != nil {
		return "", err
	}
	tokenDiscord := login.LoginDiscordResponse.Token

	// Run newman json
	result, err := u.runNewman(ctx, params.SalesOrderNumber, params.ShippingUnitId, tokenDiscord)
	if err != nil {
		return "", err
	}

	return result, nil
}

func (u *usecase) Login(ctx context.Context, params *models.LoginRequest) (*models.LoginResponse, error) {

	// Login discord
	discord, err := u.loginDiscord(ctx, params.LoginDiscordRequest.Login, params.LoginDiscordRequest.Password)
	if err != nil {
		return nil, err
	}

	return &models.LoginResponse{
		LoginDiscordResponse: *discord,
	}, nil
}

func (u *usecase) runNewman(ctx context.Context, shipmentNumber string, shippingUnitId int64, token string) (string, error) {
	cmd := exec.Command("newman",
		"run",
		"trieu1.json",
		//"--env-var", fmt.Sprintf("token=%s", token),
		"--env-var", fmt.Sprintf("shipment_number=%s", shipmentNumber),
		"--env-var", fmt.Sprintf("shipping_unit_id=%d", shippingUnitId),
	)

	// Set output
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	// Run newman
	err := cmd.Run()
	if err != nil {
		log.Printf("Lỗi stderr: %s", stderr.String())
		return "", fmt.Errorf("error running newman: %v\nStderr: %s", err, stderr.String())
	}
	a := stdout.String()
	b := stderr.String()
	// In ra kết quả chi tiết
	log.Printf("=== Kết quả Newman ===")
	log.Printf("Stdout: %s", a)
	log.Printf("Stderr: %s", b)
	log.Printf("====================")

	return a, nil
}
