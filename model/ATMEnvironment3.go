package model

// Environment of the withdrawal transaction.
type ATMEnvironment3 struct {

	// Acquirer of the withdrawal transaction, in charge of the funds settlement with the issuer.
	Acquirer *Acquirer7 `xml:"Acqrr,omitempty"`

	// Identification of the ATM manager.
	ATMManagerIdentification *Max35Text `xml:"ATMMgrId,omitempty"`

	// Entity hosting the ATM terminal.
	HostingEntity *TerminalHosting1 `xml:"HstgNtty,omitempty"`

	// ATM information.
	ATM *AutomatedTellerMachine1 `xml:"ATM"`

	// Customer involved in the transaction.
	Customer *ATMCustomer3 `xml:"Cstmr"`

	// Card performing the transaction.
	Card *PaymentCard17 `xml:"Card,omitempty"`
}

func (a *ATMEnvironment3) AddAcquirer() *Acquirer7 {
	a.Acquirer = new(Acquirer7)
	return a.Acquirer
}

func (a *ATMEnvironment3) SetATMManagerIdentification(value string) {
	a.ATMManagerIdentification = (*Max35Text)(&value)
}

func (a *ATMEnvironment3) AddHostingEntity() *TerminalHosting1 {
	a.HostingEntity = new(TerminalHosting1)
	return a.HostingEntity
}

func (a *ATMEnvironment3) AddATM() *AutomatedTellerMachine1 {
	a.ATM = new(AutomatedTellerMachine1)
	return a.ATM
}

func (a *ATMEnvironment3) AddCustomer() *ATMCustomer3 {
	a.Customer = new(ATMCustomer3)
	return a.Customer
}

func (a *ATMEnvironment3) AddCard() *PaymentCard17 {
	a.Card = new(PaymentCard17)
	return a.Card
}
