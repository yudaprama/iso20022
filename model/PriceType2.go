package model

// Specification of the price type.
type PriceType2 struct {

	// Structured format.
	Structured *TypeOfPrice6Code `xml:"Strd"`

	// Additional information about the type of charge.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (p *PriceType2) SetStructured(value string) {
	p.Structured = (*TypeOfPrice6Code)(&value)
}

func (p *PriceType2) SetAdditionalInformation(value string) {
	p.AdditionalInformation = (*Max350Text)(&value)
}
