package handlers

import (
	"anthodev/santamitsu/db"
	"anthodev/santamitsu/model"
	"anthodev/santamitsu/service"
	"anthodev/santamitsu/utils/component"
	"anthodev/santamitsu/utils/helpers"
	"anthodev/santamitsu/utils/response"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"math/rand"
)

func drawHandler(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
	ss model.SecretSanta,
) {
	service.IsMemberAuthorized(s, i, ss)

	if ss.State != 0 {
		response.SendInteractionResponse(s, i, "You can't draw the Secret Santa if it's not locked!", true)

		return
	}

	if len(ss.Participants) < 2 {
		response.SendInteractionResponse(s, i, "You need at least 3 participants to draw the Secret Santa!", true)

		return
	}

	var matchedUsers []string

	for _, p := range ss.Participants {
		matchingList := getMatchingList(p.UserId, matchedUsers, ss)

		if len(matchingList) == 0 {
			response.SendInteractionResponse(s, i, "There is no matching for the Secret Santa!", true)

			return
		}

		randomIndex := rand.Intn(len(matchingList))
		randomParticipant := matchingList[randomIndex]
		matchedPair := model.MatchedPair{
			UserId1: p.UserId,
			UserId2: randomParticipant.UserId,
		}
		ss = model.AddMatchedPair(ss, matchedPair)

		matchedUsers = append(matchedUsers, randomParticipant.UserId)
	}

	db.UpdateSantaSecret(ss)

	announceDraw(s, i, ss)
}

func getMatchingList(
	uid string,
	matchedUsers []string,
	ss model.SecretSanta,
) []model.SantaParticipant {
	var matchingList []model.SantaParticipant
	var exclusionList []string

	for _, e := range ss.ExcludedPairs {
		if e.UserId1 == uid {
			exclusionList = append(exclusionList, e.UserId2)
		}
	}

	for _, p := range ss.Participants {
		if p.UserId == uid {
			continue
		}

		if helpers.Contains(exclusionList, p.UserId) {
			continue
		}

		if helpers.Contains(matchedUsers, p.UserId) {
			continue
		}

		matchingList = append(matchingList, p)
	}

	return matchingList
}

func announceDraw(s *discordgo.Session, i *discordgo.InteractionCreate, ss model.SecretSanta) {
	embed := component.NewGenericEmbed(
		fmt.Sprintf("Result of the Secret Santa \"%s\"", ss.Title),
		"The draw has been made, your match will be announced personally in the next message!",
	)

	response.SendInteractionEmbedResponse(s, i, embed, false)

	for _, p := range ss.Participants {
		for _, m := range ss.MatchedPairs {
			if m.UserId1 == p.UserId {
				announceDrawMember(s, i, m, ss)
				break
			}
		}
	}
}

func announceDrawMember(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
	mp model.MatchedPair,
	ss model.SecretSanta,
) {
	for _, p := range ss.Participants {
		var embed *discordgo.MessageEmbed

		if p.UserId == mp.UserId1 {
			m, _ := s.GuildMember(i.GuildID, mp.UserId1)
			embed = buildAnnounceDrawMemberMessage(m, mp, ss)
			break
		}

		response.SendInteractionEmbedResponse(s, i, embed, true)
	}
}

func buildAnnounceDrawMemberMessage(
	m *discordgo.Member,
	mp model.MatchedPair,
	ss model.SecretSanta,
) *discordgo.MessageEmbed {
	embed := component.NewGenericEmbed(
		fmt.Sprintf("Result of the Secret Santa \"%s\"", ss.Title),
		"You must offer a gift to: **"+mp.UserId2+"**",
	)

	embed.Description += "\n\n"
	embed.Description += fmt.Sprintf("Remember taht the maximum price for the gift is: **%s%s**\n\n", ss.MaxPrice, ss.Currency)
	embed.Description += "Feel free to contact your match to know how to offer your gift!\n\n"
	embed.Description += "If you have any question, please contact a moderator or an admin.\n\n"
	embed.Description += fmt.Sprintf("Merry Christmas %s! 🎄\n\n", m.Mention())

	return embed
}
