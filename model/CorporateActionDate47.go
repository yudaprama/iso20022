package model

// Specifies corporate action dates.
type CorporateActionDate47 struct {

	// Date on which the movement is due to take place (cash and/or securities).
	PaymentDate *DateFormat31Choice `xml:"PmtDt"`

	// Date/time when calculating economic benefit for a cash amount.
	ValueDate *DateFormat33Choice `xml:"ValDt,omitempty"`

	// Date/time at which a foreign exchange rate will be determined.
	ForeignExchangeRateFixingDate *DateFormat31Choice `xml:"FXRateFxgDt,omitempty"`

	// Date on which a payment can be made, for example, if payment date is a non-business day or to indicate the first payment date of an offer.
	EarliestPaymentDate *DateFormat31Choice `xml:"EarlstPmtDt,omitempty"`
}

func (c *CorporateActionDate47) AddPaymentDate() *DateFormat31Choice {
	c.PaymentDate = new(DateFormat31Choice)
	return c.PaymentDate
}

func (c *CorporateActionDate47) AddValueDate() *DateFormat33Choice {
	c.ValueDate = new(DateFormat33Choice)
	return c.ValueDate
}

func (c *CorporateActionDate47) AddForeignExchangeRateFixingDate() *DateFormat31Choice {
	c.ForeignExchangeRateFixingDate = new(DateFormat31Choice)
	return c.ForeignExchangeRateFixingDate
}

func (c *CorporateActionDate47) AddEarliestPaymentDate() *DateFormat31Choice {
	c.EarliestPaymentDate = new(DateFormat31Choice)
	return c.EarliestPaymentDate
}
