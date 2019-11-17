package model

// Specifies periods.
type CorporateActionPeriod11 struct {

	// Period during which the price of a security is determined.
	PriceCalculationPeriod *Period4 `xml:"PricClctnPrd,omitempty"`

	// Period during which the specified option, or all options of the event, remains valid, for example, offer period.
	ActionPeriod *Period4 `xml:"ActnPrd,omitempty"`

	// Period during which both old and new equity may be traded simultaneously, for example, consolidation of equity or splitting of equity.
	ParallelTradingPeriod *Period4 `xml:"ParllTradgPrd,omitempty"`
}

func (c *CorporateActionPeriod11) AddPriceCalculationPeriod() *Period4 {
	c.PriceCalculationPeriod = new(Period4)
	return c.PriceCalculationPeriod
}

func (c *CorporateActionPeriod11) AddActionPeriod() *Period4 {
	c.ActionPeriod = new(Period4)
	return c.ActionPeriod
}

func (c *CorporateActionPeriod11) AddParallelTradingPeriod() *Period4 {
	c.ParallelTradingPeriod = new(Period4)
	return c.ParallelTradingPeriod
}
