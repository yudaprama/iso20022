package model

// Choice of identification of a party.
type PartyIdentification68Choice struct {

	// Code allocated to a financial or non-financial institution by the ISO 9362 Registration Authority, as described in ISO 9362 "Banking - Banking telecommunication messages - Business identifier code (BIC)".
	BIC *AnyBICIdentifier `xml:"BIC"`

	// Identification of a party with its name and address in free text.
	NameAndAddress *NameAndAddress13 `xml:"NmAndAdr"`

	// Nation with its own government, occupying a particular territory.
	Country *CountryCode `xml:"Ctry"`
}

func (p *PartyIdentification68Choice) SetBIC(value string) {
	p.BIC = (*AnyBICIdentifier)(&value)
}

func (p *PartyIdentification68Choice) AddNameAndAddress() *NameAndAddress13 {
	p.NameAndAddress = new(NameAndAddress13)
	return p.NameAndAddress
}

func (p *PartyIdentification68Choice) SetCountry(value string) {
	p.Country = (*CountryCode)(&value)
}
