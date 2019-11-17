package model

// Identification of a person or an organisation.
type PartyIdentification95 struct {

	// Identifier for an organisation that is allocated by an institution.
	Identification *PartyIdentification70Choice `xml:"Id,omitempty"`

	// Identification of the organisation with a Legal Entity Identifier. This is a code allocated to a party as described in ISO 17442 "Financial Services - Legal Entity Identifier (LEI)".
	LegalEntityIdentifier *LEIIdentifier `xml:"LglNttyIdr,omitempty"`
}

func (p *PartyIdentification95) AddIdentification() *PartyIdentification70Choice {
	p.Identification = new(PartyIdentification70Choice)
	return p.Identification
}

func (p *PartyIdentification95) SetLegalEntityIdentifier(value string) {
	p.LegalEntityIdentifier = (*LEIIdentifier)(&value)
}
