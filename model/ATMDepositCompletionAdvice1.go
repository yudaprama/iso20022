package model

// Information related to the completion of a deposit transaction on the ATM.
type ATMDepositCompletionAdvice1 struct {

	// Environment of the deposit transaction.
	Environment *ATMEnvironment13 `xml:"Envt"`

	// Context in which the transaction is performed.
	Context *ATMContext12 `xml:"Cntxt"`

	// Deposit transaction for which the completion is sent.
	Transaction *ATMTransaction19 `xml:"Tx"`
}

func (a *ATMDepositCompletionAdvice1) AddEnvironment() *ATMEnvironment13 {
	a.Environment = new(ATMEnvironment13)
	return a.Environment
}

func (a *ATMDepositCompletionAdvice1) AddContext() *ATMContext12 {
	a.Context = new(ATMContext12)
	return a.Context
}

func (a *ATMDepositCompletionAdvice1) AddTransaction() *ATMTransaction19 {
	a.Transaction = new(ATMTransaction19)
	return a.Transaction
}
