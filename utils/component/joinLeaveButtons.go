package component

import "github.com/bwmarrin/discordgo"

func JoinLeaveActionRow() *discordgo.ActionsRow {
	return &discordgo.ActionsRow{
		Components: []discordgo.MessageComponent{
			&discordgo.Button{
				Label:    "Join",
				Style:    discordgo.SuccessButton,
				CustomID: "join",
			},
			&discordgo.Button{
				Label:    "Leave",
				Style:    discordgo.PrimaryButton,
				CustomID: "leave",
			},
		},
	}
}
