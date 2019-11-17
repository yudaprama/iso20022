package model

// Original and corrected price information of an investment fund.
type PriceCorrection3 struct {

	// Information related to the price valuation of a financial instrument sent in a previous price report.
	PreviouslySentPriceDetails *PriceValuation3 `xml:"PrevslySntPricDtls"`

	// Information related to the new price valuation of a financial instrument, which overrides previously sent information.
	CorrectedPriceDetails *PriceValuation3 `xml:"CrrctdPricDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (p *PriceCorrection3) AddPreviouslySentPriceDetails() *PriceValuation3 {
	p.PreviouslySentPriceDetails = new(PriceValuation3)
	return p.PreviouslySentPriceDetails
}

func (p *PriceCorrection3) AddCorrectedPriceDetails() *PriceValuation3 {
	p.CorrectedPriceDetails = new(PriceValuation3)
	return p.CorrectedPriceDetails
}

func (p *PriceCorrection3) AddExtension() *Extension1 {
	newValue := new(Extension1)
	p.Extension = append(p.Extension, newValue)
	return newValue
}
