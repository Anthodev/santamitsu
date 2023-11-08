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

	uc := response.CreateDmChannel(s, i.Member.User.ID)

	embed := component.NewGenericEmbed(
		fmt.Sprintf("\"**%s**\" modification - Budget", ss.Title),
		"What is the new maximum budget per participant?",
	)

	response.SendDmEmbed(s, uc.ID, embed)

	s.AddHandlerOnce(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		contentMsgHandler(s, i, m)

		if m.Content != "" {
			content := m.Content

			ss.MaxPrice = content

			db.UpdateSantaSecret(ss)

			embed := component.NewGenericEmbed(
				fmt.Sprintf("\"**%s**\" modification - Budget", ss.Title),
				fmt.Sprintf("The budget has been updated to **%s%s**", ss.MaxPrice, ss.Currency),
			)

			response.SendInteractionEmbedResponse(s, i, embed, true)
		}
	})
}
