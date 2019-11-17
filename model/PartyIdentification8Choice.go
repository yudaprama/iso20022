package model

// Identification of a party by a choice between a BIC or a name and address.
type PartyIdentification8Choice struct {

	// Name and address and Alternative Identifier of a party.
	NameAndAddress *NameAndAddress8 `xml:"NmAndAdr"`

	// Identification of a party by a BIC and an Alternative Identifier.
	BICOrBEI *PartyIdentification22 `xml:"BICOrBEI"`
}

func (p *PartyIdentification8Choice) AddNameAndAddress() *NameAndAddress8 {
	p.NameAndAddress = new(NameAndAddress8)
	return p.NameAndAddress
}

func (p *PartyIdentification8Choice) AddBICOrBEI() *PartyIdentification22 {
	p.BICOrBEI = new(PartyIdentification22)
	return p.BICOrBEI
}
