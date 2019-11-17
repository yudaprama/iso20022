package model

// Specifies corporate action dates.
type CorporateActionDate4 struct {

	// Date/time at which the coupons are to be/were submitted for payment of interest.
	CouponClippingDate *DateFormat4Choice `xml:"CpnClpngDt,omitempty"`

	// Last date/time at which a holder can consent to the changes sought by the corporation.
	ConsentExpirationDate *DateFormat4Choice `xml:"CnsntXprtnDt,omitempty"`

	// Date/time used by the offeror to determine the beneficiary eligible to participate in a consent based on the registered owner of the securities, eg, beneficial owner of consent record. The consent record date qualifier is used to indicate that a record date only applies to a certain part of the offer, not the entire offer.
	ConsentRecordDate *DateFormat4Choice `xml:"CnsntRcrdDt,omitempty"`

	// Date/time at which the distribution is due to take place (cash and/or securities).
	PaymentDate *DateFormat4Choice `xml:"PmtDt,omitempty"`

	// Date/time at which a payment can be made, eg, if payment date is a non-business day or to indicate the first payment date of an offer.
	EarliestPaymentDate *DateFormat4Choice `xml:"EarlstPmtDt,omitempty"`

	// Issuer or issuer's agent deadline to respond, with an instruction, to an outstanding offer or privilege.
	MarketDeadline *DateFormat4Choice `xml:"MktDdln,omitempty"`

	// Date/time at which the account servicer has set as the deadline to respond, with instructions, to an outstanding event. This time is dependent on the reference time zone of the account servicer as specified in an SLA.
	ResponseDeadline *DateFormat4Choice `xml:"RspnDdln,omitempty"`

	// Deadline by which instructions must be received to split securities, eg, of physical certificates.
	DeadlineToSplit *DateFormat4Choice `xml:"DdlnToSplt,omitempty"`

	// Date/time at which an order expires or on which a privilege or offer terminates.
	ExpiryDate *DateFormat4Choice `xml:"XpryDt,omitempty"`

	// Date/time at which the price of a security is determined.
	QuotationSettingDate *DateFormat4Choice `xml:"QtnSetngDt,omitempty"`

	// Date/time by which cash must be in place in order to take part in the event.
	SubscriptionCostDebitDate *DateFormat4Choice `xml:"SbcptCostDbtDt,omitempty"`
}

func (c *CorporateActionDate4) AddCouponClippingDate() *DateFormat4Choice {
	c.CouponClippingDate = new(DateFormat4Choice)
	return c.CouponClippingDate
}

func (c *CorporateActionDate4) AddConsentExpirationDate() *DateFormat4Choice {
	c.ConsentExpirationDate = new(DateFormat4Choice)
	return c.ConsentExpirationDate
}

func (c *CorporateActionDate4) AddConsentRecordDate() *DateFormat4Choice {
	c.ConsentRecordDate = new(DateFormat4Choice)
	return c.ConsentRecordDate
}

func (c *CorporateActionDate4) AddPaymentDate() *DateFormat4Choice {
	c.PaymentDate = new(DateFormat4Choice)
	return c.PaymentDate
}

func (c *CorporateActionDate4) AddEarliestPaymentDate() *DateFormat4Choice {
	c.EarliestPaymentDate = new(DateFormat4Choice)
	return c.EarliestPaymentDate
}

func (c *CorporateActionDate4) AddMarketDeadline() *DateFormat4Choice {
	c.MarketDeadline = new(DateFormat4Choice)
	return c.MarketDeadline
}

func (c *CorporateActionDate4) AddResponseDeadline() *DateFormat4Choice {
	c.ResponseDeadline = new(DateFormat4Choice)
	return c.ResponseDeadline
}

func (c *CorporateActionDate4) AddDeadlineToSplit() *DateFormat4Choice {
	c.DeadlineToSplit = new(DateFormat4Choice)
	return c.DeadlineToSplit
}

func (c *CorporateActionDate4) AddExpiryDate() *DateFormat4Choice {
	c.ExpiryDate = new(DateFormat4Choice)
	return c.ExpiryDate
}

func (c *CorporateActionDate4) AddQuotationSettingDate() *DateFormat4Choice {
	c.QuotationSettingDate = new(DateFormat4Choice)
	return c.QuotationSettingDate
}

func (c *CorporateActionDate4) AddSubscriptionCostDebitDate() *DateFormat4Choice {
	c.SubscriptionCostDebitDate = new(DateFormat4Choice)
	return c.SubscriptionCostDebitDate
}
