package model

// Choice of identification of a party.
type PartyIdentification37Choice struct {

	// Code allocated to a financial or non-financial institution by the ISO 9362 Registration Authority, as described in ISO 9362 "Banking - Banking telecommunication messages - Business identifier code (BIC)".
	AnyBIC *AnyBICIdentifier `xml:"AnyBIC"`

	// Unique and unambiguous identifier, as assigned to a financial institution using a proprietary identification scheme.
	ProprietaryIdentification *GenericIdentification19 `xml:"PrtryId"`

	// Name and address of a party.
	NameAndAddress *NameAndAddress5 `xml:"NmAndAdr"`

	// Country in which a person resides (the place of a person's home). In the case of a company, it is the country from which the affairs of that company are directed.
	Country *CountryCode `xml:"Ctry"`
}

func (p *PartyIdentification37Choice) SetAnyBIC(value string) {
	p.AnyBIC = (*AnyBICIdentifier)(&value)
}

func (p *PartyIdentification37Choice) AddProprietaryIdentification() *GenericIdentification19 {
	p.ProprietaryIdentification = new(GenericIdentification19)
	return p.ProprietaryIdentification
}

func (p *PartyIdentification37Choice) AddNameAndAddress() *NameAndAddress5 {
	p.NameAndAddress = new(NameAndAddress5)
	return p.NameAndAddress
}

func (p *PartyIdentification37Choice) SetCountry(value string) {
	p.Country = (*CountryCode)(&value)
}
