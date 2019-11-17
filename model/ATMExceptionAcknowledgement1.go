package model

// Information related to the acknowledgement  of an ATM exception.
type ATMExceptionAcknowledgement1 struct {

	// ATM information.
	ATM *AutomatedTellerMachine3 `xml:"ATM"`

	// Context in which the transaction is performed, if any.
	Context *ATMContext20 `xml:"Cntxt"`

	// Acknowledgement of the exception advice.
	Transaction *ATMTransaction28 `xml:"Tx"`
}

func (a *ATMExceptionAcknowledgement1) AddATM() *AutomatedTellerMachine3 {
	a.ATM = new(AutomatedTellerMachine3)
	return a.ATM
}

func (a *ATMExceptionAcknowledgement1) AddContext() *ATMContext20 {
	a.Context = new(ATMContext20)
	return a.Context
}

func (a *ATMExceptionAcknowledgement1) AddTransaction() *ATMTransaction28 {
	a.Transaction = new(ATMTransaction28)
	return a.Transaction
}
