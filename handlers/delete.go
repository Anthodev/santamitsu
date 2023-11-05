package handlers

import (
	"anthodev/santamitsu/db"
	"anthodev/santamitsu/service"
	"anthodev/santamitsu/utils/response"
	"github.com/bwmarrin/discordgo"
)

func deleteHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	service.IsAdmin(s, i)

	db.DeleteOneSantaSecret(i.GuildID)

	response.SendInteractionResponse(s, i, "Your secret santa has been deleted", true)
}
