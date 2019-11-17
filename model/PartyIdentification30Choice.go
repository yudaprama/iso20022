package model

// Choice of identification of a party.
type PartyIdentification30Choice struct {

	// Unique and unambiguous identifier for an organisation that is allocated by an institution.
	BIC *BICIdentifier `xml:"BIC"`

	// Name by which a party is known and which is usually used to identify that party.
	NameAndAddress *NameAndAddress5 `xml:"NmAndAdr"`

	// Unique and unambiguous identifier, as assigned to a financial institution using a proprietary identification scheme.
	ProprietaryIdentification *GenericIdentification19 `xml:"PrtryId"`
}

func (p *PartyIdentification30Choice) SetBIC(value string) {
	p.BIC = (*BICIdentifier)(&value)
}

func (p *PartyIdentification30Choice) AddNameAndAddress() *NameAndAddress5 {
	p.NameAndAddress = new(NameAndAddress5)
	return p.NameAndAddress
}

func (p *PartyIdentification30Choice) AddProprietaryIdentification() *GenericIdentification19 {
	p.ProprietaryIdentification = new(GenericIdentification19)
	return p.ProprietaryIdentification
}
