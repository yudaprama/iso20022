package model

// Specifies corporate action dates.
type CorporateActionDate7 struct {

	// Date/Time of the posting (credit or debit) to the account.
	PostingDate *DateAndDateTimeChoice `xml:"PstngDt"`

	// Date/time at which assets become available to the account owner (in a credit entry), or cease to be available to the account owner (in a debit entry).
	ValueDate *DateAndDateTimeChoice `xml:"ValDt,omitempty"`

	// Date/time at which a foreign exchange rate will be determined.
	ForeignExchangeRateFixingDate *DateAndDateTimeChoice `xml:"FXRateFxgDt,omitempty"`

	// Date on which a payment can be made, for example, if payment date is a non-business day or to indicate the first payment date of an offer.
	EarliestPaymentDate *DateAndDateTimeChoice `xml:"EarlstPmtDt,omitempty"`

	// Date on which the distribution is due to take place (cash and/or securities).
	PaymentDate *DateAndDateTimeChoice `xml:"PmtDt,omitempty"`
}

func (c *CorporateActionDate7) AddPostingDate() *DateAndDateTimeChoice {
	c.PostingDate = new(DateAndDateTimeChoice)
	return c.PostingDate
}

func (c *CorporateActionDate7) AddValueDate() *DateAndDateTimeChoice {
	c.ValueDate = new(DateAndDateTimeChoice)
	return c.ValueDate
}

func (c *CorporateActionDate7) AddForeignExchangeRateFixingDate() *DateAndDateTimeChoice {
	c.ForeignExchangeRateFixingDate = new(DateAndDateTimeChoice)
	return c.ForeignExchangeRateFixingDate
}

func (c *CorporateActionDate7) AddEarliestPaymentDate() *DateAndDateTimeChoice {
	c.EarliestPaymentDate = new(DateAndDateTimeChoice)
	return c.EarliestPaymentDate
}

func (c *CorporateActionDate7) AddPaymentDate() *DateAndDateTimeChoice {
	c.PaymentDate = new(DateAndDateTimeChoice)
	return c.PaymentDate
}
