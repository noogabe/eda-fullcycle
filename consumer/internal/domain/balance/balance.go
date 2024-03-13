package balance

import (
	"errors"
)

type Balance struct {
	AccountId string
	Amount    int
}

func NewBalance(accountId string, amount int) (*Balance, error) {
	balance := &Balance{
		AccountId: accountId,
		Amount:    amount,
	}

	err := balance.Validate()

	if err != nil {
		return nil, err
	}

	return balance, nil
}

func (b *Balance) Validate() error {
	if b.Amount < 0 {
		return errors.New("amount must be greater than zero")
	}

	return nil
}

func (b *Balance) UpdateBalance(amount int) error {
	b.Amount = amount

	return b.Validate()
}
