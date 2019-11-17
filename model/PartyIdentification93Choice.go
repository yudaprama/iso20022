package model

// Choice of identification of a party.
type PartyIdentification93Choice struct {

	// Code allocated to a financial or non-financial institution by the ISO 9362 Registration Authority, as described in ISO 9362 "Banking - Banking telecommunication messages - Business identifier code (BIC)".
	AnyBIC *AnyBICIdentifier `xml:"AnyBIC"`

	// Unique and unambiguous identifier, as assigned to a financial institution using a proprietary identification scheme.
	ProprietaryIdentification *GenericIdentification36 `xml:"PrtryId"`

	// Name and address of a party.
	NameAndAddress *NameAndAddress5 `xml:"NmAndAdr"`

	// Country in which a person resides (the place of a person's home). In the case of a company, it is the country from which the affairs of that company are directed.
	Country *CountryCode `xml:"Ctry"`
}

func (p *PartyIdentification93Choice) SetAnyBIC(value string) {
	p.AnyBIC = (*AnyBICIdentifier)(&value)
}

func (p *PartyIdentification93Choice) AddProprietaryIdentification() *GenericIdentification36 {
	p.ProprietaryIdentification = new(GenericIdentification36)
	return p.ProprietaryIdentification
}

func (p *PartyIdentification93Choice) AddNameAndAddress() *NameAndAddress5 {
	p.NameAndAddress = new(NameAndAddress5)
	return p.NameAndAddress
}

func (p *PartyIdentification93Choice) SetCountry(value string) {
	p.Country = (*CountryCode)(&value)
}
