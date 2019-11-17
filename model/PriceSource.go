package model

// Source of a price quotation when it is not the local market.
type PriceSource struct {

	// Source of the price.
	PriceSource *PriceSource1Code `xml:"PricSrc"`

	// Additional information about the source of a price.
	Narrative *Max35Text `xml:"Nrrtv,omitempty"`
}

func (p *PriceSource) SetPriceSource(value string) {
	p.PriceSource = (*PriceSource1Code)(&value)
}

func (p *PriceSource) SetNarrative(value string) {
	p.Narrative = (*Max35Text)(&value)
}
