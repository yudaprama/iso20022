package model

// Specifies corporate action dates.
type CorporateActionDate54 struct {

	// Date/time at which positions are struck at the end of the day to note which parties will receive the relevant amount of entitlement, due to be distributed on payment date.
	RecordDate *DateFormat34Choice `xml:"RcrdDt,omitempty"`

	// Date/time as from which trading (including exchange and OTC trading) occurs on the underlying security without the benefit.
	ExDividendDate *DateFormat34Choice `xml:"ExDvddDt,omitempty"`

	// Date/time on which the lottery is run and applied to the holder's positions. This is also applicable to partial calls.
	LotteryDate *DateFormat34Choice `xml:"LtryDt,omitempty"`
}

func (c *CorporateActionDate54) AddRecordDate() *DateFormat34Choice {
	c.RecordDate = new(DateFormat34Choice)
	return c.RecordDate
}

func (c *CorporateActionDate54) AddExDividendDate() *DateFormat34Choice {
	c.ExDividendDate = new(DateFormat34Choice)
	return c.ExDividendDate
}

func (c *CorporateActionDate54) AddLotteryDate() *DateFormat34Choice {
	c.LotteryDate = new(DateFormat34Choice)
	return c.LotteryDate
}
