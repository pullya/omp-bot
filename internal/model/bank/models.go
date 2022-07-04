package bank

type Card struct {
	Title string
}

var AllCards = []Card{
	{Title: "debet"},
	{Title: "credit"},
	{Title: "unembossed"},
	{Title: "virtual"},
	{Title: "cash"},
	{Title: "token"},
}
