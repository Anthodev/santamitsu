package cmds

import (
	"github.com/bwmarrin/discordgo"
	"os"
)

var (
	gid = os.Getenv("DISCORD_GUILD")
)

func List() []*discordgo.ApplicationCommand {
	dmPermission := true

	return []*discordgo.ApplicationCommand{
		{
			Name:        "ping",
			Description: "Ping the bot",
		},
		{
			Name:         "dmping",
			Description:  "Ping the bot with a response in a dm",
			GuildID:      gid,
			DMPermission: &dmPermission,
		},
		//{
		//	Name:        "info",
		//	Description: "Get info about the current santa secret if one is running",
		//	GuildID:     *env.GuildID,
		//},
		{
			Name:         "setup",
			Description:  "Setup a new santa secret",
			GuildID:      gid,
			DMPermission: &dmPermission,
		},
		//{
		//	Name:        "join",
		//	Description: "Join the current santa secret",
		//	GuildID:     *env.GuildID,
		//},
		//{
		//	Name:        "leave",
		//	Description: "Leave the current santa secret",
		//	GuildID:     *env.GuildID,
		//},
		//{
		//	Name:        "lock",
		//	Description: "Lock and start the current santa secret",
		//	GuildID:     *env.GuildID,
		//},
	}
}
