package discord

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

const TIME_FORMAT = "02-01-2006 15:04:05"

func CreateEmbed() *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Title:     "Embed title",
		URL:       "https://www.google.com/",
		Timestamp: time.Now().Format(time.RFC3339),
		Image: &discordgo.MessageEmbedImage{
			URL: "https://avatars.githubusercontent.com/u/62052787?v=4",
		},
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://avatars.githubusercontent.com/u/127912210?v=4",
		},
		Color:       0x00ff00,
		Description: "embed description",
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "field A",
				Value:  "value A",
				Inline: true,
			},
		},
	}
}

