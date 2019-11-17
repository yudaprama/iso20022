package model

// Identification of a party.
type PartyIdentification33Choice struct {

	// Unique and unambiguous identifier for an organisation that is allocated by an institution, eg, Dun & Bradstreet Identification.
	AnyBIC *AnyBICIdentifier `xml:"AnyBIC"`

	// Unique and unambiguous identifier, as assigned to a financial institution using a proprietary identification scheme.
	ProprietaryIdentification *GenericIdentification29 `xml:"PrtryId"`

	// Name by which a party is known and which is usually used to identify that party.
	NameAndAddress *NameAndAddress6 `xml:"NmAndAdr"`
}

func (p *PartyIdentification33Choice) SetAnyBIC(value string) {
	p.AnyBIC = (*AnyBICIdentifier)(&value)
}

func (p *PartyIdentification33Choice) AddProprietaryIdentification() *GenericIdentification29 {
	p.ProprietaryIdentification = new(GenericIdentification29)
	return p.ProprietaryIdentification
}

func (p *PartyIdentification33Choice) AddNameAndAddress() *NameAndAddress6 {
	p.NameAndAddress = new(NameAndAddress6)
	return p.NameAndAddress
}
