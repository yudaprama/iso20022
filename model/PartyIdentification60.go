package model

// Identification of a party by fund name, name and address or an LEI.
type PartyIdentification60 struct {

	// Identification of a fund.
	FundIdentification *Max35Text `xml:"FndId"`

	// Identification of the party expressed as name and an optional address and an optional alternative identifier.
	NameAndAddress *NameAndAddress8 `xml:"NmAndAdr,omitempty"`

	// Identification of the Legal Entity Identifier. This is a code allocated to a party as described in ISO 17442 "Financial Services - Legal Entity Identifier (LEI)".
	LegalEntityIdentifier *LEIIdentifier `xml:"LglNttyIdr,omitempty"`
}

func (p *PartyIdentification60) SetFundIdentification(value string) {
	p.FundIdentification = (*Max35Text)(&value)
}

func (p *PartyIdentification60) AddNameAndAddress() *NameAndAddress8 {
	p.NameAndAddress = new(NameAndAddress8)
	return p.NameAndAddress
}

func (p *PartyIdentification60) SetLegalEntityIdentifier(value string) {
	p.LegalEntityIdentifier = (*LEIIdentifier)(&value)
}
