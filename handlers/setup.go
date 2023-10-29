package handlers

import (
	"anthodev/santamitsu/common/component"
	"anthodev/santamitsu/utils"
	"github.com/bwmarrin/discordgo"
	"strconv"
)

type dm struct {
	ChannelID   string
	UserID      string
	Title       string
	Description string
	MaxPrice    int
}

var (
	responses map[string]dm = map[string]dm{}
)

func setupHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	uc := utils.CreateDmChannel(s, i.Member.User.ID)

	setupWizard(s, i, uc)
}

func setupWizard(s *discordgo.Session, i *discordgo.InteractionCreate, uc *discordgo.Channel) {
	delete(responses, uc.ID)

	if _, ok := responses[uc.ID]; !ok {
		responses[uc.ID] = dm{
			ChannelID: uc.ID,
			UserID:    i.Member.User.ID,
		}

		askTitle(s, i, uc)
	}
}

func askTitle(s *discordgo.Session, i *discordgo.InteractionCreate, uc *discordgo.Channel) {
	embed := component.NewEmbed()
	embed.Title = "Secret Santa Setup"
	embed.Description = "Give a title to your secret santa!"

	utils.SendDmEmbed(s, uc.ID, embed.MessageEmbed)

	s.AddHandlerOnce(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		contentMsgHandler(s, m)

		if m.Content != "" {
			response := responses[uc.ID]
			response.Title = m.Content

			askDescription(uc, s)
		}
	})
}

func askDescription(uc *discordgo.Channel, s *discordgo.Session) {
	embed := component.NewEmbed()
	embed.Title = "Secret Santa Description"
	embed.Description = "Give a description to your secret santa!"

	utils.SendDmEmbed(s, uc.ID, embed.MessageEmbed)

	s.AddHandlerOnce(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		contentMsgHandler(s, m)

		if m.Content != "" {
			response := responses[uc.ID]
			response.Description = m.Content

			askPrice(uc, s)
		}
	})
}

//
func askPrice(uc *discordgo.Channel, s *discordgo.Session) {
	embed := component.NewEmbed()
	embed.Title = "Secret Santa Description"
	embed.Description = "Give a description to your secret santa!"

	utils.SendDmEmbed(s, uc.ID, embed.MessageEmbed)

	s.AddHandlerOnce(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		contentMsgHandler(s, m)

		if m.Content != "" {
			response := responses[uc.ID]

			price, err := strconv.Atoi(m.Content)
			if err != nil {
				askPrice(uc, s)
			}

			response.MaxPrice = price

			askPrice(uc, s)
		}
	})
}

func contentMsgHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.ChannelID == responses[m.ChannelID].ChannelID {
		return
	}
}
