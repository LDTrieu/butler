package helper

import (
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

func HandleHelpCommand(s *discordgo.Session, m *discordgo.MessageCreate) error {
	logrus.Infof("User [%s] needs help", m.Author.Username)
	_, err := s.ChannelMessageSendEmbed(m.ChannelID, HelpEmbed())
	if err != nil {
		logrus.Errorf("Error handle help command %v", err)
		return err
	}
	return nil
}

func HelpEmbed() *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Title: "Command List",
		Description: `
			Here is the list of commands!
			`,
		Color: 0x00ff00,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name: "Ask me anything",
				Value: `
				@butler <your request>
				Example: @butler write hello world app in go
				`,
			},
			{
				Name: "Chuyển trạng thái giỏ về available",
				Value: `
					!resetcart <mã giỏ>
					Example: !resetcart 160143
				`,
			},
			{
				Name: "pick liền đơn outbound",
				Value: `
					đang phát triển
				`,
			},
		},
	}
}

func HandleSendImage(s *discordgo.Session, m *discordgo.MessageCreate) error {
	f, err := os.Open("some_img.jpg") // relative path to the main.go file
	if err != nil {
		logrus.Errorf("Error open image %v", err)
		return err
	}
	_, err = s.ChannelFileSend(m.ChannelID, "qwe.jpg", f)
	if err != nil {
		logrus.Errorf("Error handle send image %v", err)
		return err
	}
	return nil
}
