package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
)

const prefix string = "!ask" // Command prefix for the bot.

// openAI API request structure
type OpenAIRequest struct {
	Model    string `json:"model"`
	Messages []struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"messages"`
}

// openAI api response structure
type OpenAIResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Retrieve the bot token
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		log.Fatal("BOT_TOKEN is not set in the .env file")
	}

	// Retrieve the OpenAI API key
	openaiKey := os.Getenv("OPENAI_API_KEY")
	if openaiKey == "" {
		log.Fatal("OPENAI_API_KEY is not set in the .env file")
	}

	// Initialize OpenAI client
	openaiClient := openai.NewClient(openaiKey)

	// Create a new Discord session
	sess, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatalf("Error creating Discord session: %v", err)
	}

	// Handle messages
	sess.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		log.Println("Message received!") // This will log every message received
		if m.Author.ID == s.State.User.ID {
			return
		}
		// Log incoming messages to check if the bot is receiving them
		log.Printf("Received message: %s", m.Content)

		// Check if the message starts with the command prefix
		if strings.HasPrefix(m.Content, prefix) {
			query := strings.TrimPrefix(m.Content, prefix)
			query = strings.TrimSpace(query)

			if query == "" {
				s.ChannelMessageSend(m.ChannelID, "Please provide a query after the '/ask' command.")
				return
			}

			// Process the query with OpenAI
			log.Printf("Sending request to OpenAI: %+v", openai.ChatCompletionRequest{
				Model: openai.GPT3Dot5Turbo,
				Messages: []openai.ChatCompletionMessage{
					{Role: "system", Content: "You are a helpful assistant."},
					{Role: "user", Content: query},
				},
			})

			if err != nil {
				log.Printf("OpenAI error: %v", err)
				s.ChannelMessageSend(m.ChannelID, "Error processing your request. Please try again later.")
				return
			}

			// Check if a response is available
			if len(resp.Choices) == 0 {
				s.ChannelMessageSend(m.ChannelID, "I couldn't get a response. Please try again.")
				return
			}

			// Send the response to Discord
			s.ChannelMessageSend(m.ChannelID, resp.Choices[0].Message.Content)

			log.Printf("Response from OpenAI: %+v", resp)
		}
	})
	// Open the Discord session
	err = sess.Open()
	if err != nil {
		log.Fatalf("Error opening Discord session: %v", err)
	}
	defer sess.Close()

	fmt.Println("Bot is running. Press CTRL+C to exit.")
	select {}

}
