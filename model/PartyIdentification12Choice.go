package model

// Unique and unambiguous way to identify an organisation.
type PartyIdentification12Choice struct {

	// Unique and unambiguous way to identify an organisation.
	BICOrBEI *AnyBICIdentifier `xml:"BICOrBEI"`

	// Unique and unambiguous way to identify an organisation.
	NameAndAddress *NameAndAddress5 `xml:"NmAndAdr"`

	// Unique and unambiguous way to identify an organisation.
	Country *CountryCode `xml:"Ctry"`
}

func (p *PartyIdentification12Choice) SetBICOrBEI(value string) {
	p.BICOrBEI = (*AnyBICIdentifier)(&value)
}

func (p *PartyIdentification12Choice) AddNameAndAddress() *NameAndAddress5 {
	p.NameAndAddress = new(NameAndAddress5)
	return p.NameAndAddress
}

func (p *PartyIdentification12Choice) SetCountry(value string) {
	p.Country = (*CountryCode)(&value)
}
