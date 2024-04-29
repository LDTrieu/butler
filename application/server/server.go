package server

import (
	commandsDelivery "butler/application/commands/delivery"
	initMakersuiteHandler "butler/application/domains/external/makersuite/handler"
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
	dg, err := discordgo.New("Bot " + cfg.DiscordBot.butler.Token)
	if err != nil {
		logrus.Fatalf("init discord bot err: %v", err)
	}
	dg.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	// redis
	// redisCli := redis.NewClient(&redis.Options{
	// 	Addr:     cfg.Redis.Addr,
	// 	Password: cfg.Redis.Password,
	// 	DB:       cfg.Redis.Db,
	// })

	// cron job
	// crn := cron.New()

	// sql
	// db, _, err := postgresql.InitConnection(cfg)
	// if err != nil {
	// 	logrus.Fatalf("connect database err: %v", err)
	// }
	db, err := mysql.InitConnection(cfg)
	if err != nil {
		logrus.Fatalf("connect database err: %v", err)
	}

	// grpc
	// rpcServer := grpc.NewServer(grpc.MaxRecvMsgSize(10 * constants.MB))

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
	makersuiteHandler := initMakersuiteHandler.InitHandler(s.cfg, makersuiteSv)

	// init domains
	acc := initAccount.NewInit(s.cfg, s.db)
	attendance := initAttendance.NewInit(s.cfg, s.db)

	// register handler
	commandHandler := commandsDelivery.NewCommandsDelivery(s.discordBot, acc.DiscordHandler, attendance.Handler, makersuiteHandler)
	s.discordBot.AddHandler(commandHandler.GetCommandsHandler)

	logrus.Infof("start server success")
}
