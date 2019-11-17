package model

// Information related to the reconciliation of an ATM.
type ATMReconciliationAdvice1 struct {

	// Environment of the ATM.
	Environment *ATMEnvironment10 `xml:"Envt"`

	// Command result for reinitialisation of the transaction counters.
	CommandResult []*ATMCommand5 `xml:"CmdRslt,omitempty"`

	// Party which has requested the reconciliation.
	CommandContext *ATMCommand6 `xml:"CmdCntxt,omitempty"`

	// Information about the reconciliation request.
	Transaction *ATMTransaction11 `xml:"Tx"`
}

func (a *ATMReconciliationAdvice1) AddEnvironment() *ATMEnvironment10 {
	a.Environment = new(ATMEnvironment10)
	return a.Environment
}

func (a *ATMReconciliationAdvice1) AddCommandResult() *ATMCommand5 {
	newValue := new(ATMCommand5)
	a.CommandResult = append(a.CommandResult, newValue)
	return newValue
}

func (a *ATMReconciliationAdvice1) AddCommandContext() *ATMCommand6 {
	a.CommandContext = new(ATMCommand6)
	return a.CommandContext
}

func (a *ATMReconciliationAdvice1) AddTransaction() *ATMTransaction11 {
	a.Transaction = new(ATMTransaction11)
	return a.Transaction
}
