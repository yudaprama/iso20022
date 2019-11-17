package model

// Specifies periods.
type CorporateActionPeriod2 struct {

	// Period during which the assented line is available.
	AssentedLinePeriod *Period1 `xml:"AssntdLinePrd,omitempty"`

	// Period during which the specified option, or all options of the event, remains valid, eg, offer period.
	ActionPeriod *Period1 `xml:"ActnPrd,omitempty"`

	// Period during which the privilege is not available, eg, this can happen whenever a meeting takes place or whenever a coupon payment is due.
	PrivilegeSuspensionPeriod *Period1 `xml:"PrvlgSspnsnPrd,omitempty"`

	// Period during which both old and new equity may be traded simultaneously, eg, consolidation of equity or splitting of equity.
	ParallelTradingPeriod *Period1 `xml:"ParllTradgPrd,omitempty"`

	// Period (last day included) during which an account owner can surrender or sell securities to the issuer and receive the sale proceeds.
	SellThruIssuerPeriod *Period1 `xml:"SellThruIssrPrd,omitempty"`

	// Period during which the shareholder can revoke, change or withdraw its instruction.
	RevocabilityPeriod *Period1 `xml:"RvcbltyPrd,omitempty"`

	// Period during which the price of a security is determined (for  outturn securities).
	PriceCalculationPeriod *Period1 `xml:"PricClctnPrd,omitempty"`
}

func (c *CorporateActionPeriod2) AddAssentedLinePeriod() *Period1 {
	c.AssentedLinePeriod = new(Period1)
	return c.AssentedLinePeriod
}

func (c *CorporateActionPeriod2) AddActionPeriod() *Period1 {
	c.ActionPeriod = new(Period1)
	return c.ActionPeriod
}

func (c *CorporateActionPeriod2) AddPrivilegeSuspensionPeriod() *Period1 {
	c.PrivilegeSuspensionPeriod = new(Period1)
	return c.PrivilegeSuspensionPeriod
}

func (c *CorporateActionPeriod2) AddParallelTradingPeriod() *Period1 {
	c.ParallelTradingPeriod = new(Period1)
	return c.ParallelTradingPeriod
}

func (c *CorporateActionPeriod2) AddSellThruIssuerPeriod() *Period1 {
	c.SellThruIssuerPeriod = new(Period1)
	return c.SellThruIssuerPeriod
}

func (c *CorporateActionPeriod2) AddRevocabilityPeriod() *Period1 {
	c.RevocabilityPeriod = new(Period1)
	return c.RevocabilityPeriod
}

func (c *CorporateActionPeriod2) AddPriceCalculationPeriod() *Period1 {
	c.PriceCalculationPeriod = new(Period1)
	return c.PriceCalculationPeriod
}
