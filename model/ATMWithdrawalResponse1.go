package model

// Information related to the response of an ATM withdrawal from an ATM manager.
type ATMWithdrawalResponse1 struct {

	// Environment of the withdrawal transaction.
	Environment *ATMEnvironment2 `xml:"Envt"`

	// Context in which the transaction is performed.
	Context *ATMContext2 `xml:"Cntxt"`

	// Response to the withdrawal transaction request.
	Transaction *ATMTransaction2 `xml:"Tx"`
}

func (a *ATMWithdrawalResponse1) AddEnvironment() *ATMEnvironment2 {
	a.Environment = new(ATMEnvironment2)
	return a.Environment
}

func (a *ATMWithdrawalResponse1) AddContext() *ATMContext2 {
	a.Context = new(ATMContext2)
	return a.Context
}

func (a *ATMWithdrawalResponse1) AddTransaction() *ATMTransaction2 {
	a.Transaction = new(ATMTransaction2)
	return a.Transaction
}
