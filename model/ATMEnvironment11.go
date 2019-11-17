package model

// Environment of the withdrawal transaction.
type ATMEnvironment11 struct {

	// Acquirer of the transaction, in charge of the funds settlement with the issuer.
	Acquirer *Acquirer7 `xml:"Acqrr,omitempty"`

	// Identification of the ATM manager.
	ATMManagerIdentification *Max35Text `xml:"ATMMgrId,omitempty"`

	// Entity hosting the ATM terminal.
	HostingEntity *TerminalHosting1 `xml:"HstgNtty,omitempty"`

	// ATM information.
	ATM *AutomatedTellerMachine9 `xml:"ATM"`

	// Customer involved in the transaction.
	Customer *ATMCustomer4 `xml:"Cstmr"`

	// Card performing the transaction.
	Card *PaymentCard22 `xml:"Card,omitempty"`
}

func (a *ATMEnvironment11) AddAcquirer() *Acquirer7 {
	a.Acquirer = new(Acquirer7)
	return a.Acquirer
}

func (a *ATMEnvironment11) SetATMManagerIdentification(value string) {
	a.ATMManagerIdentification = (*Max35Text)(&value)
}

func (a *ATMEnvironment11) AddHostingEntity() *TerminalHosting1 {
	a.HostingEntity = new(TerminalHosting1)
	return a.HostingEntity
}

func (a *ATMEnvironment11) AddATM() *AutomatedTellerMachine9 {
	a.ATM = new(AutomatedTellerMachine9)
	return a.ATM
}

func (a *ATMEnvironment11) AddCustomer() *ATMCustomer4 {
	a.Customer = new(ATMCustomer4)
	return a.Customer
}

func (a *ATMEnvironment11) AddCard() *PaymentCard22 {
	a.Card = new(PaymentCard22)
	return a.Card
}
