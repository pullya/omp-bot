package card

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	//	model "github.com/ozonmp/omp-bot/internal/model/bank"
	service "github.com/ozonmp/omp-bot/internal/service/bank/card"
)

type CardCommander interface {
	Help(inputMsg *tgbotapi.Message)
	Get(inputMsg *tgbotapi.Message)
	List(inputMsg *tgbotapi.Message)
	Delete(inputMsg *tgbotapi.Message)

	New(inputMsg *tgbotapi.Message)  // return error
	Edit(inputMsg *tgbotapi.Message) // return error
}

//func NewCardCommander(bot *tgbotapi.BotAPI, service service.CardService) CardCommander {
// ...
//}

type NewCardCommander struct {
	bot     *tgbotapi.BotAPI
	service *service.DummyCardService
}

func NewBankCardCommander(bot *tgbotapi.BotAPI) *NewCardCommander {
	cardService := service.NewDummyCardService()
	return &NewCardCommander{
		bot:     bot,
		service: cardService,
	}
}

func (c *NewCardCommander) Help(inputMsg *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID,
		"Help command service",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("BankCardCommander.Help: error sending reply message to chat - %v", err)
	}
}

func (c *NewCardCommander) Get(inputMsg *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID,
		"Get command service",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("BankCardCommander.Get: error sending reply message to chat - %v", err)
	}
}

func (c *NewCardCommander) List(inputMsg *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID,
		"List command service",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("BankCardCommander.List: error sending reply message to chat - %v", err)
	}
}

func (c *NewCardCommander) Delete(inputMsg *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID,
		"Delete command service",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("BankCardCommander.Delete: error sending reply message to chat - %v", err)
	}
}

func (c *NewCardCommander) New(inputMsg *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID,
		"New command service",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("BankCardCommander.New: error sending reply message to chat - %v", err)
	}
}

func (c *NewCardCommander) Edit(inputMsg *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID,
		"Edit command service",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("BankCardCommander.Edit: error sending reply message to chat - %v", err)
	}
}
