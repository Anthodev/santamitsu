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

func excludeHandler(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
	ss model.SecretSanta,
) {
	service.IsMemberAuthorized(s, i, ss)

	switch i.Type {
	case discordgo.InteractionApplicationCommand:
		data := i.ApplicationCommandData()
		subcmd := data.Options[0]

		if subcmd.Name == "member" {
			excludeUser(s, i, ss, subcmd.Options[0].UserValue(s))
			return
		}

		if subcmd.Name == "pair" {
			excludePair(s, i, ss, subcmd.Options[0].UserValue(s), subcmd.Options[1].UserValue(s))
			return
		}

		if subcmd.Name == "remove" {
			removeExcludedMember(s, i, ss, subcmd.Options[0].UserValue(s))
			return
		}

		if subcmd.Name == "remove-pair" {
			removeExcludedPair(s, i, ss, subcmd.Options[0].UserValue(s), subcmd.Options[1].UserValue(s))
			return
		}

		if subcmd.Name == "list" {
			listExcludedMembers(s, i, ss)
			return
		}
	case discordgo.InteractionApplicationCommandAutocomplete:
		data := i.ApplicationCommandData()
		var choices []*discordgo.ApplicationCommandOptionChoice

		switch data.Options[0].Options[0].Name {
		case "member":
			if data.Options[0].Options[0].StringValue() != "" {
				choices = append(choices, &discordgo.ApplicationCommandOptionChoice{
					Name:  data.Options[0].Name,
					Value: data.Options[0].UserValue(s).Mention(),
				})
			}
		case "member-1":
			if data.Options[0].Options[0].StringValue() != "" {
				choices = append(choices, &discordgo.ApplicationCommandOptionChoice{
					Name:  data.Options[0].Name,
					Value: data.Options[0].UserValue(s).Mention(),
				})
			}
		case "member-2":
			if data.Options[0].Options[1].StringValue() != "" {
				choices = append(choices, &discordgo.ApplicationCommandOptionChoice{
					Name:  data.Options[1].Name,
					Value: data.Options[1].UserValue(s).Mention(),
				})
			}
		case "rm-member":
			if data.Options[0].Options[0].StringValue() != "" {
				choices = append(choices, &discordgo.ApplicationCommandOptionChoice{
					Name:  data.Options[0].Name,
					Value: data.Options[0].UserValue(s).Mention(),
				})
			}
		case "rm-member-1":
			if data.Options[0].Options[0].StringValue() != "" {
				choices = append(choices, &discordgo.ApplicationCommandOptionChoice{
					Name:  data.Options[0].Name,
					Value: data.Options[0].UserValue(s).Mention(),
				})
			}
		case "rm-member-2":
			if data.Options[0].Options[1].StringValue() != "" {
				choices = append(choices, &discordgo.ApplicationCommandOptionChoice{
					Name:  data.Options[1].Name,
					Value: data.Options[1].UserValue(s).Mention(),
				})
			}
		}

		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionApplicationCommandAutocompleteResult,
			Data: &discordgo.InteractionResponseData{
				Choices: choices,
			},
		})

		if err != nil {
			return
		}
	}
}

func excludeUser(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
	ss model.SecretSanta,
	u *discordgo.User,
) {
	for _, p := range ss.ExcludedMembers {
		if p.UserId == u.ID {
			response.SendInteractionResponse(s, i, fmt.Sprintf("The member %s is already excluded!", u.Mention()), true)
			return
		}
	}

	e := model.ExcludedMember{UserId: u.ID}

	ss = model.AddExcludedMember(ss, e)
	db.UpdateSantaSecret(ss)

	response.SendInteractionResponse(s, i, fmt.Sprintf("Member \"%s\" is excluded!", u.Mention()), true)
}

func excludePair(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
	ss model.SecretSanta,
	u1 *discordgo.User,
	u2 *discordgo.User,
) {
	for _, p := range ss.ExcludedPairs {
		if p.UserId1 == u1.ID && p.UserId2 == u2.ID {
			response.SendInteractionResponse(s, i, fmt.Sprintf("The pair %s and %s is already excluded!", u1.Mention(), u2.Mention()), true)
			return
		}
	}

	e := model.ExcludedPair{UserId1: u1.ID, UserId2: u2.ID}

	ss = model.AddExcludedPair(ss, e)
	db.UpdateSantaSecret(ss)

	response.SendInteractionResponse(s, i, "Pair excluded!", true)
}

func removeExcludedMember(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
	ss model.SecretSanta,
	u *discordgo.User,
) {
	for _, e := range ss.ExcludedMembers {
		if e.UserId == u.ID {
			ss = model.RemoveExcludedMember(ss, e)
			db.UpdateSantaSecret(ss)

			response.SendInteractionResponse(s, i, fmt.Sprintf("The member %s is no longer excluded of the secret santa!", u.Mention()), true)
			return
		}
	}

	response.SendInteractionResponse(s, i, fmt.Sprintf("The member %s is not excluded!", u.Mention()), true)
}

func removeExcludedPair(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
	ss model.SecretSanta,
	u1 *discordgo.User,
	u2 *discordgo.User,
) {
	for _, e := range ss.ExcludedPairs {
		if e.UserId1 == u1.ID && e.UserId2 == u2.ID {
			ss = model.RemoveExcludedPair(ss, e)
			db.UpdateSantaSecret(ss)

			response.SendInteractionResponse(s, i, fmt.Sprintf("The pair %s and %s is no longer excluded of the secret santa!", u1.Mention(), u2.Mention()), true)
			return
		}
	}

	response.SendInteractionResponse(s, i, fmt.Sprintf("The pair %s and %s is not excluded!", u1.Mention(), u2.Mention()), true)
}

func listExcludedMembers(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
	ss model.SecretSanta,
) {
	embed := component.NewGenericEmbed(fmt.Sprintf("Secret Santa \"%s\" - Exclusion list", ss.Title), "List of excluded members and pairs for the secret santa")
	embed.Description += "\n\n"
	embed.Description += "Excluded members:\n"

	if len(ss.ExcludedMembers) == 0 {
		embed.Description += "No excluded members\n"
	} else {
		for _, e := range ss.ExcludedMembers {
			u, _ := s.User(e.UserId)
			embed.Description += fmt.Sprintf("%s\n", u.Mention())
		}
	}

	embed.Description += "\n"

	embed.Description += "Excluded pairs:\n"

	if len(ss.ExcludedPairs) == 0 {
		embed.Description += "No excluded pairs\n"
	} else {
		for _, e := range ss.ExcludedPairs {
			u1, _ := s.User(e.UserId1)
			u2, _ := s.User(e.UserId2)
			embed.Description += fmt.Sprintf("%s with %s\n", u1.Mention(), u2.Mention())
		}
	}

	response.SendInteractionEmbedResponse(s, i, embed, true)
}
