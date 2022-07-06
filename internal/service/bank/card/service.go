package card

import (
	"log"

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
	rlen := 0
	if cursor >= uint64(len(bank.AllCards)) {
		cursor = 0
	}

	if len(bank.AllCards) >= int(cursor+limit) {
		rlen = int(limit)
	} else {
		rlen = int(cursor+limit) - len(bank.AllCards) + 1
	}
	log.Printf("rlen: %v", rlen)

	res := make([]bank.Card, rlen)
	for i := 0; i < rlen; i++ {
		res[i] = bank.AllCards[int(cursor)+i]
	}
	return res, nil
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
