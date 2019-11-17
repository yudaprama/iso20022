package model

// Information related to the request of a withdrawal from an ATM.
type ATMWithdrawalRequest2 struct {

	// Environment of the withdrawal transaction.
	Environment *ATMEnvironment11 `xml:"Envt"`

	// Context in which the transaction is performed.
	Context *ATMContext8 `xml:"Cntxt"`

	// Withdrawal transaction for which the authorisation is requested.
	Transaction *ATMTransaction13 `xml:"Tx"`
}

func (a *ATMWithdrawalRequest2) AddEnvironment() *ATMEnvironment11 {
	a.Environment = new(ATMEnvironment11)
	return a.Environment
}

func (a *ATMWithdrawalRequest2) AddContext() *ATMContext8 {
	a.Context = new(ATMContext8)
	return a.Context
}

func (a *ATMWithdrawalRequest2) AddTransaction() *ATMTransaction13 {
	a.Transaction = new(ATMTransaction13)
	return a.Transaction
}
