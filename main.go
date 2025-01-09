package main

import (
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

// get from .env file.
godotenv.Load()

const prefix string = "/goaskaway" // command prefix for bot.


func main() {
	// to load token from .env file.
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		// to check if BOT_TOKEN is set in .env file.
		log.Println("BOT_TOKEN not set in .env file. Set BOT_TOKEN in envirorment")
	}

	sess, err := discordgo.New("Bot " + token)
	if err != nil {
		// to create error if there's not token
		log.fatal("No discord token found, check .env file if BOT_TOKEN is set in envirorment.")
	}
}

func sess.AddHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
}
