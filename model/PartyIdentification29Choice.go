package model

// Identification of a party by BIC or by name and address.
type PartyIdentification29Choice struct {

	// Code allocated to a financial institution by the ISO 9362 Registration Authority as described in ISO 9362 "Banking - Banking telecommunication messages - Business identifier code (BIC)".
	BIC *BICIdentifier `xml:"BIC"`

	// Identifies the name and address of a non-financial institution.
	NameAndAddress *PartyIdentification26 `xml:"NmAndAdr"`
}

func (p *PartyIdentification29Choice) SetBIC(value string) {
	p.BIC = (*BICIdentifier)(&value)
}

func (p *PartyIdentification29Choice) AddNameAndAddress() *PartyIdentification26 {
	p.NameAndAddress = new(PartyIdentification26)
	return p.NameAndAddress
}
