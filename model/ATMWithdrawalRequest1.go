package model

// Information related to the request of a withdrawal from an ATM.
type ATMWithdrawalRequest1 struct {

	// Environment of the withdrawal transaction.
	Environment *ATMEnvironment1 `xml:"Envt"`

	// Context in which the transaction is performed.
	Context *ATMContext1 `xml:"Cntxt"`

	// Withdrawal transaction for which the authorisation is requested.
	Transaction *ATMTransaction1 `xml:"Tx"`
}

func (a *ATMWithdrawalRequest1) AddEnvironment() *ATMEnvironment1 {
	a.Environment = new(ATMEnvironment1)
	return a.Environment
}

func (a *ATMWithdrawalRequest1) AddContext() *ATMContext1 {
	a.Context = new(ATMContext1)
	return a.Context
}

func (a *ATMWithdrawalRequest1) AddTransaction() *ATMTransaction1 {
	a.Transaction = new(ATMTransaction1)
	return a.Transaction
}
