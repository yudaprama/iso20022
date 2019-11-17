package model

// Specifies corporate action dates.
type CorporateActionDate55 struct {

	// Date/time that the account servicer has set as the deadline to respond, with instructions, to an outstanding event, giving the holder eligibility to incentives. This time is dependent on the reference time zone of the account servicer as specified in a Service Level Agreement (SLA).
	EarlyResponseDeadline *DateFormat34Choice `xml:"EarlyRspnDdln,omitempty"`

	// Last day a holder can deliver the securities that it had elected on and/or previously protected.
	CoverExpirationDate *DateFormat34Choice `xml:"CoverXprtnDt,omitempty"`

	// Last date/time a holder can request to defer delivery of securities pursuant to a notice of guaranteed delivery or other required documentation.
	ProtectDate *DateFormat34Choice `xml:"PrtctDt,omitempty"`

	// Issuer or issuer's agent deadline to respond, with an instruction, to an outstanding offer or privilege.
	MarketDeadline *DateFormat34Choice `xml:"MktDdln,omitempty"`

	// Date/time at which the account servicer has set as the deadline to respond, with instructions, to an outstanding event. This time is dependent on the reference time zone of the account servicer as specified in a Service Level Agreement (SLA).
	ResponseDeadline *DateFormat38Choice `xml:"RspnDdln,omitempty"`

	// Date/time at which an order expires or on which a privilege or offer terminates.
	ExpiryDate *DateFormat34Choice `xml:"XpryDt,omitempty"`

	// Date/time by which cash must be in place in order to take part in the event.
	SubscriptionCostDebitDate *DateFormat34Choice `xml:"SbcptCostDbtDt,omitempty"`

	// Last day that a participant of the depository can deliver securities that it had elected on and/or previously protected.
	DepositoryCoverExpirationDate *DateFormat34Choice `xml:"DpstryCoverXprtnDt,omitempty"`

	// Date/time set as the deadline to respond, with instructions, to an outstanding event, for which the underlying security is out on loan.
	StockLendingDeadline *DateFormat34Choice `xml:"StockLndgDdln,omitempty"`

	// Specifies the party borrowing stocks and the associated stock lending deadline assigned to the borrower.
	BorrowerStockLendingDeadline []*BorrowerLendingDeadline2 `xml:"BrrwrStockLndgDdln,omitempty"`
}

func (c *CorporateActionDate55) AddEarlyResponseDeadline() *DateFormat34Choice {
	c.EarlyResponseDeadline = new(DateFormat34Choice)
	return c.EarlyResponseDeadline
}

func (c *CorporateActionDate55) AddCoverExpirationDate() *DateFormat34Choice {
	c.CoverExpirationDate = new(DateFormat34Choice)
	return c.CoverExpirationDate
}

func (c *CorporateActionDate55) AddProtectDate() *DateFormat34Choice {
	c.ProtectDate = new(DateFormat34Choice)
	return c.ProtectDate
}

func (c *CorporateActionDate55) AddMarketDeadline() *DateFormat34Choice {
	c.MarketDeadline = new(DateFormat34Choice)
	return c.MarketDeadline
}

func (c *CorporateActionDate55) AddResponseDeadline() *DateFormat38Choice {
	c.ResponseDeadline = new(DateFormat38Choice)
	return c.ResponseDeadline
}

func (c *CorporateActionDate55) AddExpiryDate() *DateFormat34Choice {
	c.ExpiryDate = new(DateFormat34Choice)
	return c.ExpiryDate
}

func (c *CorporateActionDate55) AddSubscriptionCostDebitDate() *DateFormat34Choice {
	c.SubscriptionCostDebitDate = new(DateFormat34Choice)
	return c.SubscriptionCostDebitDate
}

func (c *CorporateActionDate55) AddDepositoryCoverExpirationDate() *DateFormat34Choice {
	c.DepositoryCoverExpirationDate = new(DateFormat34Choice)
	return c.DepositoryCoverExpirationDate
}

func (c *CorporateActionDate55) AddStockLendingDeadline() *DateFormat34Choice {
	c.StockLendingDeadline = new(DateFormat34Choice)
	return c.StockLendingDeadline
}

func (c *CorporateActionDate55) AddBorrowerStockLendingDeadline() *BorrowerLendingDeadline2 {
	newValue := new(BorrowerLendingDeadline2)
	c.BorrowerStockLendingDeadline = append(c.BorrowerStockLendingDeadline, newValue)
	return newValue
}
