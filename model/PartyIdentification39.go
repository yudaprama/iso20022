package model

// Identification and additional identification Information on the party.
type PartyIdentification39 struct {

	// Unique identification of the party.
	Identification *PartyIdentification9Choice `xml:"Id"`

	// Ancillary identification information about the party.
	AdditionalIdentificationInformation *PartyAdditionalIdentification2Choice `xml:"AddtlIdInf,omitempty"`
}

func (p *PartyIdentification39) AddIdentification() *PartyIdentification9Choice {
	p.Identification = new(PartyIdentification9Choice)
	return p.Identification
}

func (p *PartyIdentification39) AddAdditionalIdentificationInformation() *PartyAdditionalIdentification2Choice {
	p.AdditionalIdentificationInformation = new(PartyAdditionalIdentification2Choice)
	return p.AdditionalIdentificationInformation
}
