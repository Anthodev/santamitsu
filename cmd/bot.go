package cmd

import (
	"anthodev/santamitsu/cmds"
	"anthodev/santamitsu/handlers"
	"flag"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"strconv"
)

func init() { flag.Parse() }

var (
	s *discordgo.Session

	c = cmds.List()
	h = handlers.List()
	m = handlers.Buttons()
)

func init() {
	_ = godotenv.Load()
	dt := os.Getenv("DISCORD_TOKEN")
	var err error
	s, err = discordgo.New("Bot " + dt)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
	}

	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		switch i.Type {
		case discordgo.InteractionApplicationCommand:
			if h, ok := h[i.ApplicationCommandData().Name]; ok {
				h(s, i)
			}
		case discordgo.InteractionMessageComponent:
			if h, ok := m[i.MessageComponentData().CustomID]; ok {
				h(s, i)
			}
		}
	})
}

func Run() {
	gid := os.Getenv("DISCORD_GUILD")

	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		fmt.Printf("Logged in as %v#%v", s.State.User.Username, s.State.User.Discriminator)
		fmt.Println("")
	})

	err := s.Open()
	if err != nil {
		panic(err)
	}

	allCommands := append(c)

	fmt.Println("Adding the commands...")
	registeredCommands := make([]*discordgo.ApplicationCommand, len(allCommands))

	for i, v := range allCommands {
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, gid, v)
		if err != nil {
			log.Fatalf("Error creating command '%v': %v", v.Name, err)
		}
		registeredCommands[i] = cmd
	}

	defer func(s *discordgo.Session) {
		err := s.Close()
		if err != nil {
			log.Fatalf("Error closing Discord session: %v", err)
		}
	}(s)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	fmt.Println("Bot is running. Press Ctrl+C to exit.")
	<-stop

	removeCmds(registeredCommands)
}

func removeCmds(registeredCommands []*discordgo.ApplicationCommand) {
	rmEnv := os.Getenv("REMOVE_CMDS")
	rmEnvBool, _ := strconv.ParseBool(rmEnv)
	gid := os.Getenv("DISCORD_GUILD")

	if rmEnvBool {
		fmt.Println("")
		fmt.Println("Removing the commands...")
		for _, v := range registeredCommands {
			err := s.ApplicationCommandDelete(s.State.User.ID, gid, v.ID)
			if err != nil {
				fmt.Printf("Error removing command '%v': %v", v.Name, err)
			}
		}
	}
}
