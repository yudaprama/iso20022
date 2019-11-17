package model

// Environment of the inquiry.
type ATMEnvironment4 struct {

	// Acquirer of the withdrawal transaction, in charge of the funds settlement with the issuer.
	Acquirer *Acquirer7 `xml:"Acqrr,omitempty"`

	// Identification of the ATM manager.
	ATMManagerIdentification *Max35Text `xml:"ATMMgrId,omitempty"`

	// Entity hosting the ATM terminal.
	HostingEntity *TerminalHosting1 `xml:"HstgNtty,omitempty"`

	// ATM information.
	ATM *AutomatedTellerMachine4 `xml:"ATM"`

	// Customer involved in the withdrawal transaction.
	Customer *ATMCustomer1 `xml:"Cstmr,omitempty"`

	// Card performing the withdrawal transaction.
	Card *PaymentCard16 `xml:"Card,omitempty"`
}

func (a *ATMEnvironment4) AddAcquirer() *Acquirer7 {
	a.Acquirer = new(Acquirer7)
	return a.Acquirer
}

func (a *ATMEnvironment4) SetATMManagerIdentification(value string) {
	a.ATMManagerIdentification = (*Max35Text)(&value)
}

func (a *ATMEnvironment4) AddHostingEntity() *TerminalHosting1 {
	a.HostingEntity = new(TerminalHosting1)
	return a.HostingEntity
}

func (a *ATMEnvironment4) AddATM() *AutomatedTellerMachine4 {
	a.ATM = new(AutomatedTellerMachine4)
	return a.ATM
}

func (a *ATMEnvironment4) AddCustomer() *ATMCustomer1 {
	a.Customer = new(ATMCustomer1)
	return a.Customer
}

func (a *ATMEnvironment4) AddCard() *PaymentCard16 {
	a.Card = new(PaymentCard16)
	return a.Card
}
