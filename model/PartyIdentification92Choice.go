package model

// Choice of identification of a party.
type PartyIdentification92Choice struct {

	// Code allocated to a financial or non-financial institution by the ISO 9362 Registration Authority, as described in ISO 9362 "Banking - Banking telecommunication messages - Business identifier code (BIC)".
	AnyBIC *AnyBICIdentifier `xml:"AnyBIC"`

	// Unique and unambiguous identifier, as assigned to a financial institution using a proprietary identification scheme.
	ProprietaryIdentification *GenericIdentification36 `xml:"PrtryId"`
}

func (p *PartyIdentification92Choice) SetAnyBIC(value string) {
	p.AnyBIC = (*AnyBICIdentifier)(&value)
}

func (p *PartyIdentification92Choice) AddProprietaryIdentification() *GenericIdentification36 {
	p.ProprietaryIdentification = new(GenericIdentification36)
	return p.ProprietaryIdentification
}
