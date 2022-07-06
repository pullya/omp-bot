package card

import (
	"encoding/json"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/model/bank"
)

type CallbackListData struct {
	Offset int `json:"offset"`
}

func (c *cardCommander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	parsedData := CallbackListData{}
	err := json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData)
	if err != nil {
		log.Printf("BankCardCommander.CallbackList: "+
			"error reading json data for type CallbackListData from "+
			"input string %v - %v", callbackPath.CallbackData, err)
		return
	}

	outputMessage := ""
	cur := parsedData.Offset
	curI := cur

	products, err := c.sservice.List(uint64(cur), 5)
	if err != nil {
		log.Println("smthg went wrong")
		return
	}

	for i := range products {
		outputMessage += "\n" + products[i].Title
		log.Printf("listing cur i: %v", curI)
		curI++
	}

	if curI >= len(bank.AllCards) {
		curI = 0
	}

	msg := tgbotapi.NewMessage(callback.Message.Chat.ID,
		outputMessage,
	)
	serializedData, _ := json.Marshal(CallbackListData{
		Offset: curI,
	})

	callbackPath = path.CallbackPath{
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

	// msg := tgbotapi.NewMessage(
	// 	callback.Message.Chat.ID,
	// 	fmt.Sprintf("Parsed: %+v\n", parsedData),
	// )
	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("cardCommander.CallbackList: error sending reply message to chat - %v", err)
	}
}
