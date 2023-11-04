package component

import "github.com/bwmarrin/discordgo"

func JoinActionRow() *discordgo.ActionsRow {
	return &discordgo.ActionsRow{
		Components: []discordgo.MessageComponent{
			&discordgo.Button{
				Label:    "Join",
				Style:    discordgo.SuccessButton,
				CustomID: "join",
			},
		},
	}
}
