package handlers

import (
	"anthodev/santamitsu/db"
	"anthodev/santamitsu/model"
	"anthodev/santamitsu/service"
	"anthodev/santamitsu/utils/component"
	"anthodev/santamitsu/utils/response"
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func setBudgetHandler(s *discordgo.Session, i *discordgo.InteractionCreate, ss model.SecretSanta) {
	service.IsMemberAuthorized(s, i, ss)

	switch i.Type {
	case discordgo.InteractionApplicationCommand:
		data := i.ApplicationCommandData()
		subcmd := data.Options[0]

		if subcmd.Name == "maximum-price" {
			setBudget(s, i, ss)
			return
		}
	}
}

func setBudget(s *discordgo.Session, i *discordgo.InteractionCreate, ss model.SecretSanta) {
	nmp := i.ApplicationCommandData().Options[0].StringValue()
	
	ss.MaxPrice = nmp

	ss = db.UpdateSantaSecret(ss)

	embed := component.NewGenericEmbed(
		fmt.Sprintf("\"**%s**\" - Budget modification", ss.Title),
		fmt.Sprintf("The budget has been updated to **%s%s**", nmp, ss.Currency),
	)

	response.SendInteractionEmbedResponse(s, i, embed, true)
}
