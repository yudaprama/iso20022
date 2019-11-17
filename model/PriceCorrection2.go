package model

// Original and corrected price information of an investment fund.
type PriceCorrection2 struct {

	// Information related to the price valuation of a financial instrument sent in a previous price report.
	PreviouslySentPriceDetails *PriceValuation2 `xml:"PrevslySntPricDtls"`

	// Information related to the new price valuation of a financial instrument, which overrides previously sent information.
	CorrectedPriceDetails *PriceValuation2 `xml:"CrrctdPricDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (p *PriceCorrection2) AddPreviouslySentPriceDetails() *PriceValuation2 {
	p.PreviouslySentPriceDetails = new(PriceValuation2)
	return p.PreviouslySentPriceDetails
}

func (p *PriceCorrection2) AddCorrectedPriceDetails() *PriceValuation2 {
	p.CorrectedPriceDetails = new(PriceValuation2)
	return p.CorrectedPriceDetails
}

func (p *PriceCorrection2) AddExtension() *Extension1 {
	newValue := new(Extension1)
	p.Extension = append(p.Extension, newValue)
	return newValue
}
