package bank

type Card struct {
	Title string
}

var AllCards = []Card{
	{Title: "debet"},
	{Title: "credit"},
	{Title: "unempossed"},
	{Title: "virtual"},
	{Title: "cash"},
	{Title: "token"},
}
