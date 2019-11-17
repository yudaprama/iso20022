package model

// Choice of identification of a party.
type PartyIdentification32Choice struct {

	// Code allocated to a financial or non-financial institution by the ISO 9362 Registration Authority, as described in ISO 9362 "Banking - Banking telecommunication messages - Business identifier code (BIC)".
	BIC *AnyBICIdentifier `xml:"BIC"`

	// Unique and unambiguous identifier, as assigned to a financial institution using a proprietary identification scheme.
	ProprietaryIdentification *GenericIdentification29 `xml:"PrtryId"`

	// Identification of a party with its name and address in free text.
	NameAndAddress *NameAndAddress13 `xml:"NmAndAdr"`
}

func (p *PartyIdentification32Choice) SetBIC(value string) {
	p.BIC = (*AnyBICIdentifier)(&value)
}

func (p *PartyIdentification32Choice) AddProprietaryIdentification() *GenericIdentification29 {
	p.ProprietaryIdentification = new(GenericIdentification29)
	return p.ProprietaryIdentification
}

func (p *PartyIdentification32Choice) AddNameAndAddress() *NameAndAddress13 {
	p.NameAndAddress = new(NameAndAddress13)
	return p.NameAndAddress
}
