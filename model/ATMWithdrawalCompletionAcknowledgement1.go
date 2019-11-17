package model

// Information related to the acknowledgement  of an ATM withdrawal from the ATM manager.
type ATMWithdrawalCompletionAcknowledgement1 struct {

	// ATM information.
	ATM *AutomatedTellerMachine3 `xml:"ATM"`

	// Context in which the transaction is performed.
	Context *ATMContext2 `xml:"Cntxt"`

	// Acknowledgement of the withdrawal completion advice.
	Transaction *ATMTransaction4 `xml:"Tx"`
}

func (a *ATMWithdrawalCompletionAcknowledgement1) AddATM() *AutomatedTellerMachine3 {
	a.ATM = new(AutomatedTellerMachine3)
	return a.ATM
}

func (a *ATMWithdrawalCompletionAcknowledgement1) AddContext() *ATMContext2 {
	a.Context = new(ATMContext2)
	return a.Context
}

func (a *ATMWithdrawalCompletionAcknowledgement1) AddTransaction() *ATMTransaction4 {
	a.Transaction = new(ATMTransaction4)
	return a.Transaction
}
