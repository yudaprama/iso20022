package model

// Environment of the transaction.
type ATMEnvironment5 struct {

	// ATM information.
	ATM *AutomatedTellerMachine2 `xml:"ATM"`

	// Customer involved in the withdrawal transaction.
	Customer *ATMCustomer2 `xml:"Cstmr"`

	// Encryption of the sensitive card data.
	ProtectedCardData *ContentInformationType10 `xml:"PrtctdCardData,omitempty"`

	// Sensitive data associated with the card performing the transaction.
	PlainCardData *PlainCardData14 `xml:"PlainCardData,omitempty"`
}

func (a *ATMEnvironment5) AddATM() *AutomatedTellerMachine2 {
	a.ATM = new(AutomatedTellerMachine2)
	return a.ATM
}

func (a *ATMEnvironment5) AddCustomer() *ATMCustomer2 {
	a.Customer = new(ATMCustomer2)
	return a.Customer
}

func (a *ATMEnvironment5) AddProtectedCardData() *ContentInformationType10 {
	a.ProtectedCardData = new(ContentInformationType10)
	return a.ProtectedCardData
}

func (a *ATMEnvironment5) AddPlainCardData() *PlainCardData14 {
	a.PlainCardData = new(PlainCardData14)
	return a.PlainCardData
}
