package model

// Specifies coprorate action dates.
type CorporateActionDate3 struct {

	// Date/time at which the distribution is due to take place (cash and/or securities).
	PaymentDate *DateFormat4Choice `xml:"PmtDt,omitempty"`

	// Date/time at which securities become available for sale.
	AvailableDate *DateFormat4Choice `xml:"AvlblDt,omitempty"`

	// Date/time at which a security will be entitled to a dividend.
	DividendRankingDate *DateFormat4Choice `xml:"DvddRnkgDt,omitempty"`

	// Date on which security will assimilate, become fungible, or have the same rights to dividends as the parent issue.
	PariPassuDate *DateFormat4Choice `xml:"PrpssDt,omitempty"`

	// Date/time at which new securities begin trading.
	FirstDealingDate *DateFormat4Choice `xml:"FrstDealgDt,omitempty"`

	// Date/time at which a payment can be made, eg, if payment date is a non-business day or to indicate the first payment date of an offer.
	EarliestPaymentDate *DateFormat4Choice `xml:"EarlstPmtDt,omitempty"`
}

func (c *CorporateActionDate3) AddPaymentDate() *DateFormat4Choice {
	c.PaymentDate = new(DateFormat4Choice)
	return c.PaymentDate
}

func (c *CorporateActionDate3) AddAvailableDate() *DateFormat4Choice {
	c.AvailableDate = new(DateFormat4Choice)
	return c.AvailableDate
}

func (c *CorporateActionDate3) AddDividendRankingDate() *DateFormat4Choice {
	c.DividendRankingDate = new(DateFormat4Choice)
	return c.DividendRankingDate
}

func (c *CorporateActionDate3) AddPariPassuDate() *DateFormat4Choice {
	c.PariPassuDate = new(DateFormat4Choice)
	return c.PariPassuDate
}

func (c *CorporateActionDate3) AddFirstDealingDate() *DateFormat4Choice {
	c.FirstDealingDate = new(DateFormat4Choice)
	return c.FirstDealingDate
}

func (c *CorporateActionDate3) AddEarliestPaymentDate() *DateFormat4Choice {
	c.EarliestPaymentDate = new(DateFormat4Choice)
	return c.EarliestPaymentDate
}
