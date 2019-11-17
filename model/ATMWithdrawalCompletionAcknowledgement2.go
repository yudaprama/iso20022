package model

// Information related to the acknowledgement  of an ATM withdrawal from the ATM manager.
type ATMWithdrawalCompletionAcknowledgement2 struct {

	// ATM information.
	ATM *AutomatedTellerMachine3 `xml:"ATM"`

	// Context in which the transaction is performed.
	Context *ATMContext9 `xml:"Cntxt"`

	// Acknowledgement of the withdrawal completion advice.
	Transaction *ATMTransaction18 `xml:"Tx"`
}

func (a *ATMWithdrawalCompletionAcknowledgement2) AddATM() *AutomatedTellerMachine3 {
	a.ATM = new(AutomatedTellerMachine3)
	return a.ATM
}

func (a *ATMWithdrawalCompletionAcknowledgement2) AddContext() *ATMContext9 {
	a.Context = new(ATMContext9)
	return a.Context
}

func (a *ATMWithdrawalCompletionAcknowledgement2) AddTransaction() *ATMTransaction18 {
	a.Transaction = new(ATMTransaction18)
	return a.Transaction
}
