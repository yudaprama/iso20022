package model

// Information related to the completion of a withdrawal on the ATM.
type ATMWithdrawalCompletionAdvice1 struct {

	// Environment of the withdrawal transaction.
	Environment *ATMEnvironment3 `xml:"Envt"`

	// Context in which the transaction is performed.
	Context *ATMContext2 `xml:"Cntxt"`

	// Withdrawal transaction for which the completion is sent.
	Transaction *ATMTransaction3 `xml:"Tx"`
}

func (a *ATMWithdrawalCompletionAdvice1) AddEnvironment() *ATMEnvironment3 {
	a.Environment = new(ATMEnvironment3)
	return a.Environment
}

func (a *ATMWithdrawalCompletionAdvice1) AddContext() *ATMContext2 {
	a.Context = new(ATMContext2)
	return a.Context
}

func (a *ATMWithdrawalCompletionAdvice1) AddTransaction() *ATMTransaction3 {
	a.Transaction = new(ATMTransaction3)
	return a.Transaction
}
