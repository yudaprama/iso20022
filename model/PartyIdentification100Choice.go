package model

// Identification of a party.
type PartyIdentification100Choice struct {

	// Unique and unambiguous identifier for an organisation that is allocated by an institution, for example, Dun & Bradstreet Identification.
	AnyBIC *AnyBICIdentifier `xml:"AnyBIC"`

	// Unique and unambiguous identifier, as assigned to a financial institution using a proprietary identification scheme.
	ProprietaryIdentification *GenericIdentification36 `xml:"PrtryId"`

	// Name by which a party is known and which is usually used to identify that party.
	NameAndAddress *NameAndAddress6 `xml:"NmAndAdr"`
}

func (p *PartyIdentification100Choice) SetAnyBIC(value string) {
	p.AnyBIC = (*AnyBICIdentifier)(&value)
}

func (p *PartyIdentification100Choice) AddProprietaryIdentification() *GenericIdentification36 {
	p.ProprietaryIdentification = new(GenericIdentification36)
	return p.ProprietaryIdentification
}

func (p *PartyIdentification100Choice) AddNameAndAddress() *NameAndAddress6 {
	p.NameAndAddress = new(NameAndAddress6)
	return p.NameAndAddress
}
