package handlers

import (
	"anthodev/santamitsu/db"
	"anthodev/santamitsu/model"
	"anthodev/santamitsu/service"
	"anthodev/santamitsu/utils/component"
	"anthodev/santamitsu/utils/response"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strconv"
)

var (
	responses = map[string]model.SetupSettings{}
)

func setupHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	service.IsAdmin(s, i)

	uc := response.CreateDmChannel(s, i.Member.User.ID)

	setupWizard(s, i, uc)
}

func setupWizard(s *discordgo.Session, i *discordgo.InteractionCreate, uc *discordgo.Channel) {
	delete(responses, i.GuildID)

	if _, ok := responses[i.GuildID]; !ok {
		responses[i.GuildID] = model.SetupSettings{
			ChannelID: i.GuildID,
			UserID:    i.Member.User.ID,
			State:     1,
		}

		askTitle(s, i, uc)
	}
}

func askTitle(s *discordgo.Session, i *discordgo.InteractionCreate, uc *discordgo.Channel) {
	embed := component.NewGenericEmbed(
		"Secret Santa Setup",
		"Give a title to your secret santa",
	)

	response.SendDmEmbed(s, uc.ID, embed)

	s.AddHandlerOnce(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		contentMsgHandler(s, i, m)

		if m.Content != "" {
			r := responses[i.GuildID]
			r.Title = m.Content

			askDescription(uc, s, i, r)
		}
	})
}

func askDescription(uc *discordgo.Channel, s *discordgo.Session, i *discordgo.InteractionCreate, r model.SetupSettings) {
	embed := component.NewGenericEmbed(
		"Description",
		"Give a description to your secret santa",
	)

	response.SendDmEmbed(s, uc.ID, embed)

	s.AddHandlerOnce(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		contentMsgHandler(s, i, m)

		if m.Content != "" {
			r.Description = m.Content

			askPrice(uc, s, i, r)
		}
	})
}

func askPrice(uc *discordgo.Channel, s *discordgo.Session, i *discordgo.InteractionCreate, r model.SetupSettings) {
	embed := component.NewGenericEmbed(
		"Price",
		"Indicate the maximum price for a gift",
	)

	response.SendDmEmbed(s, uc.ID, embed)

	s.AddHandlerOnce(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		contentMsgHandler(s, i, m)

		if m.Content != "" {
			r.MaxPrice = m.Content

			askCurrency(uc, s, i, r)
		}
	})
}

func askCurrency(uc *discordgo.Channel, s *discordgo.Session, i *discordgo.InteractionCreate, r model.SetupSettings) {
	embed := component.NewGenericEmbed(
		"Currency",
		"Indicate the currency of the price",
	)

	currencies := model.BuildCurrencies()

	for _, v := range currencies {
		embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
			Name:   fmt.Sprintf("`%v`: ", v.Id),
			Value:  fmt.Sprintf("%v", v.Value),
			Inline: true,
		})
	}

	response.SendDmEmbed(s, uc.ID, embed)

	s.AddHandlerOnce(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		contentMsgHandler(s, i, m)

		if m.Content != "" {
			content := m.Content

			_, err := strconv.Atoi(content)

			if err != nil {
				askCurrency(uc, s, i, r)
				return
			}

			r.Currency = convertCurrency(content)
			responses[i.GuildID] = r

			ss := model.CreateSantaSecret(r)
			db.InsertSantaSecret(ss)
			completedWizard(uc, s)

			delete(responses, i.GuildID)
		}
	})
}

func convertCurrency(currency string) string {
	switch currency {
	case "1":
		return "€"
	case "2":
		return "$"
	default:
		return ""
	}
}

func completedWizard(uc *discordgo.Channel, s *discordgo.Session) {
	embed := component.NewGenericEmbed(
		"Completed",
		"Your secret santa has been created! Use `/announce` on your server to announce it!",
	)

	response.SendDmEmbed(s, uc.ID, embed)
}

func contentMsgHandler(s *discordgo.Session, i *discordgo.InteractionCreate, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if i.ChannelID != responses[i.ChannelID].ChannelID {
		return
	}
}
