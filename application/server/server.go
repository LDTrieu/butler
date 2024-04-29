package server

import (
	commandHandler "butler/application/commands/handler"
	initPromtAiHandler "butler/application/domains/promt_ai/makersuite/handler"
	initServices "butler/application/domains/services/init"
	"context"

	"butler/config"
	"butler/pkg/sql/mysql"

	"github.com/bwmarrin/discordgo"
	"github.com/google/generative-ai-go/genai"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"google.golang.org/api/option"
)

type Server struct {
	cfg *config.Config
	// cronJob     *cron.Cron
	// redisCli    *redis.Client
	discordBot *discordgo.Session
	db         *gorm.DB
	// grpcSv      *grpc.Server
	genaiClient *genai.Client
}

func NewServer(cfg *config.Config) *Server {
	// discord
	dg, err := discordgo.New("Bot " + cfg.DiscordBot.Butler.Token)
	if err != nil {
		logrus.Fatalf("init discord bot err: %v", err)
	}
	dg.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	db, err := mysql.InitConnection(cfg)
	if err != nil {
		logrus.Fatalf("connect database err: %v", err)
	}

	// genai client
	genaiClient, err := genai.NewClient(context.Background(), option.WithAPIKey(cfg.Makersuite.ApiKey))
	if err != nil {
		logrus.Fatalf("init genai client err: %s", err)
	}

	return &Server{
		cfg:        cfg,
		discordBot: dg,
		// redisCli:    redisCli,
		// cronJob:     crn,
		db: db,
		// grpcSv:      rpcServer,
		genaiClient: genaiClient,
	}
}

func (s *Server) Start() {
	s.run()
}

func (s *Server) Stop() {
	s.discordBot.Close()
	// s.redisCli.Close()
	// s.cronJob.Stop()
	s.genaiClient.Close()
	// s.grpcSv.Stop()
}

func (s *Server) run() {
	err := s.discordBot.Open()
	if err != nil {
		logrus.Fatalf("opening connection discord err: %v", err)
		return
	}
	// s.cronJob.Start()

	// init services
	services := initServices.InitService(s.cfg, s.genaiClient)

	// init external
	promtAiHandler := initPromtAiHandler.InitHandler(s.cfg, services)

	// register handler
	commandHandler := commandHandler.NewCommandHandler(s.discordBot, promtAiHandler)
	s.discordBot.AddHandler(commandHandler.GetCommandsHandler)

	logrus.Infof("start server success")
}
