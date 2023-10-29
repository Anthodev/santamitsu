package handlers

import (
	"anthodev/santamitsu/utils"
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func pingResponse(s *discordgo.Session, i *discordgo.InteractionCreate) {
	utils.SendInteractionResponse(s, i, "Pong!", false)
}

func dmPingResponse(s *discordgo.Session, i *discordgo.InteractionCreate) {
	uc, err := s.UserChannelCreate(i.Member.User.ID)

	if err != nil {
		// If an error occurred, we failed to create the channel.
		//
		// Some common causes are:
		// 1. We don't share a server with the user (not possible here).
		// 2. We opened enough DM channels quickly enough for Discord to
		//    label us as abusing the endpoint, blocking us from opening
		//    new ones.
		fmt.Println("error creating channel:", err)
		_, _ = s.ChannelMessageSend(
			uc.ID,
			"Something went wrong while sending the DM!",
		)
		return
	}

	_, errDm := s.ChannelMessageSend(uc.ID, "DM Pong!")

	if errDm != nil {
		// If an error occurred, we failed to send the message.
		//
		// It may occur either when we do not share a server with the
		// user (highly unlikely as we just received a message) or
		// the user disabled DM in their settings (more likely).
		fmt.Println("error sending DM message:", errDm)
		_, _ = s.ChannelMessageSend(
			uc.ID,
			"Failed to send you a DM. "+
				"Did you disable DM in your privacy settings?",
		)
	}

	utils.SendInteractionResponse(s, i, "Check your DMs!", true)
}
