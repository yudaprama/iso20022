package model

// Information related to the request of a deposit transaction from an ATM.
type ATMDepositRequest1 struct {

	// Environment in which the transaction is performed.
	Environment *ATMEnvironment11 `xml:"Envt"`

	// Context in which the deposit transaction is performed.
	Context *ATMContext10 `xml:"Cntxt"`

	// Deposit transaction for which the service is requested.
	Transaction *ATMTransaction15 `xml:"Tx"`
}

func (a *ATMDepositRequest1) AddEnvironment() *ATMEnvironment11 {
	a.Environment = new(ATMEnvironment11)
	return a.Environment
}

func (a *ATMDepositRequest1) AddContext() *ATMContext10 {
	a.Context = new(ATMContext10)
	return a.Context
}

func (a *ATMDepositRequest1) AddTransaction() *ATMTransaction15 {
	a.Transaction = new(ATMTransaction15)
	return a.Transaction
}
