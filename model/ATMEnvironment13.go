package model

// Environment of the withdrawal transaction.
type ATMEnvironment13 struct {

	// Acquirer of the transactions, in charge of the funds settlement with the issuer.
	Acquirer *Acquirer7 `xml:"Acqrr,omitempty"`

	// Identification of the ATM manager.
	ATMManagerIdentification *Max35Text `xml:"ATMMgrId,omitempty"`

	// Entity hosting the ATM terminal.
	HostingEntity *TerminalHosting1 `xml:"HstgNtty,omitempty"`

	// ATM information.
	ATM *AutomatedTellerMachine9 `xml:"ATM"`

	// Customer involved in the transaction.
	Customer *ATMCustomer6 `xml:"Cstmr"`

	// Card performing the transaction.
	Card *PaymentCard23 `xml:"Card,omitempty"`
}

func (a *ATMEnvironment13) AddAcquirer() *Acquirer7 {
	a.Acquirer = new(Acquirer7)
	return a.Acquirer
}

func (a *ATMEnvironment13) SetATMManagerIdentification(value string) {
	a.ATMManagerIdentification = (*Max35Text)(&value)
}

func (a *ATMEnvironment13) AddHostingEntity() *TerminalHosting1 {
	a.HostingEntity = new(TerminalHosting1)
	return a.HostingEntity
}

func (a *ATMEnvironment13) AddATM() *AutomatedTellerMachine9 {
	a.ATM = new(AutomatedTellerMachine9)
	return a.ATM
}

func (a *ATMEnvironment13) AddCustomer() *ATMCustomer6 {
	a.Customer = new(ATMCustomer6)
	return a.Customer
}

func (a *ATMEnvironment13) AddCard() *PaymentCard23 {
	a.Card = new(PaymentCard23)
	return a.Card
}
