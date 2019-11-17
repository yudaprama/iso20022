package model

// Information about a price report.
type PriceReport3 struct {

	// Information related to the price valuation of a financial instrument.
	PriceValuationDetails []*PriceValuation4 `xml:"PricValtnDtls"`
}

func (p *PriceReport3) AddPriceValuationDetails() *PriceValuation4 {
	newValue := new(PriceValuation4)
	p.PriceValuationDetails = append(p.PriceValuationDetails, newValue)
	return newValue
}
