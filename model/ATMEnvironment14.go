package model

// Environment of the inquiry.
type ATMEnvironment14 struct {

	// Acquirer of the transactions, in charge of the funds settlement with the issuer.
	Acquirer *Acquirer7 `xml:"Acqrr,omitempty"`

	// Identification of the ATM manager.
	ATMManagerIdentification *Max35Text `xml:"ATMMgrId,omitempty"`

	// Entity hosting the ATM terminal.
	HostingEntity *TerminalHosting1 `xml:"HstgNtty,omitempty"`

	// ATM information.
	ATM *AutomatedTellerMachine10 `xml:"ATM"`

	// Customer involved in the withdrawal transaction.
	Customer *ATMCustomer4 `xml:"Cstmr,omitempty"`

	// Card performing the withdrawal transaction.
	Card *PaymentCard22 `xml:"Card,omitempty"`
}

func (a *ATMEnvironment14) AddAcquirer() *Acquirer7 {
	a.Acquirer = new(Acquirer7)
	return a.Acquirer
}

func (a *ATMEnvironment14) SetATMManagerIdentification(value string) {
	a.ATMManagerIdentification = (*Max35Text)(&value)
}

func (a *ATMEnvironment14) AddHostingEntity() *TerminalHosting1 {
	a.HostingEntity = new(TerminalHosting1)
	return a.HostingEntity
}

func (a *ATMEnvironment14) AddATM() *AutomatedTellerMachine10 {
	a.ATM = new(AutomatedTellerMachine10)
	return a.ATM
}

func (a *ATMEnvironment14) AddCustomer() *ATMCustomer4 {
	a.Customer = new(ATMCustomer4)
	return a.Customer
}

func (a *ATMEnvironment14) AddCard() *PaymentCard22 {
	a.Card = new(PaymentCard22)
	return a.Card
}
