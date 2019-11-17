package model

// Specifies corporate action date.
type CorporateActionDate46 struct {

	// Date/time at which the account servicer has set as the deadline to respond, with instructions, to an outstanding event. This time is dependent on the reference time zone of the account servicer as specified in a Service Level Agreement (SLA).
	ResponseDeadline *DateFormat31Choice `xml:"RspnDdln,omitempty"`

	// Date/time by which cash must be in place in order to take part in the event.
	SubscriptionCostDebitDate *DateFormat31Choice `xml:"SbcptCostDbtDt,omitempty"`

	// Issuer or issuer's agent deadline to respond, with an instruction, to an outstanding offer or privilege.
	MarketDeadline *DateFormat31Choice `xml:"MktDdln,omitempty"`

	// Date/time at which an order expires or on which a privilege or offer terminates.
	ExpiryDate *DateFormat31Choice `xml:"XpryDt,omitempty"`

	// Last day a holder can deliver the securities that it had elected on and/or previously protected.
	CoverExpirationDate *DateFormat31Choice `xml:"CoverXprtnDt,omitempty"`

	// Last date/time a holder can request to defer delivery of securities pursuant to a notice of guaranteed delivery or other required documentation.
	ProtectDate *DateFormat31Choice `xml:"PrtctDt,omitempty"`

	// Date/time at which the deal (rights) was agreed.
	TradingDate *DateFormat31Choice `xml:"TradgDt,omitempty"`
}

func (c *CorporateActionDate46) AddResponseDeadline() *DateFormat31Choice {
	c.ResponseDeadline = new(DateFormat31Choice)
	return c.ResponseDeadline
}

func (c *CorporateActionDate46) AddSubscriptionCostDebitDate() *DateFormat31Choice {
	c.SubscriptionCostDebitDate = new(DateFormat31Choice)
	return c.SubscriptionCostDebitDate
}

func (c *CorporateActionDate46) AddMarketDeadline() *DateFormat31Choice {
	c.MarketDeadline = new(DateFormat31Choice)
	return c.MarketDeadline
}

func (c *CorporateActionDate46) AddExpiryDate() *DateFormat31Choice {
	c.ExpiryDate = new(DateFormat31Choice)
	return c.ExpiryDate
}

func (c *CorporateActionDate46) AddCoverExpirationDate() *DateFormat31Choice {
	c.CoverExpirationDate = new(DateFormat31Choice)
	return c.CoverExpirationDate
}

func (c *CorporateActionDate46) AddProtectDate() *DateFormat31Choice {
	c.ProtectDate = new(DateFormat31Choice)
	return c.ProtectDate
}

func (c *CorporateActionDate46) AddTradingDate() *DateFormat31Choice {
	c.TradingDate = new(DateFormat31Choice)
	return c.TradingDate
}
