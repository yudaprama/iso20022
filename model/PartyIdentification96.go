package model

// Identification of a person or an organisation.
type PartyIdentification96 struct {

	// Identification of the organisation.
	Identification *PartyIdentification96Choice `xml:"Id,omitempty"`

	// Identification of the organisation with a Legal Entity Identifier. This is a code allocated to a party as described in ISO 17442 "Financial Services - Legal Entity Identifier (LEI)".
	LegalEntityIdentifier *LEIIdentifier `xml:"LglNttyIdr,omitempty"`
}

func (p *PartyIdentification96) AddIdentification() *PartyIdentification96Choice {
	p.Identification = new(PartyIdentification96Choice)
	return p.Identification
}

func (p *PartyIdentification96) SetLegalEntityIdentifier(value string) {
	p.LegalEntityIdentifier = (*LEIIdentifier)(&value)
}
