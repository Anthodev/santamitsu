package handlers

import (
	"anthodev/santamitsu/db"
	"anthodev/santamitsu/model"
	"anthodev/santamitsu/utils/response"
	"github.com/bwmarrin/discordgo"
)

func LockHandler(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
	ss model.SantaSecret,
) {
	ss.State = 0

	db.UpdateSantaSecret(ss)

	response.SendInteractionResponse(s, i, "Your secret santa has been locked", true)
}

func UnlockHandler(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
	ss model.SantaSecret,
) {
	ss.State = 1

	db.UpdateSantaSecret(ss)

	response.SendInteractionResponse(s, i, "Your secret santa has been unlocked", true)
}
