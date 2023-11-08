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

	embed := component.NewGenericEmbed(
		fmt.Sprintf("\"**%s**\" - Budget modification", ss.Title),
		"Check your DMs to set the budget for your secret santa!",
	)

	response.SendInteractionEmbedResponse(s, i, embed, true)

	uc := response.CreateDmChannel(s, i.Member.User.ID)

	embed = component.NewGenericEmbed(
		fmt.Sprintf("\"**%s**\" - Budget modification", ss.Title),
		"What is the new maximum budget per participant?",
	)

	response.SendDmEmbed(s, uc.ID, embed)

	s.AddHandlerOnce(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		contentMsgHandler(s, i, m)

		if m.Content != "" {
			content := m.Content

			ss.MaxPrice = content

			ss = db.UpdateSantaSecret(ss)

			embed := component.NewGenericEmbed(
				fmt.Sprintf("\"**%s**\" - Budget modification", ss.Title),
				fmt.Sprintf("The budget has been updated to **%s%s**", ss.MaxPrice, ss.Currency),
			)

			response.SendDmEmbed(s, uc.ID, embed)
		}
	})
}
