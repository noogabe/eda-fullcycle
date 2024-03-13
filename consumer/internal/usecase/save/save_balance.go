package save_balance

import "github.com.br/noogabe/eda-fullcycle/consumer/internal/domain/balance"

type SaveBalanceInputDto struct {
	AccountId string `json:"account_id"`
	Amount    int    `json:"amount"`
}

type SaveBalanceOutputDto struct {
	AccountId string
	Amount    int
}

type SaveBalanceUsecase struct {
	BalanceGateway balance.BalanceGateway
}

func NewSaveBalanceUsecase(balanceGateway balance.BalanceGateway) *SaveBalanceUsecase {
	return &SaveBalanceUsecase{
		BalanceGateway: balanceGateway,
	}
}

func (u *SaveBalanceUsecase) Execute(input SaveBalanceInputDto) (*SaveBalanceOutputDto, error) {

	aBalance := u.BalanceGateway.Get(input.AccountId)

	var err error

	if aBalance == nil {
		aBalance, err = balance.NewBalance(input.AccountId, input.Amount)

		if err != nil {
			return nil, err
		}

		err = u.BalanceGateway.Create(*aBalance)
	} else {
		err = aBalance.UpdateBalance(input.Amount)
		if err != nil {
			return nil, err
		}
		err = u.BalanceGateway.Update(*aBalance)
	}

	if err != nil {
		return nil, err
	}

	return &SaveBalanceOutputDto{
		AccountId: input.AccountId,
		Amount:    input.Amount,
	}, nil
}
