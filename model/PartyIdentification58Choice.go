package model

// Unique and unambiguous way to identify an organisation.
type PartyIdentification58Choice struct {

	// Unique and unambiguous way to identify an organisation.
	AnyBIC *AnyBICIdentifier `xml:"AnyBIC"`

	// Unique and unambiguous way to identify an organisation.
	NameAndAddress *NameAndAddress12 `xml:"NmAndAdr"`

	// Unique and unambiguous way to identify an organisation.
	Country *CountryCode `xml:"Ctry"`
}

func (p *PartyIdentification58Choice) SetAnyBIC(value string) {
	p.AnyBIC = (*AnyBICIdentifier)(&value)
}

func (p *PartyIdentification58Choice) AddNameAndAddress() *NameAndAddress12 {
	p.NameAndAddress = new(NameAndAddress12)
	return p.NameAndAddress
}

func (p *PartyIdentification58Choice) SetCountry(value string) {
	p.Country = (*CountryCode)(&value)
}
