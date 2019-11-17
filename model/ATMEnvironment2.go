package model

// Environment of the withdrawal transaction.
type ATMEnvironment2 struct {

	// Acquirer of the withdrawal transaction, in charge of the funds settlement with the issuer.
	Acquirer *Acquirer7 `xml:"Acqrr,omitempty"`

	// Institution in charge of managing the ATM.
	ATMManager *Acquirer8 `xml:"ATMMgr,omitempty"`

	// Entity hosting the ATM terminal.
	HostingEntity *TerminalHosting1 `xml:"HstgNtty,omitempty"`

	// ATM information.
	ATM *AutomatedTellerMachine2 `xml:"ATM"`

	// Customer involved in the withdrawal transaction.
	Customer *ATMCustomer2 `xml:"Cstmr"`

	// Encryption of the sensitive card data.
	ProtectedCardData *ContentInformationType10 `xml:"PrtctdCardData,omitempty"`

	// Sensitive data associated with the card performing the transaction.
	PlainCardData *PlainCardData14 `xml:"PlainCardData,omitempty"`
}

func (a *ATMEnvironment2) AddAcquirer() *Acquirer7 {
	a.Acquirer = new(Acquirer7)
	return a.Acquirer
}

func (a *ATMEnvironment2) AddATMManager() *Acquirer8 {
	a.ATMManager = new(Acquirer8)
	return a.ATMManager
}

func (a *ATMEnvironment2) AddHostingEntity() *TerminalHosting1 {
	a.HostingEntity = new(TerminalHosting1)
	return a.HostingEntity
}

func (a *ATMEnvironment2) AddATM() *AutomatedTellerMachine2 {
	a.ATM = new(AutomatedTellerMachine2)
	return a.ATM
}

func (a *ATMEnvironment2) AddCustomer() *ATMCustomer2 {
	a.Customer = new(ATMCustomer2)
	return a.Customer
}

func (a *ATMEnvironment2) AddProtectedCardData() *ContentInformationType10 {
	a.ProtectedCardData = new(ContentInformationType10)
	return a.ProtectedCardData
}

func (a *ATMEnvironment2) AddPlainCardData() *PlainCardData14 {
	a.PlainCardData = new(PlainCardData14)
	return a.PlainCardData
}
