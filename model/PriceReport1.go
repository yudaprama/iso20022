package model

// Information about a price report.
type PriceReport1 struct {

	// Information related to the price valuation of a financial instrument.
	PriceValuationDetails []*PriceValuation2 `xml:"PricValtnDtls"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (p *PriceReport1) AddPriceValuationDetails() *PriceValuation2 {
	newValue := new(PriceValuation2)
	p.PriceValuationDetails = append(p.PriceValuationDetails, newValue)
	return newValue
}

func (p *PriceReport1) AddExtension() *Extension1 {
	newValue := new(Extension1)
	p.Extension = append(p.Extension, newValue)
	return newValue
}
