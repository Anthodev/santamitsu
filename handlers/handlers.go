package handlers

import (
	"anthodev/santamitsu/db"
	"anthodev/santamitsu/utils/response"
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
		"info": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			if i.Type != discordgo.InteractionApplicationCommand {
				return
			}

			ss := db.FindOneSantaSecret(i.GuildID)

			if ss.Title == "" {
				response.SendInteractionResponse(s, i, "You don't have any secret santa running!", true)

				return
			}

			InfoHandler(s, i, ss)
		},
		"announce": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			if i.Type != discordgo.InteractionApplicationCommand {
				return
			}

			ss := db.FindOneSantaSecret(i.GuildID)

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
		"join": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			if i.Type != discordgo.InteractionMessageComponent {
				return
			}

			ss := db.FindOneSantaSecret(i.GuildID)

			if ss.Title == "" {
				response.SendInteractionResponse(s, i, "No Secret Santa is active!", true)

				return
			}

			JoinHandler(s, i, ss, false)
		},
		"leave": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			if i.Type != discordgo.InteractionApplicationCommand {
				return
			}

			ss := db.FindOneSantaSecret(i.GuildID)

			if ss.Title == "" {
				response.SendInteractionResponse(s, i, "No Secret Santa is active!", true)

				return
			}

			LeaveHandler(s, i, ss)
		},
	}
}

func Buttons() map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"join": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			if i.Type != discordgo.InteractionMessageComponent {
				return
			}

			ss := db.FindOneSantaSecret(i.GuildID)

			if ss.Title == "" {
				response.SendInteractionResponse(s, i, "No Secret Santa is active!", true)

				return
			}

			JoinHandler(s, i, ss, true)
		},
	}
}
