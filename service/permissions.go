package service

import (
	"anthodev/santamitsu/model"
	"anthodev/santamitsu/utils/response"
	"github.com/bwmarrin/discordgo"
)

func isMemberAuthorizedCheck(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
	ss model.SecretSanta,
) bool {
	m, _ := s.GuildMember(i.GuildID, i.Member.User.ID)

	for _, mr := range ss.ModeratorRoles {
		for _, r := range m.Roles {
			if mr.RoleId == r {
				return true
			}
		}
	}

	g, _ := s.Guild(i.GuildID)

	if m.User.ID == g.OwnerID {
		return true
	}

	return false
}

func isAdminCheck(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
) bool {
	m, _ := s.GuildMember(i.GuildID, i.Member.User.ID)
	g, _ := s.Guild(i.GuildID)

	if m.User.ID == g.OwnerID {
		return true
	}

	return false
}

func IsMemberAuthorized(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
	ss model.SecretSanta,
) {
	if isMemberAuthorizedCheck(s, i, ss) == false {
		response.SendInteractionResponse(s, i, "You are not authorized to do this.", true)
		return
	}
}

func IsAdmin(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
) {
	if isAdminCheck(s, i) == false {
		response.SendInteractionResponse(s, i, "You are not authorized to do this.", true)
		return
	}
}
