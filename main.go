package main

import (
	"github.com/joho/godotenv"
	"gobot/bot"
	"gobot/task_scheduler"
	"log"
	"os"
)

// https://discord.com/oauth2/authorize?client_id=1257416801734889503&permissions=8&integration_type=0&scope=bot
func main() {
	defer task_scheduler.TaskScheduler.Stop()

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	botToken := os.Getenv("BOT_TOKEN")
	guildId := os.Getenv("GUILD_ID")
	appId := os.Getenv("APP_ID")

	if botToken == "" || guildId == "" || appId == "" {
		log.Fatal("Missing environment variables (BOT_TOKEN, GUILD_ID, APP_ID  required)")
	}

	b := bot.New(botToken, guildId, appId)
	b.Run()
}
