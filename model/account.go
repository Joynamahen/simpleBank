package model

type Account interface {
	Deposit(amount float64)
	Withdraw(amount float64) bool
	GetBalance() float64
}

type BaseAccount struct {
	Balance float64
}

func (a *BaseAccount) Deposit(amount float64) {
	a.Balance += amount
}

func (a *BaseAccount) Withdraw(amount float64) bool {
	if amount > a.Balance {
		return false
	}
	a.Balance -= amount
	return true
}

func (a *BaseAccount) GetBalance() float64 {
	return a.Balance
}
