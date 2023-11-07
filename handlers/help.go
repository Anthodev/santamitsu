package handlers

import (
	"anthodev/santamitsu/utils/component"
	"anthodev/santamitsu/utils/response"
	"github.com/bwmarrin/discordgo"
)

func helpHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	t := "Secret Santa Creator - List of most useful commands"
	msg := "Most of the commands are self-explanatory, but here's a quick rundown of the most useful commands:\n" +
		"- `/setup` - Setup the secret santa (admin only)\n" +
		"- `/announce` - Announce the secret santa (admin only)\n" +
		"- `/info` - Get info about the current secret santa\n" +
		"- `/join` - Join the current secret santa\n" +
		"- `/leave` - Leave the current secret santa\n" +
		"- `/cancel` - Cancel the current secret santa (admin only)\n" +
		"- `/delete` - Delete the current secret santa (admin only)\n"

	embed := component.NewGenericEmbed(t, msg)

	response.SendInteractionEmbedResponse(s, i, embed, true)
}
