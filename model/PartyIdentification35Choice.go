package model

// Choice of identification of a party. The party can be identified by giving a BIC or a proprietary code.
type PartyIdentification35Choice struct {

	// Unique and unambiguous identifier for an organisation that is allocated by an institution, eg, Dun & Bradstreet Identification.
	BIC *AnyBICIdentifier `xml:"BIC"`

	// Unique and unambiguous identifier, as assigned to a financial institution using a proprietary identification scheme.
	ProprietaryIdentification *GenericIdentification29 `xml:"PrtryId"`
}

func (p *PartyIdentification35Choice) SetBIC(value string) {
	p.BIC = (*AnyBICIdentifier)(&value)
}

func (p *PartyIdentification35Choice) AddProprietaryIdentification() *GenericIdentification29 {
	p.ProprietaryIdentification = new(GenericIdentification29)
	return p.ProprietaryIdentification
}
