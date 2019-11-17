package model

// Information related to the completion of a withdrawal on the ATM.
type ATMWithdrawalCompletionAdvice2 struct {

	// Environment of the withdrawal transaction.
	Environment *ATMEnvironment13 `xml:"Envt"`

	// Context in which the transaction is performed.
	Context *ATMContext9 `xml:"Cntxt"`

	// Withdrawal transaction for which the completion is sent.
	Transaction *ATMTransaction17 `xml:"Tx"`
}

func (a *ATMWithdrawalCompletionAdvice2) AddEnvironment() *ATMEnvironment13 {
	a.Environment = new(ATMEnvironment13)
	return a.Environment
}

func (a *ATMWithdrawalCompletionAdvice2) AddContext() *ATMContext9 {
	a.Context = new(ATMContext9)
	return a.Context
}

func (a *ATMWithdrawalCompletionAdvice2) AddTransaction() *ATMTransaction17 {
	a.Transaction = new(ATMTransaction17)
	return a.Transaction
}
