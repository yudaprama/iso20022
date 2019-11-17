package model

// Information related to the response of an ATM withdrawal from an ATM manager.
type ATMWithdrawalResponse2 struct {

	// Environment of the withdrawal transaction.
	Environment *ATMEnvironment12 `xml:"Envt"`

	// Context in which the transaction is performed.
	Context *ATMContext9 `xml:"Cntxt"`

	// Response to the withdrawal transaction request.
	Transaction *ATMTransaction14 `xml:"Tx"`
}

func (a *ATMWithdrawalResponse2) AddEnvironment() *ATMEnvironment12 {
	a.Environment = new(ATMEnvironment12)
	return a.Environment
}

func (a *ATMWithdrawalResponse2) AddContext() *ATMContext9 {
	a.Context = new(ATMContext9)
	return a.Context
}

func (a *ATMWithdrawalResponse2) AddTransaction() *ATMTransaction14 {
	a.Transaction = new(ATMTransaction14)
	return a.Transaction
}
