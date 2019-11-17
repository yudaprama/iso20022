package model

// Environment of the withdrawal transaction.
type ATMEnvironment12 struct {

	// Acquirer of the transactions, in charge of the funds settlement with the issuer.
	Acquirer *Acquirer7 `xml:"Acqrr,omitempty"`

	// Institution in charge of managing the ATM.
	ATMManager *Acquirer8 `xml:"ATMMgr,omitempty"`

	// Entity hosting the ATM terminal.
	HostingEntity *TerminalHosting1 `xml:"HstgNtty,omitempty"`

	// ATM information.
	ATM *AutomatedTellerMachine2 `xml:"ATM"`

	// Customer involved in the transaction.
	Customer *ATMCustomer5 `xml:"Cstmr"`

	// Encryption of the sensitive card data.
	ProtectedCardData *ContentInformationType10 `xml:"PrtctdCardData,omitempty"`

	// Sensitive data associated with the card performing the transaction.
	PlainCardData *PlainCardData19 `xml:"PlainCardData,omitempty"`
}

func (a *ATMEnvironment12) AddAcquirer() *Acquirer7 {
	a.Acquirer = new(Acquirer7)
	return a.Acquirer
}

func (a *ATMEnvironment12) AddATMManager() *Acquirer8 {
	a.ATMManager = new(Acquirer8)
	return a.ATMManager
}

func (a *ATMEnvironment12) AddHostingEntity() *TerminalHosting1 {
	a.HostingEntity = new(TerminalHosting1)
	return a.HostingEntity
}

func (a *ATMEnvironment12) AddATM() *AutomatedTellerMachine2 {
	a.ATM = new(AutomatedTellerMachine2)
	return a.ATM
}

func (a *ATMEnvironment12) AddCustomer() *ATMCustomer5 {
	a.Customer = new(ATMCustomer5)
	return a.Customer
}

func (a *ATMEnvironment12) AddProtectedCardData() *ContentInformationType10 {
	a.ProtectedCardData = new(ContentInformationType10)
	return a.ProtectedCardData
}

func (a *ATMEnvironment12) AddPlainCardData() *PlainCardData19 {
	a.PlainCardData = new(PlainCardData19)
	return a.PlainCardData
}
