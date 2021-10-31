package bot

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	tools "kubegram/tools"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type BotAPI struct {
	Token string
}

type FileConfig struct {
	FileID string
}

type File struct {
	FileID string `json:"file_id"`
}

func KubeTelegramBot(Token string) (*BotAPI, error) {

	token := BotAPI{Token: Token}

	var chat_id_string string = os.Getenv("CHAT_ID")
	chat_id, _ := strconv.Atoi(chat_id_string)

	bot, err := tgbotapi.NewBotAPI(token.Token)

	if err != nil {
		panic(err)
	}
	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(30)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		msg.ParseMode = "Markdown"

		if update.Message.Chat.ID == int64(chat_id) {

			switch update.Message.Text {
			case "ls":
				msg.Text = tools.GetFiles("ls", os.Getenv("FILE_STORAGE"))
			case "help":
				msg.Text = "kubectl client for telegram"
			case "k3":
				msg.Text = tools.RewriteK3()
			case "mykub":
				msg.Text = tools.MyKubeConfig()
			default:
				respone := update.Message.Text
				words := strings.Split(respone, " ")

				if strings.Contains(respone, "cat") {
					replace_response := strings.Replace(respone, "cat ", "", -1)
					fmt.Println(replace_response)
					msg.Text = tools.CatFiles("cat", os.Getenv("FILE_STORAGE")+replace_response)
				} else {
					msg.Text = tools.GetCommand("kubectl", words...)
				}

			}
		} else {
			msg.Text = "Not validate user, fuck off"
		}

		if _, err := bot.Send(msg); err != nil {
			fmt.Println("just bad")
		}
	}
	return &token, err
}
