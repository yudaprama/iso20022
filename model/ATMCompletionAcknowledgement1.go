package model

// Information related to the acknowledgement  of an ATM completion from the ATM manager.
type ATMCompletionAcknowledgement1 struct {

	// ATM information.
	ATM *AutomatedTellerMachine3 `xml:"ATM"`

	// Context in which the transaction is performed.
	Context *ATMContext3 `xml:"Cntxt"`

	// Acknowledgement of the completion advice.
	Transaction *ATMTransaction4 `xml:"Tx"`
}

func (a *ATMCompletionAcknowledgement1) AddATM() *AutomatedTellerMachine3 {
	a.ATM = new(AutomatedTellerMachine3)
	return a.ATM
}

func (a *ATMCompletionAcknowledgement1) AddContext() *ATMContext3 {
	a.Context = new(ATMContext3)
	return a.Context
}

func (a *ATMCompletionAcknowledgement1) AddTransaction() *ATMTransaction4 {
	a.Transaction = new(ATMTransaction4)
	return a.Transaction
}
