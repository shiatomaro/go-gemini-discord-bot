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
	// to load key from .env file.
	err := godotenv.Load()
	if err != nil {
		// to check if BOT_TOKEN is set in .env file.
		log.Fatal("OPEN_AI_KEY not set in .env file.")
	}

	// to get OpenAI API KEY
	openaiKey := os.Getenv("OPENAI_API_KEY")
	if openaiKey == "" {
		log.Fatal("OPENAI_API_KEY is not set in the .env file.")
	}

	// initiialize OpenAI Client
	aiClient := openai.NewClient(openaiKey)

	// to load token from .env file.
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		// to check if BOT_TOKEN is set in .env file.
		log.Println("BOT_TOKEN not set in .env file. Set BOT_TOKEN in envirorment")

	// to create new discord session.
	sess, err := discordgo.New("Bot " + token)
	if err != nil {
		// to create error if there's no token
		log.fatal("No discord token found, check .env file if BOT_TOKEN is set in envirorment.", err)
	}
}
	// handle messages
	dg.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		//Ignore bots own response
		if m.Author.ID == s.State.User.ID {
		return
	}
		// Command Prefix 
		const prefix = "/ask"

		// to check if the message starts with prefix 
		if String.HasPrefix(m.Content, prefix) {
			query := string.TrimPrefix(m.Content, prefix)
			query := string.TrimSpace(query)

			if query == "" {
				s.ChannelMessageSend(m.ChannelID, "Please use '/ask' to use command.")
				return
			}
		}

}
