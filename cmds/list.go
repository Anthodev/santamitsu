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
			Name:        "info",
			Description: "Get info about the current Secret Santa if one is running",
			GuildID:     gid,
		},
		{
			Name:        "announce",
			Description: "Announce the Secret Santa",
			GuildID:     gid,
		},
		{
			Name:         "setup",
			Description:  "Setup a new Secret Santa",
			GuildID:      gid,
			DMPermission: &dmPermission,
		},
		{
			Name:        "cancel",
			Description: "Cancel a Secret Santa",
			GuildID:     gid,
		},
		{
			Name:        "join",
			Description: "Join the current Secret Santa",
			GuildID:     gid,
		},
		{
			Name:        "leave",
			Description: "Leave the current Secret Santa",
			GuildID:     gid,
		},
		{
			Name:        "exclude",
			Description: "Exclude a member from the current Secret Santa",
			GuildID:     gid,
			Type:        discordgo.ChatApplicationCommand,
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "member",
					Description: "The member to exclude for the santa",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Options: []*discordgo.ApplicationCommandOption{
						{
							Name:         "user",
							Description:  "The member to exclude for the santa",
							Type:         discordgo.ApplicationCommandOptionUser,
							Required:     true,
							Autocomplete: true,
						},
					},
				},
				{
					Name:        "pair",
					Description: "The pair to exclude for the santa",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Options: []*discordgo.ApplicationCommandOption{
						{
							Name:         "member-1",
							Description:  "First member to exclude for the pair",
							Type:         discordgo.ApplicationCommandOptionUser,
							Required:     true,
							Autocomplete: true,
						},
						{
							Name:         "member-2",
							Description:  "Second member to exclude for the pair",
							Type:         discordgo.ApplicationCommandOptionUser,
							Required:     true,
							Autocomplete: true,
						},
					},
				},
				{
					Name:        "remove",
					Description: "Remove an excluded member for the santa",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Options: []*discordgo.ApplicationCommandOption{
						{
							Name:         "rm-member",
							Description:  "The member to remove from the excluded list",
							Type:         discordgo.ApplicationCommandOptionUser,
							Required:     true,
							Autocomplete: true,
						},
					},
				},
				{
					Name:        "remove-pair",
					Description: "Remove an excluded pair for the santa",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Options: []*discordgo.ApplicationCommandOption{
						{
							Name:         "rm-member-1",
							Description:  "First member of the pair to remove from the excluded list",
							Type:         discordgo.ApplicationCommandOptionUser,
							Required:     true,
							Autocomplete: true,
						},
						{
							Name:         "rm-member-2",
							Description:  "Second member of the pair to remove from the excluded list",
							Type:         discordgo.ApplicationCommandOptionUser,
							Required:     true,
							Autocomplete: true,
						},
					},
				},
				{
					Name:        "list",
					Description: "List all excluded members and pairs for the santa",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
				},
			},
		},
		{
			Name:        "moderator-role",
			Description: "Add a moderator role to the Secret Santa",
			GuildID:     gid,
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "add-role",
					Description: "Add a moderator role to the Secret Santa",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Options: []*discordgo.ApplicationCommandOption{
						{
							Name:         "role",
							Description:  "The role to add as a moderator",
							Type:         discordgo.ApplicationCommandOptionRole,
							Required:     true,
							Autocomplete: true,
						},
					},
				},
				{
					Name:        "remove-role",
					Description: "Remove a moderator role from the Secret Santa",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Options: []*discordgo.ApplicationCommandOption{
						{
							Name:         "role",
							Description:  "The role to remove as a moderator",
							Type:         discordgo.ApplicationCommandOptionRole,
							Required:     true,
							Autocomplete: true,
						},
					},
				},
				{
					Name:        "list-roles",
					Description: "List all moderator roles for the Secret Santa",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
				},
			},
		},
		{
			Name:        "lock",
			Description: "Lock the current Secret Santa",
			GuildID:     gid,
		},
		{
			Name:        "unlock",
			Description: "Unlock the current Secret Santa",
			GuildID:     gid,
		},
		{
			Name:        "draw",
			Description: "Draw the pairs for the current Secret Santa",
			GuildID:     gid,
		},
		{
			Name:        "get-pair",
			Description: "Get your match for the Secret Santa",
			GuildID:     gid,
		},
		{
			Name:        "delete",
			Description: "Delete the Secret Santa",
			GuildID:     gid,
		},
		{
			Name:        "set-budget",
			Description: "Set the budget for the Secret Santa",
			GuildID:     gid,
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "maximum-price",
					Description: "Set the new maximum budget for the Secret Santa",
					Type:        discordgo.ApplicationCommandOptionString,
					Required:    true,
				},
			},
		},
		{
			Name:        "help",
			Description: "Get help about the bot",
		},
	}
}
