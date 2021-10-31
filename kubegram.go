package main

import (
	"os"
	f "kubegram/tools"
	bot "kubegram/bot"
)

func main() {
	f.CreateFolder()
	bot.KubeTelegramBot(os.Getenv("BOT_TOKEN"))
}
