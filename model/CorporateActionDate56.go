package model

// Specifies corporate action dates.
type CorporateActionDate56 struct {

	// Date on which the movement is due to take place (cash and/or securities).
	PaymentDate *DateFormat34Choice `xml:"PmtDt"`

	// Date/time when calculating economic benefit for a cash amount.
	ValueDate *DateFormat39Choice `xml:"ValDt,omitempty"`

	// Date/time at which a foreign exchange rate will be determined.
	ForeignExchangeRateFixingDate *DateFormat34Choice `xml:"FXRateFxgDt,omitempty"`

	// Date on which a payment can be made, for example, if payment date is a non-business day or to indicate the first payment date of an offer.
	EarliestPaymentDate *DateFormat34Choice `xml:"EarlstPmtDt,omitempty"`
}

func (c *CorporateActionDate56) AddPaymentDate() *DateFormat34Choice {
	c.PaymentDate = new(DateFormat34Choice)
	return c.PaymentDate
}

func (c *CorporateActionDate56) AddValueDate() *DateFormat39Choice {
	c.ValueDate = new(DateFormat39Choice)
	return c.ValueDate
}

func (c *CorporateActionDate56) AddForeignExchangeRateFixingDate() *DateFormat34Choice {
	c.ForeignExchangeRateFixingDate = new(DateFormat34Choice)
	return c.ForeignExchangeRateFixingDate
}

func (c *CorporateActionDate56) AddEarliestPaymentDate() *DateFormat34Choice {
	c.EarliestPaymentDate = new(DateFormat34Choice)
	return c.EarliestPaymentDate
}
