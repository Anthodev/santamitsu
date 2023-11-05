package handlers

import (
	"anthodev/santamitsu/db"
	"anthodev/santamitsu/model"
	"anthodev/santamitsu/utils/response"
	"github.com/bwmarrin/discordgo"
)

func lockHandler(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
	ss model.SecretSanta,
) {
	ss.State = 0

	db.UpdateSantaSecret(ss)

	response.SendInteractionResponse(s, i, "Your secret santa has been locked", true)
}

func unlockHandler(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
	ss model.SecretSanta,
) {
	ss.State = 1

	db.UpdateSantaSecret(ss)

	response.SendInteractionResponse(s, i, "Your secret santa has been unlocked", true)
}
