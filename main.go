package main

import (
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

// go get from .env file
godotenv.Load()

const prefix string = "/"

func main() {
	token := os.Getenv("BOT_TOKEN")
	// to print if there's a token
	log.Println("the bot token is " + token)
	sess, err := discordgo.New("Bot " + token)
	if err != nil {
		// to create error if there's not token
		log.fatal(err)
	}
}
