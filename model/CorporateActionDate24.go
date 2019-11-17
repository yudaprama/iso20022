package model

// Specifies corporate action dates.
type CorporateActionDate24 struct {

	// Date/Time of the posting (credit or debit) to the account.
	PostingDate *DateAndDateTimeChoice `xml:"PstngDt"`

	// Date/time when calculating economic benefit for a cash amount.
	ValueDate *DateAndDateTimeChoice `xml:"ValDt,omitempty"`

	// Date/time at which a foreign exchange rate will be determined.
	ForeignExchangeRateFixingDate *DateAndDateTimeChoice `xml:"FXRateFxgDt,omitempty"`

	// Date on which a payment can be made, for example, if payment date is a non-business day or to indicate the first payment date of an offer.
	EarliestPaymentDate *DateAndDateTimeChoice `xml:"EarlstPmtDt,omitempty"`

	// Date on which the distribution is due to take place (cash and/or securities).
	PaymentDate *DateAndDateTimeChoice `xml:"PmtDt,omitempty"`
}

func (c *CorporateActionDate24) AddPostingDate() *DateAndDateTimeChoice {
	c.PostingDate = new(DateAndDateTimeChoice)
	return c.PostingDate
}

func (c *CorporateActionDate24) AddValueDate() *DateAndDateTimeChoice {
	c.ValueDate = new(DateAndDateTimeChoice)
	return c.ValueDate
}

func (c *CorporateActionDate24) AddForeignExchangeRateFixingDate() *DateAndDateTimeChoice {
	c.ForeignExchangeRateFixingDate = new(DateAndDateTimeChoice)
	return c.ForeignExchangeRateFixingDate
}

func (c *CorporateActionDate24) AddEarliestPaymentDate() *DateAndDateTimeChoice {
	c.EarliestPaymentDate = new(DateAndDateTimeChoice)
	return c.EarliestPaymentDate
}

func (c *CorporateActionDate24) AddPaymentDate() *DateAndDateTimeChoice {
	c.PaymentDate = new(DateAndDateTimeChoice)
	return c.PaymentDate
}
