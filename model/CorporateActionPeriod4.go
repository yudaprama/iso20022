package model

// Specifies periods.
type CorporateActionPeriod4 struct {

	// Period during which the price of a security is determined.
	PriceCalculationPeriod *Period3 `xml:"PricClctnPrd,omitempty"`

	// Period during which the specified option, or all options of the event, remains valid, for example, offer period.
	ActionPeriod *Period3 `xml:"ActnPrd,omitempty"`

	// Period during which both old and new equity may be traded simultaneously, for example, consolidation of equity or splitting of equity.
	ParallelTradingPeriod *Period3 `xml:"ParllTradgPrd,omitempty"`
}

func (c *CorporateActionPeriod4) AddPriceCalculationPeriod() *Period3 {
	c.PriceCalculationPeriod = new(Period3)
	return c.PriceCalculationPeriod
}

func (c *CorporateActionPeriod4) AddActionPeriod() *Period3 {
	c.ActionPeriod = new(Period3)
	return c.ActionPeriod
}

func (c *CorporateActionPeriod4) AddParallelTradingPeriod() *Period3 {
	c.ParallelTradingPeriod = new(Period3)
	return c.ParallelTradingPeriod
}
