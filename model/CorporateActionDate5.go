package model

// Specifies corporate action dates.
type CorporateActionDate5 struct {

	// Date/time at which a foreign exchange rate will be determined.
	ForeignExchangeRateFixingDate *DateFormat4Choice `xml:"FXRateFxgDt,omitempty"`

	// Date/time at which assets become available to the account owner (in a credit entry), or cease to be available to the account owner (in a debit entry).
	ValueDate *DateFormat4Choice `xml:"ValDt,omitempty"`

	// Date/time at which the distribution is due to take place (cash and/or securities).
	PaymentDate *DateFormat4Choice `xml:"PmtDt,omitempty"`

	// Date/time at which a payment can be made, eg, if payment date is a non-business day or to indicate the first payment date of an offer.
	EarliestPaymentDate *DateFormat4Choice `xml:"EarlstPmtDt,omitempty"`
}

func (c *CorporateActionDate5) AddForeignExchangeRateFixingDate() *DateFormat4Choice {
	c.ForeignExchangeRateFixingDate = new(DateFormat4Choice)
	return c.ForeignExchangeRateFixingDate
}

func (c *CorporateActionDate5) AddValueDate() *DateFormat4Choice {
	c.ValueDate = new(DateFormat4Choice)
	return c.ValueDate
}

func (c *CorporateActionDate5) AddPaymentDate() *DateFormat4Choice {
	c.PaymentDate = new(DateFormat4Choice)
	return c.PaymentDate
}

func (c *CorporateActionDate5) AddEarliestPaymentDate() *DateFormat4Choice {
	c.EarliestPaymentDate = new(DateFormat4Choice)
	return c.EarliestPaymentDate
}
