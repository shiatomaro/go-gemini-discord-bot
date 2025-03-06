# Go Ask Away - Discord Bot with Gemini API

This is a **Discord bot** built with **Go**, leveraging the [DiscordGo](https://github.com/bwmarrin/discordgo) library to interact with Discord and the **Gemini API** to provide AI-powered responses to user queries. The bot registers a slash command `/ask` which allows users to ask any question and receive AI-generated responses.

## Features  
- **Discord Slash Command** (`/ask`) for user interaction  
- **AI-powered responses** using the Gemini API  
- **Handles errors and API response limits**  
- **Loads secrets securely** from a `.env` file  

## Prerequisites  
- Go installed on your system  
- A Discord bot token (set in a `.env` file)  
- A Gemini API key (set in a `.env` file)  

## Installation & Setup  

1. Clone the repository:  
   ```sh
   git clone https://github.com/your-username/go-ask-away.git
   cd go-ask-away
   ```  
2. Install dependencies:  
   ```sh
   go mod tidy
   ```  
3. Create a `.env` file and add your credentials:  
   ```ini
   BOT_TOKEN=your_discord_bot_token
   GEMINI_API_KEY=your_gemini_api_key
   ```  
4. Run the bot:  
   ```sh
   go run main.go
   ```  

## Usage  
Once the bot is running and invited to a server, use the `/ask` command followed by a question. Example:  
```
/ask What is the capital of France?
```
The bot will fetch a response from the Gemini API and return it to the user.

## License  
This project is licensed under the MIT License.  

