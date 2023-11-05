package service

import (
	"anthodev/santamitsu/model"
	"anthodev/santamitsu/utils/response"
	"github.com/bwmarrin/discordgo"
)

func isLocked(
	ss model.SecretSanta,
) bool {
	if ss.State == 0 {
		return true
	}

	return false
}

func LockState(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
	ss model.SecretSanta,
) {
	if isLocked(ss) {
		response.SendInteractionResponse(s, i, "You can't perform this action, the Secret Santa is locked.", true)
		return
	}
}
