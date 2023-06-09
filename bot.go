package main

import (
	"log"
	"strings"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"commands/admin.go"
)
// Rest of the code...


func main() {
	bot, err := tgbotapi.NewBotAPI("6124526418:AAGKcmirRbl2MYEQuiun1r3oV_ZBmTiWaME")
	if err != nil {
		log.Fatal(err)
	}


	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		message := update.Message

		// Check if the command is for the admin
		if strings.HasPrefix(message.Text, "/admin") {
			commands.HandleAdminCommand(bot, message)
		} else {
			// Handle other commands
			// ...
		}
	}
}