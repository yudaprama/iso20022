package model

// Choice of formats for the identification of a party.
type PartyIdentification90Choice struct {

	// Identification of the party expressed as a BIC.
	AnyBIC *AnyBICIdentifier `xml:"AnyBIC"`

	// Unique and unambiguous identifier, as assigned to the party using a proprietary identification scheme.
	ProprietaryIdentification *GenericIdentification1 `xml:"PrtryId"`

	// Name and address of the party.
	NameAndAddress *NameAndAddress5 `xml:"NmAndAdr"`
}

func (p *PartyIdentification90Choice) SetAnyBIC(value string) {
	p.AnyBIC = (*AnyBICIdentifier)(&value)
}

func (p *PartyIdentification90Choice) AddProprietaryIdentification() *GenericIdentification1 {
	p.ProprietaryIdentification = new(GenericIdentification1)
	return p.ProprietaryIdentification
}

func (p *PartyIdentification90Choice) AddNameAndAddress() *NameAndAddress5 {
	p.NameAndAddress = new(NameAndAddress5)
	return p.NameAndAddress
}
