package model

// Information related to the acknowledgement  of an ATM deposit transaction from the ATM manager.
type ATMDepositCompletionAcknowledgement1 struct {

	// ATM information.
	ATM *AutomatedTellerMachine3 `xml:"ATM"`

	// Context in which the transaction is performed.
	Context *ATMContext12 `xml:"Cntxt"`

	// Acknowledgement of the deposit completion advice.
	Transaction *ATMTransaction18 `xml:"Tx"`
}

func (a *ATMDepositCompletionAcknowledgement1) AddATM() *AutomatedTellerMachine3 {
	a.ATM = new(AutomatedTellerMachine3)
	return a.ATM
}

func (a *ATMDepositCompletionAcknowledgement1) AddContext() *ATMContext12 {
	a.Context = new(ATMContext12)
	return a.Context
}

func (a *ATMDepositCompletionAcknowledgement1) AddTransaction() *ATMTransaction18 {
	a.Transaction = new(ATMTransaction18)
	return a.Transaction
}
