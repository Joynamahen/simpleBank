package model

type SavingsAccount struct {
	BaseAccount
	InterestRate float64
}

func (sa *SavingsAccount) AddInterest() {
	sa.Balance += sa.Balance * sa.InterestRate
}
