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
				Name: "chuẩn bị cho đơn outbound có thể được đi pick",
				Value: `
					!readypick <mã source number>
					Eample: !readypick 100224050700001
				`,
			},
			{
				Name: "cho kho xuất hiện để đi pick ở vị trí kho 29 HOANG VIET",
				Value: `
					!showwarehouse <tên kho>
					Eample: !showwarehouse SHOP - 29
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
