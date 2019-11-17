package model

// Response to a deposit request.
type ATMDepositResponse1 struct {

	// Environment of the deposit transaction.
	Environment *ATMEnvironment12 `xml:"Envt"`

	// Context in which the transaction is performed.
	Context *ATMContext11 `xml:"Cntxt"`

	// Response to the deposit request.
	Transaction *ATMTransaction16 `xml:"Tx"`
}

func (a *ATMDepositResponse1) AddEnvironment() *ATMEnvironment12 {
	a.Environment = new(ATMEnvironment12)
	return a.Environment
}

func (a *ATMDepositResponse1) AddContext() *ATMContext11 {
	a.Context = new(ATMContext11)
	return a.Context
}

func (a *ATMDepositResponse1) AddTransaction() *ATMTransaction16 {
	a.Transaction = new(ATMTransaction16)
	return a.Transaction
}
