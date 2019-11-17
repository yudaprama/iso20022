package model

// Identification of a party.
type PartyIdentification1Choice struct {

	// Unique and unambiguous identifier for an organisation that is allocated by an institution.
	BICOrBEI *AnyBICIdentifier `xml:"BICOrBEI"`

	// Unique and unambiguous identifier, as assigned to a financial institution using a proprietary identification scheme.
	ProprietaryIdentification *GenericIdentification1 `xml:"PrtryId"`

	// Name by which a party is known and which is usually used to identify that party.
	NameAndAddress *NameAndAddress2 `xml:"NmAndAdr"`
}

func (p *PartyIdentification1Choice) SetBICOrBEI(value string) {
	p.BICOrBEI = (*AnyBICIdentifier)(&value)
}

func (p *PartyIdentification1Choice) AddProprietaryIdentification() *GenericIdentification1 {
	p.ProprietaryIdentification = new(GenericIdentification1)
	return p.ProprietaryIdentification
}

func (p *PartyIdentification1Choice) AddNameAndAddress() *NameAndAddress2 {
	p.NameAndAddress = new(NameAndAddress2)
	return p.NameAndAddress
}
