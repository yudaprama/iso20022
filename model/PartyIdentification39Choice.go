package model

// Unique and unambiguous way to identify an organisation.
type PartyIdentification39Choice struct {

	// Unique and unambiguous way to identify an organisation.
	AnyBIC *AnyBICIdentifier `xml:"AnyBIC"`

	// Unique and unambiguous way to identify an organisation.
	NameAndAddress *NameAndAddress5 `xml:"NmAndAdr"`

	// Unique and unambiguous way to identify an organisation.
	Country *CountryCode `xml:"Ctry"`
}

func (p *PartyIdentification39Choice) SetAnyBIC(value string) {
	p.AnyBIC = (*AnyBICIdentifier)(&value)
}

func (p *PartyIdentification39Choice) AddNameAndAddress() *NameAndAddress5 {
	p.NameAndAddress = new(NameAndAddress5)
	return p.NameAndAddress
}

func (p *PartyIdentification39Choice) SetCountry(value string) {
	p.Country = (*CountryCode)(&value)
}
