package handlers

import (
	"anthodev/santamitsu/db"
	"anthodev/santamitsu/service"
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

			infoHandler(s, i, ss)
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

			announceHandler(s, i, ss)
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

			service.IsMemberAuthorized(s, i, ss)

			cancelHandler(s, i)
		},
		"join": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			if i.Type != discordgo.InteractionApplicationCommand {
				return
			}

			ss := db.FindOneSantaSecret(i.GuildID)

			if ss.Title == "" {
				response.SendInteractionResponse(s, i, "No Secret Santa is active!", true)

				return
			}

			joinHandler(s, i, ss, false)
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

			leaveHandler(s, i, ss, false)
		},
		"exclude": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			ss := db.FindOneSantaSecret(i.GuildID)

			if ss.Title == "" {
				response.SendInteractionResponse(s, i, "No Secret Santa is active!", true)

				return
			}

			excludeHandler(s, i, ss)
		},
		"moderator-role": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			ss := db.FindOneSantaSecret(i.GuildID)

			if ss.Title == "" {
				response.SendInteractionResponse(s, i, "No Secret Santa is active!", true)

				return
			}

			moderatorRoleHandler(s, i, ss)
		},
		"lock": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			ss := db.FindOneSantaSecret(i.GuildID)

			if ss.Title == "" {
				response.SendInteractionResponse(s, i, "No Secret Santa is active!", true)

				return
			}

			lockHandler(s, i, ss)
		},
		"unlock": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			ss := db.FindOneSantaSecret(i.GuildID)

			if ss.Title == "" {
				response.SendInteractionResponse(s, i, "No Secret Santa is active!", true)

				return
			}

			unlockHandler(s, i, ss)
		},
		"draw": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			ss := db.FindOneSantaSecret(i.GuildID)

			if ss.Title == "" {
				response.SendInteractionResponse(s, i, "No Secret Santa is active!", true)

				return
			}

			drawHandler(s, i, ss)
		},
		"get-pair": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			ss := db.FindOneSantaSecret(i.GuildID)

			if ss.Title == "" {
				response.SendInteractionResponse(s, i, "No Secret Santa is active!", true)

				return
			}

			getPairHandler(s, i, ss)
		},
		"delete": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			ss := db.FindOneSantaSecret(i.GuildID)

			if ss.Title != "" {
				response.SendInteractionResponse(s, i, "You already have a secret santa running!", true)

				return
			}

			deleteHandler(s, i)
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

			joinHandler(s, i, ss, true)
		},
		"leave": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			if i.Type != discordgo.InteractionMessageComponent {
				return
			}

			ss := db.FindOneSantaSecret(i.GuildID)

			if ss.Title == "" {
				response.SendInteractionResponse(s, i, "No Secret Santa is active!", true)

				return
			}

			leaveHandler(s, i, ss, true)
		},
	}
}
