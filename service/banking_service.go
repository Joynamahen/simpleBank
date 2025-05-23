package service

import (
	"fmt"
	"simpleBank/model"
)

func StartBankingService() {
	user := model.User{
		Name:    "John Doe",
		Account: &model.SavingsAccount{BaseAccount: model.BaseAccount{Balance: 1000}, InterestRate: 0.05},
	}

	fmt.Printf("User: %s, Balance: %.2f\n", user.Name, user.Account.GetBalance())
	user.Account.Deposit(500)
	fmt.Printf("Deposited 500. New Balance: %.2f\n", user.Account.GetBalance())
}
