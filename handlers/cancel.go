package handlers

import (
	"anthodev/santamitsu/db"
	"anthodev/santamitsu/utils/response"
	"github.com/bwmarrin/discordgo"
)

func CancelHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	db.DeleteOneSantaSecret(i.GuildID)

	response.SendInteractionResponse(s, i, "Your secret santa has been cancelled", true)
}
