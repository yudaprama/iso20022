package model

// Environment of the transaction.
type ATMEnvironment7 struct {

	// Acquirer of the ATM transaction, in charge of the funds settlement with the issuer.
	Acquirer *Acquirer7 `xml:"Acqrr,omitempty"`

	// Institution in charge of managing the ATM.
	ATMManager *Acquirer8 `xml:"ATMMgr,omitempty"`

	// Entity hosting the ATM terminal.
	HostingEntity *TerminalHosting1 `xml:"HstgNtty,omitempty"`

	// ATM information.
	ATM *AutomatedTellerMachine3 `xml:"ATM"`
}

func (a *ATMEnvironment7) AddAcquirer() *Acquirer7 {
	a.Acquirer = new(Acquirer7)
	return a.Acquirer
}

func (a *ATMEnvironment7) AddATMManager() *Acquirer8 {
	a.ATMManager = new(Acquirer8)
	return a.ATMManager
}

func (a *ATMEnvironment7) AddHostingEntity() *TerminalHosting1 {
	a.HostingEntity = new(TerminalHosting1)
	return a.HostingEntity
}

func (a *ATMEnvironment7) AddATM() *AutomatedTellerMachine3 {
	a.ATM = new(AutomatedTellerMachine3)
	return a.ATM
}
