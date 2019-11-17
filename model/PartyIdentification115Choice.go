package model

// Choice of identification of a party.
type PartyIdentification115Choice struct {

	// Code allocated to a financial or non-financial institution by the ISO 9362 Registration Authority, as described in ISO 9362 "Banking - Banking telecommunication messages - Business identifier code (BIC)".
	AnyBIC *AnyBICIdentifier `xml:"AnyBIC"`

	// Unique and unambiguous identifier, as assigned to a financial institution using a proprietary identification scheme.
	ProprietaryIdentification *GenericIdentification84 `xml:"PrtryId"`

	// Name and address of a party.
	NameAndAddress *NameAndAddress12 `xml:"NmAndAdr"`

	// Country in which a person resides (the place of a person's home). In the case of a company, it is the country from which the affairs of that company are directed.
	Country *CountryCode `xml:"Ctry"`
}

func (p *PartyIdentification115Choice) SetAnyBIC(value string) {
	p.AnyBIC = (*AnyBICIdentifier)(&value)
}

func (p *PartyIdentification115Choice) AddProprietaryIdentification() *GenericIdentification84 {
	p.ProprietaryIdentification = new(GenericIdentification84)
	return p.ProprietaryIdentification
}

func (p *PartyIdentification115Choice) AddNameAndAddress() *NameAndAddress12 {
	p.NameAndAddress = new(NameAndAddress12)
	return p.NameAndAddress
}

func (p *PartyIdentification115Choice) SetCountry(value string) {
	p.Country = (*CountryCode)(&value)
}
