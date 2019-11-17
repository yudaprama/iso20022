package model

// Environment of the transaction.
type ATMEnvironment10 struct {

	// Acquirer of the ATM transaction, in charge of the funds settlement with the issuer.
	Acquirer *Acquirer7 `xml:"Acqrr,omitempty"`

	// Identification of the ATM manager.
	ATMManagerIdentification *Max35Text `xml:"ATMMgrId,omitempty"`

	// Entity hosting the ATM terminal.
	HostingEntity *TerminalHosting1 `xml:"HstgNtty,omitempty"`

	// ATM information.
	ATM *AutomatedTellerMachine8 `xml:"ATM"`
}

func (a *ATMEnvironment10) AddAcquirer() *Acquirer7 {
	a.Acquirer = new(Acquirer7)
	return a.Acquirer
}

func (a *ATMEnvironment10) SetATMManagerIdentification(value string) {
	a.ATMManagerIdentification = (*Max35Text)(&value)
}

func (a *ATMEnvironment10) AddHostingEntity() *TerminalHosting1 {
	a.HostingEntity = new(TerminalHosting1)
	return a.HostingEntity
}

func (a *ATMEnvironment10) AddATM() *AutomatedTellerMachine8 {
	a.ATM = new(AutomatedTellerMachine8)
	return a.ATM
}
