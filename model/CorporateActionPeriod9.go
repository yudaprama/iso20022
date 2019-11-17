package model

// Specifies periods.
type CorporateActionPeriod9 struct {

	// Period during which the price of a security is determined.
	PriceCalculationPeriod *Period5 `xml:"PricClctnPrd,omitempty"`

	// Period during which the specified option, or all options of the event, remains valid, for example, offer period.
	ActionPeriod *Period5 `xml:"ActnPrd,omitempty"`

	// Period during which both old and new equity may be traded simultaneously, for example, consolidation of equity or splitting of equity.
	ParallelTradingPeriod *Period5 `xml:"ParllTradgPrd,omitempty"`
}

func (c *CorporateActionPeriod9) AddPriceCalculationPeriod() *Period5 {
	c.PriceCalculationPeriod = new(Period5)
	return c.PriceCalculationPeriod
}

func (c *CorporateActionPeriod9) AddActionPeriod() *Period5 {
	c.ActionPeriod = new(Period5)
	return c.ActionPeriod
}

func (c *CorporateActionPeriod9) AddParallelTradingPeriod() *Period5 {
	c.ParallelTradingPeriod = new(Period5)
	return c.ParallelTradingPeriod
}
