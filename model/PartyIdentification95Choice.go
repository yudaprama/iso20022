package model

// Choice of identification of a party.
type PartyIdentification95Choice struct {

	// Code allocated to a financial or non-financial institution by the ISO 9362 Registration Authority, as described in ISO 9362 "Banking - Banking telecommunication messages - Business identifier code (BIC)".
	AnyBIC *AnyBICIdentifier `xml:"AnyBIC"`

	// Unique and unambiguous identifier, as assigned to a financial institution using a proprietary identification scheme.
	ProprietaryIdentification *GenericIdentification79 `xml:"PrtryId"`
}

func (p *PartyIdentification95Choice) SetAnyBIC(value string) {
	p.AnyBIC = (*AnyBICIdentifier)(&value)
}

func (p *PartyIdentification95Choice) AddProprietaryIdentification() *GenericIdentification79 {
	p.ProprietaryIdentification = new(GenericIdentification79)
	return p.ProprietaryIdentification
}
