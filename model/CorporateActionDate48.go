package model

// Specifies corporate action dates.
type CorporateActionDate48 struct {

	// Date/time that the account servicer has set as the deadline to respond, with instructions, to an outstanding event, giving the holder eligibility to incentives. This time is dependent on the reference time zone of the account servicer as specified in a Service Level Agreement (SLA).
	EarlyResponseDeadline *DateFormat31Choice `xml:"EarlyRspnDdln,omitempty"`

	// Last day a holder can deliver the securities that it had elected on and/or previously protected.
	CoverExpirationDate *DateFormat31Choice `xml:"CoverXprtnDt,omitempty"`

	// Last date/time a holder can request to defer delivery of securities pursuant to a notice of guaranteed delivery or other required documentation.
	ProtectDate *DateFormat31Choice `xml:"PrtctDt,omitempty"`

	// Issuer or issuer's agent deadline to respond, with an instruction, to an outstanding offer or privilege.
	MarketDeadline *DateFormat31Choice `xml:"MktDdln,omitempty"`

	// Date/time at which the account servicer has set as the deadline to respond, with instructions, to an outstanding event. This time is dependent on the reference time zone of the account servicer as specified in a Service Level Agreement (SLA).
	ResponseDeadline *DateFormat32Choice `xml:"RspnDdln,omitempty"`

	// Date/time at which an order expires or on which a privilege or offer terminates.
	ExpiryDate *DateFormat31Choice `xml:"XpryDt,omitempty"`

	// Date/time by which cash must be in place in order to take part in the event.
	SubscriptionCostDebitDate *DateFormat31Choice `xml:"SbcptCostDbtDt,omitempty"`

	// Last day that a participant of the depository can deliver securities that it had elected on and/or previously protected.
	DepositoryCoverExpirationDate *DateFormat31Choice `xml:"DpstryCoverXprtnDt,omitempty"`

	// Date/time set as the deadline to respond, with instructions, to an outstanding event, for which the underlying security is out on loan.
	StockLendingDeadline *DateFormat31Choice `xml:"StockLndgDdln,omitempty"`

	// Specifies the party borrowing stocks and the associated stock lending deadline assigned to the borrower.
	BorrowerStockLendingDeadline []*BorrowerLendingDeadline1 `xml:"BrrwrStockLndgDdln,omitempty"`
}

func (c *CorporateActionDate48) AddEarlyResponseDeadline() *DateFormat31Choice {
	c.EarlyResponseDeadline = new(DateFormat31Choice)
	return c.EarlyResponseDeadline
}

func (c *CorporateActionDate48) AddCoverExpirationDate() *DateFormat31Choice {
	c.CoverExpirationDate = new(DateFormat31Choice)
	return c.CoverExpirationDate
}

func (c *CorporateActionDate48) AddProtectDate() *DateFormat31Choice {
	c.ProtectDate = new(DateFormat31Choice)
	return c.ProtectDate
}

func (c *CorporateActionDate48) AddMarketDeadline() *DateFormat31Choice {
	c.MarketDeadline = new(DateFormat31Choice)
	return c.MarketDeadline
}

func (c *CorporateActionDate48) AddResponseDeadline() *DateFormat32Choice {
	c.ResponseDeadline = new(DateFormat32Choice)
	return c.ResponseDeadline
}

func (c *CorporateActionDate48) AddExpiryDate() *DateFormat31Choice {
	c.ExpiryDate = new(DateFormat31Choice)
	return c.ExpiryDate
}

func (c *CorporateActionDate48) AddSubscriptionCostDebitDate() *DateFormat31Choice {
	c.SubscriptionCostDebitDate = new(DateFormat31Choice)
	return c.SubscriptionCostDebitDate
}

func (c *CorporateActionDate48) AddDepositoryCoverExpirationDate() *DateFormat31Choice {
	c.DepositoryCoverExpirationDate = new(DateFormat31Choice)
	return c.DepositoryCoverExpirationDate
}

func (c *CorporateActionDate48) AddStockLendingDeadline() *DateFormat31Choice {
	c.StockLendingDeadline = new(DateFormat31Choice)
	return c.StockLendingDeadline
}

func (c *CorporateActionDate48) AddBorrowerStockLendingDeadline() *BorrowerLendingDeadline1 {
	newValue := new(BorrowerLendingDeadline1)
	c.BorrowerStockLendingDeadline = append(c.BorrowerStockLendingDeadline, newValue)
	return newValue
}
