package card

import (
	bank "github.com/ozonmp/omp-bot/internal/model/bank"
)

type CardService interface {
	Describe(cardId uint64) (*bank.Card, error)
	List(cursor uint64, limit uint64) ([]bank.Card, error)
	Create(bank.Card) (int, error)
	Update(cardID uint64, card bank.Card) error
	Remove(cardID uint64) (bool, error)
}

type DummyCardService struct{}

func NewDummyCardService() *DummyCardService {
	return &DummyCardService{}
}

func (d *DummyCardService) Describe(cardId uint64) (*bank.Card, error) {
	return &bank.AllCards[cardId], nil
}

func (d *DummyCardService) List(cursor uint64, limit uint64) ([]bank.Card, error) {
	return bank.AllCards, nil
}

func (d *DummyCardService) Create(bank.Card) (int, error) {
	return len(bank.AllCards), nil
}

func (d *DummyCardService) Update(cardID uint64, card bank.Card) error {
	return nil
}

func (d *DummyCardService) Remove(cardID uint64) (bool, error) {
	return true, nil
}
