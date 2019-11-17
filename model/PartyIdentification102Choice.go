package model

// Unique and unambiguous way to identify an organisation using a choice between a BIC or the name and addres or the country code.
type PartyIdentification102Choice struct {

	// Unique and unambiguous way to identify an organisation.
	AnyBIC *AnyBICIdentifier `xml:"AnyBIC"`

	// Unique and unambiguous way to identify an organisation, using the name and address.
	NameAndAddress *NameAndAddress5 `xml:"NmAndAdr"`

	// Unique and unambiguous way to identify an organisation using the country code, using the country code.
	Country *CountryCode `xml:"Ctry"`
}

func (p *PartyIdentification102Choice) SetAnyBIC(value string) {
	p.AnyBIC = (*AnyBICIdentifier)(&value)
}

func (p *PartyIdentification102Choice) AddNameAndAddress() *NameAndAddress5 {
	p.NameAndAddress = new(NameAndAddress5)
	return p.NameAndAddress
}

func (p *PartyIdentification102Choice) SetCountry(value string) {
	p.Country = (*CountryCode)(&value)
}
