package model

// Information related to the response of an ATM PIN Management from an ATM manager.
type ATMPINManagementResponse2 struct {

	// Environment of the PIN management transaction.
	Environment *ATMEnvironment12 `xml:"Envt"`

	// Context in which the transaction is performed.
	Context *ATMContext17 `xml:"Cntxt"`

	// Response to the PIN management transaction.
	Transaction *ATMTransaction22 `xml:"Tx"`
}

func (a *ATMPINManagementResponse2) AddEnvironment() *ATMEnvironment12 {
	a.Environment = new(ATMEnvironment12)
	return a.Environment
}

func (a *ATMPINManagementResponse2) AddContext() *ATMContext17 {
	a.Context = new(ATMContext17)
	return a.Context
}

func (a *ATMPINManagementResponse2) AddTransaction() *ATMTransaction22 {
	a.Transaction = new(ATMTransaction22)
	return a.Transaction
}
