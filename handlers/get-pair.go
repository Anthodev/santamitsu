package handlers

import (
	"anthodev/santamitsu/model"
	"anthodev/santamitsu/utils/response"
	"github.com/bwmarrin/discordgo"
)

func getPairHandler(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
	ss model.SecretSanta,
) {
	for _, m := range ss.MatchedPairs {
		if m.UserId1 == i.Member.User.ID {
			announceDrawMember(s, i, m, ss)
			return
		}
	}

	response.SendInteractionResponse(s, i, "You have not been assigned a match yet!", true)
}
