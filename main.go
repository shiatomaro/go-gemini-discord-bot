package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
)

// Gemini API request structure
type GeminiRequest struct {
	Contents []struct {
		Parts []struct {
			Text string `json:"text"`
		} `json:"parts"`
	} `json:"contents"`
}

// Gemini API response structure
type GeminiResponse struct {
	Candidates []struct {
		Content struct {
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
		} `json:"content"`
	} `json:"candidates"`
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

	// Retrieve the Gemini API key
	geminiKey := os.Getenv("GEMINI_API_KEY")
	if geminiKey == "" {
		log.Fatal("GEMINI_API_KEY is not set in the .env file")
	}

	// Create a new Discord session
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatalf("Error creating Discord session: %v", err)
	}

	// start a discord session
	err = dg.Open()
	if err != nil {
		log.Fatalf("Error in starting a discord session: %v", err)
	}

	// Slash command actions
	registerSlashCommand(dg)

	dg.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		log.Printf("Received interaction: %v", i.ApplicationCommandData().Name)
		if i.ApplicationCommandData().Name == "ask" {
			handleAskCommand(s, i, geminiKey)
		}
	})

	// to close discord session
	defer dg.Close()
	log.Println(" Go ask away is running to exist, please press CTRL+C.")
	select {}

}

// Function to call OpenAI API
func registerSlashCommand(s *discordgo.Session) {
	command := &discordgo.ApplicationCommand{
		Name:        "ask",
		Description: "go ask away any question",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "question",
				Description: "the question for go ask away to answer",
				Type:        discordgo.ApplicationCommandOptionString,
				Required:    true,
			},
		},
	}

	_, err := s.ApplicationCommandCreate(s.State.User.ID, "", command)
	if err != nil {
		log.Fatalf("Error in creating a slash command: %v", err)
	}
	log.Println("Slash command / registered , success!")
}

func handleAskCommand(s *discordgo.Session, i *discordgo.InteractionCreate, geminiKey string) {
	log.Println("Handling /ask command interaction.")

	question := i.ApplicationCommandData().Options[0].StringValue()
	log.Printf("User's question: %s", question)

	// Send a deferred response to Discord
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
	})
	if err != nil {
		log.Printf("Error sending preliminary response: %v", err)
		return
	}

	// Call Gemini API
	response, err := getGeminiResponse(geminiKey, question)
	if err != nil {
		log.Printf("Error getting response from Gemini API: %v", err)
		// Send an error message as a follow-up
		followUpMessage(s, i.Interaction, "Failed to get a response from Gemini API.")
		return
	}

	// Truncate response if too long for Discord
	if len(response) > 2000 {
		response = response[:1997] + "..."
	}

	// Send the final response as a follow-up message
	followUpMessage(s, i.Interaction, response)
}

// Function to send a follow-up message
func followUpMessage(s *discordgo.Session, i *discordgo.Interaction, message string) {
	_, err := s.FollowupMessageCreate(i, false, &discordgo.WebhookParams{
		Content: message,
	})
	if err != nil {
		log.Printf("Error sending follow-up message: %v", err)
	}
}

// Function to call gemini
func getGeminiResponse(geminiKey, userInput string) (string, error) {
	client := resty.New()

	request := GeminiRequest{
		Contents: []struct {
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
		}{
			{
				Parts: []struct {
					Text string `json:"text"`
				}{
					{Text: userInput},
				},
			},
		},
	}

	log.Println("Sending request to Gemini API...")
	url := fmt.Sprintf("https://generativelanguage.googleapis.com/v1beta/models/gemini-1.5-flash:generateContent?key=%s", geminiKey)

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(request).
		Post(url)

	if err != nil {
		return "", err
	}

	log.Printf("Gemini API Response: %s", string(resp.Body()))

	var geminiResponse GeminiResponse
	err = json.Unmarshal(resp.Body(), &geminiResponse)
	if err != nil {
		return "", err
	}

	if len(geminiResponse.Candidates) > 0 && len(geminiResponse.Candidates[0].Content.Parts) > 0 {
		return geminiResponse.Candidates[0].Content.Parts[0].Text, nil
	}

	return "No response from Gemini.", nil
}
