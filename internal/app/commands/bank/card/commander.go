package card

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	//  	model "github.com/ozonmp/omp-bot/internal/model/bank"
	"github.com/ozonmp/omp-bot/internal/app/path"
	service "github.com/ozonmp/omp-bot/internal/service/bank/card"
)

type BankCardService struct{}

func NewBankCardService() *BankCardService {
	return &BankCardService{}
}

type BankCardCommander struct {
	bot             *tgbotapi.BotAPI
	bankCardService CardCommander
}

func NewBankCardCommander(bot *tgbotapi.BotAPI) *BankCardCommander {
	dummyService := service.NewDummyCardService()
	return &BankCardCommander{
		bot:             bot,
		bankCardService: NewCardCommander(bot, dummyService),
	}
}

func (c *BankCardCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.bankCardService.Help(msg)
	case "get":
		c.bankCardService.Get(msg)
	case "list":
		c.bankCardService.List(msg)
	case "edit":
		c.bankCardService.Edit(msg)
	case "delete":
		c.bankCardService.Delete(msg)
	case "new":
		c.bankCardService.New(msg)
	default:
		log.Printf("TBD - %s", commandPath.Subdomain)
	}
}

func (c *BankCardCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.bankCardService.CallbackList(callback, callbackPath)
	default:
		log.Printf("DemoSubdomainCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

type CardCommander interface {
	Help(inputMsg *tgbotapi.Message)
	Get(inputMsg *tgbotapi.Message)
	List(inputMsg *tgbotapi.Message)
	Delete(inputMsg *tgbotapi.Message)

	New(inputMsg *tgbotapi.Message)  // return error
	Edit(inputMsg *tgbotapi.Message) // return error

	CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
}

type cardCommander struct {
	bot      *tgbotapi.BotAPI
	sservice *service.DummyCardService
}

func NewCardCommander(bot *tgbotapi.BotAPI, sservice service.CardService) *cardCommander {
	dummyService := service.NewDummyCardService()
	return &cardCommander{
		bot:      bot,
		sservice: dummyService,
	}
}
