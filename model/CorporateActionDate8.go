package model

// Specifies corporate action dates.
type CorporateActionDate8 struct {

	// Date/time that the account servicer has set as the deadline to respond, with instructions, to an outstanding event, giving the holder eligibility to incentives. This time is dependent on the reference time zone of the account servicer as specified in a Service Level Agreement (SLA).
	EarlyResponseDeadline *DateFormat6Choice `xml:"EarlyRspnDdln,omitempty"`

	// Last day a holder can deliver the securities that it had elected on and/or previously protected.
	CoverExpirationDate *DateFormat6Choice `xml:"CoverXprtnDt,omitempty"`

	// Last date/time a holder can request to defer delivery of securities pursuant to a notice of guaranteed delivery or other required documentation.
	ProtectDate *DateFormat6Choice `xml:"PrtctDt,omitempty"`

	// Issuer or issuer's agent deadline to respond, with an instruction, to an outstanding offer or privilege.
	MarketDeadline *DateFormat6Choice `xml:"MktDdln,omitempty"`

	// Date/time at which the account servicer has set as the deadline to respond, with instructions, to an outstanding event. This time is dependent on the reference time zone of the account servicer as specified in a Service Level Agreement (SLA).
	ResponseDeadline *DateFormat7Choice `xml:"RspnDdln,omitempty"`

	// Date/time at which an order expires or on which a privilege or offer terminates.
	ExpiryDate *DateFormat6Choice `xml:"XpryDt,omitempty"`

	// Date/time by which cash must be in place in order to take part in the event.
	SubscriptionCostDebitDate *DateFormat6Choice `xml:"SbcptCostDbtDt,omitempty"`

	// Last day that a participant of the depository can deliver securities that it had elected on and/or previously protected.
	DepositoryCoverExpirationDate *DateFormat6Choice `xml:"DpstryCoverXprtnDt,omitempty"`

	// Last day an investor can become a lead plaintiff.
	LeadPlaintiffDeadline *DateFormat6Choice `xml:"LeadPlntffDdln,omitempty"`
}

func (c *CorporateActionDate8) AddEarlyResponseDeadline() *DateFormat6Choice {
	c.EarlyResponseDeadline = new(DateFormat6Choice)
	return c.EarlyResponseDeadline
}

func (c *CorporateActionDate8) AddCoverExpirationDate() *DateFormat6Choice {
	c.CoverExpirationDate = new(DateFormat6Choice)
	return c.CoverExpirationDate
}

func (c *CorporateActionDate8) AddProtectDate() *DateFormat6Choice {
	c.ProtectDate = new(DateFormat6Choice)
	return c.ProtectDate
}

func (c *CorporateActionDate8) AddMarketDeadline() *DateFormat6Choice {
	c.MarketDeadline = new(DateFormat6Choice)
	return c.MarketDeadline
}

func (c *CorporateActionDate8) AddResponseDeadline() *DateFormat7Choice {
	c.ResponseDeadline = new(DateFormat7Choice)
	return c.ResponseDeadline
}

func (c *CorporateActionDate8) AddExpiryDate() *DateFormat6Choice {
	c.ExpiryDate = new(DateFormat6Choice)
	return c.ExpiryDate
}

func (c *CorporateActionDate8) AddSubscriptionCostDebitDate() *DateFormat6Choice {
	c.SubscriptionCostDebitDate = new(DateFormat6Choice)
	return c.SubscriptionCostDebitDate
}

func (c *CorporateActionDate8) AddDepositoryCoverExpirationDate() *DateFormat6Choice {
	c.DepositoryCoverExpirationDate = new(DateFormat6Choice)
	return c.DepositoryCoverExpirationDate
}

func (c *CorporateActionDate8) AddLeadPlaintiffDeadline() *DateFormat6Choice {
	c.LeadPlaintiffDeadline = new(DateFormat6Choice)
	return c.LeadPlaintiffDeadline
}
