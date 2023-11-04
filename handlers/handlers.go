package handlers

import (
	"anthodev/santamitsu/db"
	"anthodev/santamitsu/utils/response"
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func List() map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"setup": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			if i.Type != discordgo.InteractionApplicationCommand {
				return
			}

			ss := db.FindOneSantaSecret(i.GuildID)

			if ss.Title != "" {
				response.SendInteractionResponse(s, i, "You already have a secret santa running!", true)

				return
			}

			response.SendInteractionResponse(s, i, "Check your DMs to setup the secret santa!", true)

			setupHandler(s, i)
		},
		"announce": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			if i.Type != discordgo.InteractionApplicationCommand {
				return
			}

			fmt.Println(i.GuildID)

			ss := db.FindOneSantaSecret(i.GuildID)

			fmt.Println(ss)

			if ss.Title == "" {
				response.SendInteractionResponse(s, i, "You don't have any secret santa running!", true)

				return
			}

			AnnounceHandler(s, i, ss)
		},
		"cancel": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			if i.Type != discordgo.InteractionApplicationCommand {
				return
			}

			ss := db.FindOneSantaSecret(i.GuildID)

			if ss.Title == "" {
				response.SendInteractionResponse(s, i, "You don't have any secret santa running!", true)

				return
			}

			CancelHandler(s, i)
		},
	}
}
