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

func ModeratorRoleHandler(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
	ss model.SantaSecret,
) {
	service.IsAdmin(s, i)

	switch i.Type {
	case discordgo.InteractionApplicationCommand:
		data := i.ApplicationCommandData()
		subcmd := data.Options[0]

		if subcmd.Name == "add-role" {
			addRole(s, i, ss, subcmd.Options[0].RoleValue(s, i.GuildID))
			return
		}

		if subcmd.Name == "remove-role" {
			removeRole(s, i, ss, subcmd.Options[0].RoleValue(s, i.GuildID))
			return
		}

		if subcmd.Name == "list-roles" {
			listRoles(s, i, ss)
			return
		}
	case discordgo.InteractionApplicationCommandAutocomplete:
		data := i.ApplicationCommandData()
		var choices []*discordgo.ApplicationCommandOptionChoice

		switch data.Options[0].Options[0].Name {
		case "role":
			if data.Options[0].Options[0].StringValue() != "" {
				choices = append(choices, &discordgo.ApplicationCommandOptionChoice{
					Name:  data.Options[0].Name,
					Value: data.Options[0].RoleValue(s, i.GuildID).Mention(),
				})
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
}

func addRole(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
	ss model.SantaSecret,
	r *discordgo.Role,
) {
	mr := model.ModeratorRole{
		RoleId: r.ID,
	}

	ss = model.AddModeratorRole(ss, mr)

	db.UpdateSantaSecret(ss)

	response.SendInteractionResponse(s, i, fmt.Sprintf("Role %s added as a moderator role of the secret santa", r.Mention()), true)
}

func removeRole(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
	ss model.SantaSecret,
	r *discordgo.Role,
) {
	for _, mr := range ss.ModeratorRoles {
		if mr.RoleId == r.ID {
			ss = model.RemoveModeratorRole(ss, mr)
			db.UpdateSantaSecret(ss)
			response.SendInteractionResponse(s, i, fmt.Sprintf("Role %s removed as a moderator role of the secret santa", r.Mention()), true)
			return
		}
	}

	response.SendInteractionResponse(s, i, fmt.Sprintf("Role %s is not a moderator role of the secret santa", r.Mention()), true)
}

func listRoles(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
	ss model.SantaSecret,
) {
	embed := component.NewGenericEmbed(fmt.Sprintf("Moderator roles of %s", ss.Title), "List of the roles that can manage the secret santa")
	embed.Description += "\n\n"
	embed.Description += "Moderator roles:\n"

	if len(ss.ModeratorRoles) == 0 {
		embed.Description += "No roles added\n"
	} else {
		for _, mr := range ss.ModeratorRoles {
			mr, _ := s.State.Role(i.GuildID, mr.RoleId)
			embed.Description += fmt.Sprintf("%s\n", mr.Mention())
		}
	}

	response.SendInteractionEmbedResponse(s, i, embed, true)
}
