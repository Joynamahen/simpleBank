package model

type InvestmentAccount struct {
	BaseAccount
	Fund Fund
}

func (ia *InvestmentAccount) ApplyGrowth() {
	ia.Balance *= (1 + ia.Fund.Growth)
}
