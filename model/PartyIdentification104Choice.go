package model

// Choice between different formats for the identification of a party.
type PartyIdentification104Choice struct {

	// Code allocated to a financial or non-financial institution by the ISO 9362 Registration Authority, as described in ISO 9362 "Banking - Banking telecommunication messages - Business identifier code (BIC)".
	AnyBIC *AnyBICIdentifier `xml:"AnyBIC"`

	// Unique and unambiguous identifier, as assigned to a financial institution using a proprietary identification scheme.
	ProprietaryIdentification *GenericIdentification84 `xml:"PrtryId"`

	// Name and address of a party.
	NameAndAddress *NameAndAddress12 `xml:"NmAndAdr"`
}

func (p *PartyIdentification104Choice) SetAnyBIC(value string) {
	p.AnyBIC = (*AnyBICIdentifier)(&value)
}

func (p *PartyIdentification104Choice) AddProprietaryIdentification() *GenericIdentification84 {
	p.ProprietaryIdentification = new(GenericIdentification84)
	return p.ProprietaryIdentification
}

func (p *PartyIdentification104Choice) AddNameAndAddress() *NameAndAddress12 {
	p.NameAndAddress = new(NameAndAddress12)
	return p.NameAndAddress
}
