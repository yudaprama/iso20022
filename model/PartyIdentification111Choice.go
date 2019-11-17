package model

// Unique and unambiguous way to identify an organisation.
type PartyIdentification111Choice struct {

	// Unique and unambiguous way to identify an organisation.
	AnyBIC *AnyBICIdentifier `xml:"AnyBIC"`

	// Unique and unambiguous way to identify an organisation.
	NameAndAddress *NameAndAddress12 `xml:"NmAndAdr"`

	// Unique and unambiguous way to identify an organisation.
	Country *CountryCode `xml:"Ctry"`
}

func (p *PartyIdentification111Choice) SetAnyBIC(value string) {
	p.AnyBIC = (*AnyBICIdentifier)(&value)
}

func (p *PartyIdentification111Choice) AddNameAndAddress() *NameAndAddress12 {
	p.NameAndAddress = new(NameAndAddress12)
	return p.NameAndAddress
}

func (p *PartyIdentification111Choice) SetCountry(value string) {
	p.Country = (*CountryCode)(&value)
}
