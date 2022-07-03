package bank

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/ozonmp/omp-bot/internal/app/commands/bank/card"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type BankCommander struct {
	bot           *tgbotapi.BotAPI
	cardCommander Commander
}

func NewBankCommander(
	bot *tgbotapi.BotAPI,
) *BankCommander {
	return &BankCommander{
		bot:           bot,
		cardCommander: card.NewBankCardCommander(bot),
	}
}

func (c *BankCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "card":
		c.cardCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("BankCommander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

func (c *BankCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "card":
		c.cardCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("BankCommander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
	}
}
