package model

// Information related to the response of an ATM transfer from an ATM manager.
type ATMTransferResponse1 struct {

	// Environment of the transaction.
	Environment *ATMEnvironment12 `xml:"Envt"`

	// Context in which the transfer is performed.
	Context *ATMContext19 `xml:"Cntxt"`

	// Transfer information for the transaction.
	Transaction *ATMTransaction24 `xml:"Tx"`
}

func (a *ATMTransferResponse1) AddEnvironment() *ATMEnvironment12 {
	a.Environment = new(ATMEnvironment12)
	return a.Environment
}

func (a *ATMTransferResponse1) AddContext() *ATMContext19 {
	a.Context = new(ATMContext19)
	return a.Context
}

func (a *ATMTransferResponse1) AddTransaction() *ATMTransaction24 {
	a.Transaction = new(ATMTransaction24)
	return a.Transaction
}
