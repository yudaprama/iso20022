package model

// Specifies corporate action dates.
type CorporateActionDate9 struct {

	// Date on which the movement is due to take place (cash and/or securities).
	PaymentDate *DateFormat6Choice `xml:"PmtDt"`

	// Date at which assets become available to the account owner (in a credit entry), or cease to be available to the account owner (in a debit entry).
	ValueDate *DateFormat11Choice `xml:"ValDt,omitempty"`

	// Date/time at which a foreign exchange rate will be determined.
	ForeignExchangeRateFixingDate *DateFormat6Choice `xml:"FXRateFxgDt,omitempty"`

	// Date on which a payment can be made, for example, if payment date is a non-business day or to indicate the first payment date of an offer.
	EarliestPaymentDate *DateFormat6Choice `xml:"EarlstPmtDt,omitempty"`
}

func (c *CorporateActionDate9) AddPaymentDate() *DateFormat6Choice {
	c.PaymentDate = new(DateFormat6Choice)
	return c.PaymentDate
}

func (c *CorporateActionDate9) AddValueDate() *DateFormat11Choice {
	c.ValueDate = new(DateFormat11Choice)
	return c.ValueDate
}

func (c *CorporateActionDate9) AddForeignExchangeRateFixingDate() *DateFormat6Choice {
	c.ForeignExchangeRateFixingDate = new(DateFormat6Choice)
	return c.ForeignExchangeRateFixingDate
}

func (c *CorporateActionDate9) AddEarliestPaymentDate() *DateFormat6Choice {
	c.EarliestPaymentDate = new(DateFormat6Choice)
	return c.EarliestPaymentDate
}
