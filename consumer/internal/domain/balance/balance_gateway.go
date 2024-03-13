package balance

type BalanceGateway interface {
	Get(accountId string) *Balance
	Create(balance Balance) error
	Update(balance Balance) error
}
