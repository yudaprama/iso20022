package model

// Identification of a party by a choice between a BIC or a name and address.
type PartyIdentification99Choice struct {

	// Name and address and Alternative Identifier of a party.
	NameAndAddress *NameAndAddress8 `xml:"NmAndAdr"`

	// Identification of a party by a BIC and an Alternative Identifier.
	AnyBIC *PartyIdentification44 `xml:"AnyBIC"`
}

func (p *PartyIdentification99Choice) AddNameAndAddress() *NameAndAddress8 {
	p.NameAndAddress = new(NameAndAddress8)
	return p.NameAndAddress
}

func (p *PartyIdentification99Choice) AddAnyBIC() *PartyIdentification44 {
	p.AnyBIC = new(PartyIdentification44)
	return p.AnyBIC
}
