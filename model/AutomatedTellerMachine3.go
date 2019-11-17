package model

// ATM information.
type AutomatedTellerMachine3 struct {

	// ATM terminal device identification for the acquirer and the issuer.
	Identification *Max35Text `xml:"Id"`

	// ATM terminal device identification for the ATM manager.
	AdditionalIdentification *Max35Text `xml:"AddtlId,omitempty"`

	// ATM terminal device identification for the branch.
	SequenceNumber *Max35Text `xml:"SeqNb,omitempty"`

	// Location of the ATM.
	Location *PostalAddress17 `xml:"Lctn,omitempty"`
}

func (a *AutomatedTellerMachine3) SetIdentification(value string) {
	a.Identification = (*Max35Text)(&value)
}

func (a *AutomatedTellerMachine3) SetAdditionalIdentification(value string) {
	a.AdditionalIdentification = (*Max35Text)(&value)
}

func (a *AutomatedTellerMachine3) SetSequenceNumber(value string) {
	a.SequenceNumber = (*Max35Text)(&value)
}

func (a *AutomatedTellerMachine3) AddLocation() *PostalAddress17 {
	a.Location = new(PostalAddress17)
	return a.Location
}
