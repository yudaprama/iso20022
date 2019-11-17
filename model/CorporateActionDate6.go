package model

// Specifies corporate action date.
type CorporateActionDate6 struct {

	// Date/time at which the account servicer has set as the deadline to respond, with instructions, to an outstanding event. This time is dependent on the reference time zone of the account servicer as specified in a Service Level Agreement (SLA).
	ResponseDeadline *DateFormat6Choice `xml:"RspnDdln,omitempty"`

	// Date/time by which cash must be in place in order to take part in the event.
	SubscriptionCostDebitDate *DateFormat6Choice `xml:"SbcptCostDbtDt,omitempty"`

	// Issuer or issuer's agent deadline to respond, with an instruction, to an outstanding offer or privilege.
	MarketDeadline *DateFormat6Choice `xml:"MktDdln,omitempty"`

	// Date/time at which an order expires or on which a privilege or offer terminates.
	ExpiryDate *DateFormat6Choice `xml:"XpryDt,omitempty"`

	// Last day a holder can deliver the securities that it had elected on and/or previously protected.
	CoverExpirationDate *DateFormat6Choice `xml:"CoverXprtnDt,omitempty"`

	// Last date/time a holder can request to defer delivery of securities pursuant to a notice of guaranteed delivery or other required documentation.
	ProtectDate *DateFormat6Choice `xml:"PrtctDt,omitempty"`

	// Date/time at which the deal (rights) was agreed.
	TradingDate *DateFormat6Choice `xml:"TradgDt,omitempty"`
}

func (c *CorporateActionDate6) AddResponseDeadline() *DateFormat6Choice {
	c.ResponseDeadline = new(DateFormat6Choice)
	return c.ResponseDeadline
}

func (c *CorporateActionDate6) AddSubscriptionCostDebitDate() *DateFormat6Choice {
	c.SubscriptionCostDebitDate = new(DateFormat6Choice)
	return c.SubscriptionCostDebitDate
}

func (c *CorporateActionDate6) AddMarketDeadline() *DateFormat6Choice {
	c.MarketDeadline = new(DateFormat6Choice)
	return c.MarketDeadline
}

func (c *CorporateActionDate6) AddExpiryDate() *DateFormat6Choice {
	c.ExpiryDate = new(DateFormat6Choice)
	return c.ExpiryDate
}

func (c *CorporateActionDate6) AddCoverExpirationDate() *DateFormat6Choice {
	c.CoverExpirationDate = new(DateFormat6Choice)
	return c.CoverExpirationDate
}

func (c *CorporateActionDate6) AddProtectDate() *DateFormat6Choice {
	c.ProtectDate = new(DateFormat6Choice)
	return c.ProtectDate
}

func (c *CorporateActionDate6) AddTradingDate() *DateFormat6Choice {
	c.TradingDate = new(DateFormat6Choice)
	return c.TradingDate
}
