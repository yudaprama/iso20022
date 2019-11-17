package model

// ATM information.
type AutomatedTellerMachine6 struct {

	// ATM terminal device identification for the acquirer and the issuer.
	Identification *Max35Text `xml:"Id"`

	// ATM terminal device identification for the ATM manager.
	AdditionalIdentification *Max35Text `xml:"AddtlId,omitempty"`

	// ATM terminal device identification for the branch.
	SequenceNumber *Max35Text `xml:"SeqNb,omitempty"`

	// Location of the ATM.
	Location *PostalAddress17 `xml:"Lctn,omitempty"`

	// Indicates the environment of the transaction.
	LocationCategory *TransactionEnvironment2Code `xml:"LctnCtgy,omitempty"`

	// ATM terminal equipment.
	Equipment *ATMEquipment1 `xml:"Eqpmnt,omitempty"`
}

func (a *AutomatedTellerMachine6) SetIdentification(value string) {
	a.Identification = (*Max35Text)(&value)
}

func (a *AutomatedTellerMachine6) SetAdditionalIdentification(value string) {
	a.AdditionalIdentification = (*Max35Text)(&value)
}

func (a *AutomatedTellerMachine6) SetSequenceNumber(value string) {
	a.SequenceNumber = (*Max35Text)(&value)
}

func (a *AutomatedTellerMachine6) AddLocation() *PostalAddress17 {
	a.Location = new(PostalAddress17)
	return a.Location
}

func (a *AutomatedTellerMachine6) SetLocationCategory(value string) {
	a.LocationCategory = (*TransactionEnvironment2Code)(&value)
}

func (a *AutomatedTellerMachine6) AddEquipment() *ATMEquipment1 {
	a.Equipment = new(ATMEquipment1)
	return a.Equipment
}
