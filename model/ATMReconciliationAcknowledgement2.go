package model

// Information related to the acknowledgement  of an ATM reconciliation from the ATM. manager.
type ATMReconciliationAcknowledgement2 struct {

	// ATM information.
	ATM *AutomatedTellerMachine3 `xml:"ATM"`

	// Information about the reconciliation response.
	Transaction *ATMTransaction26 `xml:"Tx"`
}

func (a *ATMReconciliationAcknowledgement2) AddATM() *AutomatedTellerMachine3 {
	a.ATM = new(AutomatedTellerMachine3)
	return a.ATM
}

func (a *ATMReconciliationAcknowledgement2) AddTransaction() *ATMTransaction26 {
	a.Transaction = new(ATMTransaction26)
	return a.Transaction
}
