package model

// Unique and unambiguous way to identify an organisation.
type PartyIdentification44 struct {

	// Code allocated to a financial or non-financial institution by the ISO 9362 Registration Authority, as described in ISO 9362 "Banking - Banking telecommunication messages - Business identifier code (BIC)".
	AnyBIC *AnyBICIdentifier `xml:"AnyBIC"`

	// Unique and unambiguous identifier, as assigned to a financial institution using a proprietary identification scheme.
	AlternativeIdentifier []*Max35Text `xml:"AltrntvIdr,omitempty"`
}

func (p *PartyIdentification44) SetAnyBIC(value string) {
	p.AnyBIC = (*AnyBICIdentifier)(&value)
}

func (p *PartyIdentification44) AddAlternativeIdentifier(value string) {
	p.AlternativeIdentifier = append(p.AlternativeIdentifier, (*Max35Text)(&value))
}
