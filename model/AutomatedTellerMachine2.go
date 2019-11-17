package model

// ATM information.
type AutomatedTellerMachine2 struct {

	// ATM terminal device identification for the acquirer and the issuer.
	Identification *Max35Text `xml:"Id"`

	// ATM terminal device identification for the ATM manager.
	AdditionalIdentification *Max35Text `xml:"AddtlId,omitempty"`

	// ATM terminal device identification for the branch.
	SequenceNumber *Max35Text `xml:"SeqNb,omitempty"`

	// Reference currency of the ATM.
	BaseCurrency *ActiveCurrencyCode `xml:"BaseCcy,omitempty"`

	// Location of the ATM.
	Location *PostalAddress17 `xml:"Lctn,omitempty"`
}

func (a *AutomatedTellerMachine2) SetIdentification(value string) {
	a.Identification = (*Max35Text)(&value)
}

func (a *AutomatedTellerMachine2) SetAdditionalIdentification(value string) {
	a.AdditionalIdentification = (*Max35Text)(&value)
}

func (a *AutomatedTellerMachine2) SetSequenceNumber(value string) {
	a.SequenceNumber = (*Max35Text)(&value)
}

func (a *AutomatedTellerMachine2) SetBaseCurrency(value string) {
	a.BaseCurrency = (*ActiveCurrencyCode)(&value)
}

func (a *AutomatedTellerMachine2) AddLocation() *PostalAddress17 {
	a.Location = new(PostalAddress17)
	return a.Location
}
