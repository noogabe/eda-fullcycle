package get_balance

import (
	"errors"

	"github.com.br/noogabe/eda-fullcycle/consumer/internal/domain/balance"
)

type GetBalanceInputDto struct {
	AccountId string `json:"account_id"`
}

type GetBalanceOutputDto struct {
	AccountId string `json:"account_id"`
	Amount    int    `json:"amount"`
}

type GetBalanceUsecase struct {
	BalanceGateway balance.BalanceGateway
}

func NewGetBalanceUsecase(balanceGateway balance.BalanceGateway) *GetBalanceUsecase {
	return &GetBalanceUsecase{
		BalanceGateway: balanceGateway,
	}
}

func (u *GetBalanceUsecase) Execute(input GetBalanceInputDto) (*GetBalanceOutputDto, error) {
	balance := u.BalanceGateway.Get(input.AccountId)

	if balance == nil {
		return nil, errors.New("balance not found")
	}

	return &GetBalanceOutputDto{
		AccountId: balance.AccountId,
		Amount:    balance.Amount,
	}, nil
}
