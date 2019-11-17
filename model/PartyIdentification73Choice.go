package model

// Identification of a party by a choice between a BIC or a name and address or an LEI.
type PartyIdentification73Choice struct {

	// Identification of the party expressed as name and an optional address and an optional alternative identifier.
	NameAndAddress *NameAndAddress8 `xml:"NmAndAdr"`

	// Identification of the party expressed as a BIC and an optional alternative identifier.
	AnyBIC *PartyIdentification44 `xml:"AnyBIC"`

	// Party Identification specified as a list of values per element
	PartyIdentification *PartyIdentification59 `xml:"PtyId"`
}

func (p *PartyIdentification73Choice) AddNameAndAddress() *NameAndAddress8 {
	p.NameAndAddress = new(NameAndAddress8)
	return p.NameAndAddress
}

func (p *PartyIdentification73Choice) AddAnyBIC() *PartyIdentification44 {
	p.AnyBIC = new(PartyIdentification44)
	return p.AnyBIC
}

func (p *PartyIdentification73Choice) AddPartyIdentification() *PartyIdentification59 {
	p.PartyIdentification = new(PartyIdentification59)
	return p.PartyIdentification
}
