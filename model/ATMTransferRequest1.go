package model

// Information related to the request of a fund transfer from an ATM.
type ATMTransferRequest1 struct {

	// Environment in which the fund transfer is performed.
	Environment *ATMEnvironment11 `xml:"Envt"`

	// Context in which the fund transfer is performed.
	Context *ATMContext18 `xml:"Cntxt,omitempty"`

	// Transfer information for the transaction.
	Transaction *ATMTransaction23 `xml:"Tx"`
}

func (a *ATMTransferRequest1) AddEnvironment() *ATMEnvironment11 {
	a.Environment = new(ATMEnvironment11)
	return a.Environment
}

func (a *ATMTransferRequest1) AddContext() *ATMContext18 {
	a.Context = new(ATMContext18)
	return a.Context
}

func (a *ATMTransferRequest1) AddTransaction() *ATMTransaction23 {
	a.Transaction = new(ATMTransaction23)
	return a.Transaction
}
