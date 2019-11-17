package model

// ATM information.
type AutomatedTellerMachine4 struct {

	// ATM terminal device identification for the acquirer and the issuer.
	Identification *Max35Text `xml:"Id"`

	// ATM terminal device identification for the ATM manager.
	AdditionalIdentification *Max35Text `xml:"AddtlId,omitempty"`

	// ATM terminal device identification for the branch.
	SequenceNumber *Max35Text `xml:"SeqNb,omitempty"`

	// Reference currency of the ATM.
	BaseCurrency *ActiveCurrencyCode `xml:"BaseCcy"`

	// Location of the ATM.
	Location *PostalAddress17 `xml:"Lctn,omitempty"`

	// Indicates the environment of the transaction.
	LocationCategory *TransactionEnvironment2Code `xml:"LctnCtgy,omitempty"`

	// Capabilities of the ATM terminal performing the transaction.
	Capabilities *PointOfInteractionCapabilities5 `xml:"Cpblties,omitempty"`

	// ATM terminal equipment.
	Equipment *ATMEquipment1 `xml:"Eqpmnt,omitempty"`

	// List of ATM devices out of service.
	AvailableDevice []*ATMDevice2Code `xml:"AvlblDvc,omitempty"`
}

func (a *AutomatedTellerMachine4) SetIdentification(value string) {
	a.Identification = (*Max35Text)(&value)
}

func (a *AutomatedTellerMachine4) SetAdditionalIdentification(value string) {
	a.AdditionalIdentification = (*Max35Text)(&value)
}

func (a *AutomatedTellerMachine4) SetSequenceNumber(value string) {
	a.SequenceNumber = (*Max35Text)(&value)
}

func (a *AutomatedTellerMachine4) SetBaseCurrency(value string) {
	a.BaseCurrency = (*ActiveCurrencyCode)(&value)
}

func (a *AutomatedTellerMachine4) AddLocation() *PostalAddress17 {
	a.Location = new(PostalAddress17)
	return a.Location
}

func (a *AutomatedTellerMachine4) SetLocationCategory(value string) {
	a.LocationCategory = (*TransactionEnvironment2Code)(&value)
}

func (a *AutomatedTellerMachine4) AddCapabilities() *PointOfInteractionCapabilities5 {
	a.Capabilities = new(PointOfInteractionCapabilities5)
	return a.Capabilities
}

func (a *AutomatedTellerMachine4) AddEquipment() *ATMEquipment1 {
	a.Equipment = new(ATMEquipment1)
	return a.Equipment
}

func (a *AutomatedTellerMachine4) AddAvailableDevice(value string) {
	a.AvailableDevice = append(a.AvailableDevice, (*ATMDevice2Code)(&value))
}
