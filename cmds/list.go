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
		//{
		//	Name:        "info",
		//	Description: "Get info about the current santa secret if one is running",
		//	GuildID:     *env.GuildID,
		//},
		{
			Name:        "announce",
			Description: "Announce the santa secret",
			GuildID:     gid,
		},
		{
			Name:         "setup",
			Description:  "Setup a new santa secret",
			GuildID:      gid,
			DMPermission: &dmPermission,
		},
		{
			Name:        "cancel",
			Description: "Cancel a santa secret",
			GuildID:     gid,
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
