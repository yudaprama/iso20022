package model

// Specifies corporate action date.
type CorporateActionDate52 struct {

	// Date/time at which the account servicer has set as the deadline to respond, with instructions, to an outstanding event. This time is dependent on the reference time zone of the account servicer as specified in a Service Level Agreement (SLA).
	ResponseDeadline *DateFormat34Choice `xml:"RspnDdln,omitempty"`

	// Date/time by which cash must be in place in order to take part in the event.
	SubscriptionCostDebitDate *DateFormat34Choice `xml:"SbcptCostDbtDt,omitempty"`

	// Issuer or issuer's agent deadline to respond, with an instruction, to an outstanding offer or privilege.
	MarketDeadline *DateFormat34Choice `xml:"MktDdln,omitempty"`

	// Date/time at which an order expires or on which a privilege or offer terminates.
	ExpiryDate *DateFormat34Choice `xml:"XpryDt,omitempty"`

	// Last day a holder can deliver the securities that it had elected on and/or previously protected.
	CoverExpirationDate *DateFormat34Choice `xml:"CoverXprtnDt,omitempty"`

	// Last date/time a holder can request to defer delivery of securities pursuant to a notice of guaranteed delivery or other required documentation.
	ProtectDate *DateFormat34Choice `xml:"PrtctDt,omitempty"`

	// Date/time at which the deal (rights) was agreed.
	TradingDate *DateFormat34Choice `xml:"TradgDt,omitempty"`
}

func (c *CorporateActionDate52) AddResponseDeadline() *DateFormat34Choice {
	c.ResponseDeadline = new(DateFormat34Choice)
	return c.ResponseDeadline
}

func (c *CorporateActionDate52) AddSubscriptionCostDebitDate() *DateFormat34Choice {
	c.SubscriptionCostDebitDate = new(DateFormat34Choice)
	return c.SubscriptionCostDebitDate
}

func (c *CorporateActionDate52) AddMarketDeadline() *DateFormat34Choice {
	c.MarketDeadline = new(DateFormat34Choice)
	return c.MarketDeadline
}

func (c *CorporateActionDate52) AddExpiryDate() *DateFormat34Choice {
	c.ExpiryDate = new(DateFormat34Choice)
	return c.ExpiryDate
}

func (c *CorporateActionDate52) AddCoverExpirationDate() *DateFormat34Choice {
	c.CoverExpirationDate = new(DateFormat34Choice)
	return c.CoverExpirationDate
}

func (c *CorporateActionDate52) AddProtectDate() *DateFormat34Choice {
	c.ProtectDate = new(DateFormat34Choice)
	return c.ProtectDate
}

func (c *CorporateActionDate52) AddTradingDate() *DateFormat34Choice {
	c.TradingDate = new(DateFormat34Choice)
	return c.TradingDate
}
