package model

// Specifies corporate action date.
type CorporateActionDate18 struct {

	// Date/time at which the account servicer has set as the deadline to respond, with instructions, to an outstanding event. This time is dependent on the reference time zone of the account servicer as specified in a Service Level Agreement (SLA).
	ResponseDeadline *DateFormat19Choice `xml:"RspnDdln,omitempty"`

	// Date/time by which cash must be in place in order to take part in the event.
	SubscriptionCostDebitDate *DateFormat19Choice `xml:"SbcptCostDbtDt,omitempty"`

	// Issuer or issuer's agent deadline to respond, with an instruction, to an outstanding offer or privilege.
	MarketDeadline *DateFormat19Choice `xml:"MktDdln,omitempty"`

	// Date/time at which an order expires or on which a privilege or offer terminates.
	ExpiryDate *DateFormat19Choice `xml:"XpryDt,omitempty"`

	// Last day a holder can deliver the securities that it had elected on and/or previously protected.
	CoverExpirationDate *DateFormat19Choice `xml:"CoverXprtnDt,omitempty"`

	// Last date/time a holder can request to defer delivery of securities pursuant to a notice of guaranteed delivery or other required documentation.
	ProtectDate *DateFormat19Choice `xml:"PrtctDt,omitempty"`

	// Date/time at which the deal (rights) was agreed.
	TradingDate *DateFormat19Choice `xml:"TradgDt,omitempty"`
}

func (c *CorporateActionDate18) AddResponseDeadline() *DateFormat19Choice {
	c.ResponseDeadline = new(DateFormat19Choice)
	return c.ResponseDeadline
}

func (c *CorporateActionDate18) AddSubscriptionCostDebitDate() *DateFormat19Choice {
	c.SubscriptionCostDebitDate = new(DateFormat19Choice)
	return c.SubscriptionCostDebitDate
}

func (c *CorporateActionDate18) AddMarketDeadline() *DateFormat19Choice {
	c.MarketDeadline = new(DateFormat19Choice)
	return c.MarketDeadline
}

func (c *CorporateActionDate18) AddExpiryDate() *DateFormat19Choice {
	c.ExpiryDate = new(DateFormat19Choice)
	return c.ExpiryDate
}

func (c *CorporateActionDate18) AddCoverExpirationDate() *DateFormat19Choice {
	c.CoverExpirationDate = new(DateFormat19Choice)
	return c.CoverExpirationDate
}

func (c *CorporateActionDate18) AddProtectDate() *DateFormat19Choice {
	c.ProtectDate = new(DateFormat19Choice)
	return c.ProtectDate
}

func (c *CorporateActionDate18) AddTradingDate() *DateFormat19Choice {
	c.TradingDate = new(DateFormat19Choice)
	return c.TradingDate
}
