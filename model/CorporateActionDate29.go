package model

// Specifies corporate action dates.
type CorporateActionDate29 struct {

	// Date/time that the account servicer has set as the deadline to respond, with instructions, to an outstanding event, giving the holder eligibility to incentives. This time is dependent on the reference time zone of the account servicer as specified in a Service Level Agreement (SLA).
	EarlyResponseDeadline *DateFormat19Choice `xml:"EarlyRspnDdln,omitempty"`

	// Last day a holder can deliver the securities that it had elected on and/or previously protected.
	CoverExpirationDate *DateFormat19Choice `xml:"CoverXprtnDt,omitempty"`

	// Last date/time a holder can request to defer delivery of securities pursuant to a notice of guaranteed delivery or other required documentation.
	ProtectDate *DateFormat19Choice `xml:"PrtctDt,omitempty"`

	// Issuer or issuer's agent deadline to respond, with an instruction, to an outstanding offer or privilege.
	MarketDeadline *DateFormat19Choice `xml:"MktDdln,omitempty"`

	// Date/time at which the account servicer has set as the deadline to respond, with instructions, to an outstanding event. This time is dependent on the reference time zone of the account servicer as specified in a Service Level Agreement (SLA).
	ResponseDeadline *DateFormat20Choice `xml:"RspnDdln,omitempty"`

	// Date/time at which an order expires or on which a privilege or offer terminates.
	ExpiryDate *DateFormat19Choice `xml:"XpryDt,omitempty"`

	// Date/time by which cash must be in place in order to take part in the event.
	SubscriptionCostDebitDate *DateFormat19Choice `xml:"SbcptCostDbtDt,omitempty"`

	// Last day that a participant of the depository can deliver securities that it had elected on and/or previously protected.
	DepositoryCoverExpirationDate *DateFormat19Choice `xml:"DpstryCoverXprtnDt,omitempty"`

	// Date/time that the account servicer has set as the deadline to respond, with instructions, to an outstanding event, for which the underlying security is out on loan. This time is dependent on the reference time zone of the account servicer as specified in an SLA.
	StockLendingDeadline *DateFormat19Choice `xml:"StockLndgDdln,omitempty"`
}

func (c *CorporateActionDate29) AddEarlyResponseDeadline() *DateFormat19Choice {
	c.EarlyResponseDeadline = new(DateFormat19Choice)
	return c.EarlyResponseDeadline
}

func (c *CorporateActionDate29) AddCoverExpirationDate() *DateFormat19Choice {
	c.CoverExpirationDate = new(DateFormat19Choice)
	return c.CoverExpirationDate
}

func (c *CorporateActionDate29) AddProtectDate() *DateFormat19Choice {
	c.ProtectDate = new(DateFormat19Choice)
	return c.ProtectDate
}

func (c *CorporateActionDate29) AddMarketDeadline() *DateFormat19Choice {
	c.MarketDeadline = new(DateFormat19Choice)
	return c.MarketDeadline
}

func (c *CorporateActionDate29) AddResponseDeadline() *DateFormat20Choice {
	c.ResponseDeadline = new(DateFormat20Choice)
	return c.ResponseDeadline
}

func (c *CorporateActionDate29) AddExpiryDate() *DateFormat19Choice {
	c.ExpiryDate = new(DateFormat19Choice)
	return c.ExpiryDate
}

func (c *CorporateActionDate29) AddSubscriptionCostDebitDate() *DateFormat19Choice {
	c.SubscriptionCostDebitDate = new(DateFormat19Choice)
	return c.SubscriptionCostDebitDate
}

func (c *CorporateActionDate29) AddDepositoryCoverExpirationDate() *DateFormat19Choice {
	c.DepositoryCoverExpirationDate = new(DateFormat19Choice)
	return c.DepositoryCoverExpirationDate
}

func (c *CorporateActionDate29) AddStockLendingDeadline() *DateFormat19Choice {
	c.StockLendingDeadline = new(DateFormat19Choice)
	return c.StockLendingDeadline
}
