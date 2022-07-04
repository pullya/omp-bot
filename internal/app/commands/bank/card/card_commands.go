package card

import (
	"encoding/json"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

func (c *cardCommander) Help(inputMsg *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID,
		"/help - print list of commands\n"+
			"/get - print info about card\n"+
			"/list - print list of cards\n"+
			"/delete - delete card from list\n"+
			"/add - add card to list\n"+
			"/edit - edit card information",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("BankCardCommander.Help: error senn ding reply message to chat - %v", err)
	}
}

func (c *cardCommander) Get(inputMsg *tgbotapi.Message) {
	args := inputMsg.CommandArguments()

	idx, err := strconv.ParseUint(args, 10, 64)
	if err != nil {
		log.Println("wrong args", args)
		return
	}

	product, err := c.sservice.Describe(idx)
	//	product, err := c.subdomainService.Get(idx)
	if err != nil {
		log.Printf("fail to get product with idx %d: %v", idx, err)
		return
	}

	msg := tgbotapi.NewMessage(inputMsg.Chat.ID,
		product.Title,
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("BankCardCommander.Get: error sending reply message to chat - %v", err)
	}
}

func (c *cardCommander) List(inputMsg *tgbotapi.Message) {
	outputMessage := "Here all card types:"

	products, err := c.sservice.List(0, 3)
	if err != nil {
		log.Println("smthg went wrong")
		return
	}

	for i := range products {
		outputMessage += "\n" + products[i].Title
	}

	msg := tgbotapi.NewMessage(inputMsg.Chat.ID,
		outputMessage,
	)
	serializedData, _ := json.Marshal(CallbackListData{
		Offset: 19,
	})

	callbackPath := path.CallbackPath{
		Domain:       "bank",
		Subdomain:    "card",
		CallbackName: "list",
		CallbackData: string(serializedData),
	}

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPath.String()),
		),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("BankCardCommander.List: error sending reply message to chat - %v", err)
	}
}

func (c *cardCommander) Delete(inputMsg *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID,
		"Delete command service",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("BankCardCommander.Delete: error sending reply message to chat - %v", err)
	}
}

func (c *cardCommander) New(inputMsg *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID,
		"New command service",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("BankCardCommander.New: error sending reply message to chat - %v", err)
	}
}

func (c *cardCommander) Edit(inputMsg *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID,
		"Edit command service",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("BankCardCommander.Edit: error sending reply message to chat - %v", err)
	}
}
