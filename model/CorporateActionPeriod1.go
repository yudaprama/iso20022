package model

// Specifies periods.
type CorporateActionPeriod1 struct {

	// Period during which the specified option, or all options of the event, remains valid, eg, offer period.
	ActionPeriod *Period1 `xml:"ActnPrd,omitempty"`

	// Period during a take-over where any outstanding equity must be purchased by the take-over company.
	CompulsoryPurchasePeriod *Period1 `xml:"CmplsryPurchsPrd,omitempty"`

	// Period during which the interest rate has been applied.
	InterestPeriod *Period1 `xml:"IntrstPrd,omitempty"`

	// Period during which the security is blocked.
	BlockingPeriod *Period1 `xml:"BlckgPrd,omitempty"`

	// Period during which the price of a security is determined.
	PriceCalculationPeriod *Period1 `xml:"PricClctnPrd,omitempty"`
}

func (c *CorporateActionPeriod1) AddActionPeriod() *Period1 {
	c.ActionPeriod = new(Period1)
	return c.ActionPeriod
}

func (c *CorporateActionPeriod1) AddCompulsoryPurchasePeriod() *Period1 {
	c.CompulsoryPurchasePeriod = new(Period1)
	return c.CompulsoryPurchasePeriod
}

func (c *CorporateActionPeriod1) AddInterestPeriod() *Period1 {
	c.InterestPeriod = new(Period1)
	return c.InterestPeriod
}

func (c *CorporateActionPeriod1) AddBlockingPeriod() *Period1 {
	c.BlockingPeriod = new(Period1)
	return c.BlockingPeriod
}

func (c *CorporateActionPeriod1) AddPriceCalculationPeriod() *Period1 {
	c.PriceCalculationPeriod = new(Period1)
	return c.PriceCalculationPeriod
}
