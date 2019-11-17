package model

// Specifies corporate action dates.
type CorporateActionDate45 struct {

	// Date/time at which positions are struck at the end of the day to note which parties will receive the relevant amount of entitlement, due to be distributed on payment date.
	RecordDate *DateFormat31Choice `xml:"RcrdDt,omitempty"`

	// Date/time as from which trading (including exchange and OTC trading) occurs on the underlying security without the benefit.
	ExDividendDate *DateFormat31Choice `xml:"ExDvddDt,omitempty"`

	// Date/time on which the lottery is run and applied to the holder's positions. This is also applicable to partial calls.
	LotteryDate *DateFormat31Choice `xml:"LtryDt,omitempty"`
}

func (c *CorporateActionDate45) AddRecordDate() *DateFormat31Choice {
	c.RecordDate = new(DateFormat31Choice)
	return c.RecordDate
}

func (c *CorporateActionDate45) AddExDividendDate() *DateFormat31Choice {
	c.ExDividendDate = new(DateFormat31Choice)
	return c.ExDividendDate
}

func (c *CorporateActionDate45) AddLotteryDate() *DateFormat31Choice {
	c.LotteryDate = new(DateFormat31Choice)
	return c.LotteryDate
}
