package model

// Performance factors of the investment fund / fund class.
type PerformanceFactors1 struct {

	// Value of the NAV before all corporate events of the valuation date, divided by the value of the NAV after the corporate event.
	CorporateActionFactor *DecimalNumber `xml:"CorpActnFctr,omitempty"`

	// Value of the NAV before a corporate event, divided by the value of the NAV after the corporate event, accumulated for a number of corporate events over the defined period of time.
	CumulativeCorporateActionFactor *DecimalNumber `xml:"CmltvCorpActnFctr,omitempty"`

	// Period of time for the calculation of the cumulative corporate action factor.
	AccumulationPeriod *DatePeriodDetails `xml:"AcmltnPrd,omitempty"`

	// Normal performance value of the NAV.
	NormalPerformance *DecimalNumber `xml:"NrmlPrfrmnc,omitempty"`
}

func (p *PerformanceFactors1) SetCorporateActionFactor(value string) {
	p.CorporateActionFactor = (*DecimalNumber)(&value)
}

func (p *PerformanceFactors1) SetCumulativeCorporateActionFactor(value string) {
	p.CumulativeCorporateActionFactor = (*DecimalNumber)(&value)
}

func (p *PerformanceFactors1) AddAccumulationPeriod() *DatePeriodDetails {
	p.AccumulationPeriod = new(DatePeriodDetails)
	return p.AccumulationPeriod
}

func (p *PerformanceFactors1) SetNormalPerformance(value string) {
	p.NormalPerformance = (*DecimalNumber)(&value)
}
