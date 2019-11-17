package model

// Set of elements used to identify a person or an organisation.
type PartyIdentification71 struct {

	// Unique identification of the party.
	Identification *PartyIdentification40Choice `xml:"Id"`

	// Ancillary identification information about the party.
	AdditionalIdentificationInformation *PartyAdditionalIdentification2Choice `xml:"AddtlIdInf,omitempty"`
}

func (p *PartyIdentification71) AddIdentification() *PartyIdentification40Choice {
	p.Identification = new(PartyIdentification40Choice)
	return p.Identification
}

func (p *PartyIdentification71) AddAdditionalIdentificationInformation() *PartyAdditionalIdentification2Choice {
	p.AdditionalIdentificationInformation = new(PartyAdditionalIdentification2Choice)
	return p.AdditionalIdentificationInformation
}
