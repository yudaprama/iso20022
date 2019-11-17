package model

// Choice of identification of a party.
type PartyIdentification103Choice struct {

	// Code allocated to a financial or non-financial institution by the ISO 9362 Registration Authority, as described in ISO 9362 "Banking - Banking telecommunication messages - Business identifier code (BIC)".
	AnyBIC *AnyBICIdentifier `xml:"AnyBIC"`

	// Unique and unambiguous identifier, as assigned to a financial institution using a proprietary identification scheme.
	ProprietaryIdentification *GenericIdentification84 `xml:"PrtryId"`
}

func (p *PartyIdentification103Choice) SetAnyBIC(value string) {
	p.AnyBIC = (*AnyBICIdentifier)(&value)
}

func (p *PartyIdentification103Choice) AddProprietaryIdentification() *GenericIdentification84 {
	p.ProprietaryIdentification = new(GenericIdentification84)
	return p.ProprietaryIdentification
}
