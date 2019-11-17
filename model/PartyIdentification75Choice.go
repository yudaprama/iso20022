package model

// Choice of identification of a party.
type PartyIdentification75Choice struct {

	// Identification of the party expressed as a BIC.
	AnyBIC *AnyBICIdentifier `xml:"AnyBIC"`

	// Name and address of a party.
	NameAndAddress *NameAndAddress5 `xml:"NmAndAdr"`

	// Country of the party.
	Country *CountryCode `xml:"Ctry"`
}

func (p *PartyIdentification75Choice) SetAnyBIC(value string) {
	p.AnyBIC = (*AnyBICIdentifier)(&value)
}

func (p *PartyIdentification75Choice) AddNameAndAddress() *NameAndAddress5 {
	p.NameAndAddress = new(NameAndAddress5)
	return p.NameAndAddress
}

func (p *PartyIdentification75Choice) SetCountry(value string) {
	p.Country = (*CountryCode)(&value)
}
