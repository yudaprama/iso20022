package model

// Specifies corporate action dates.
type CorporateActionDate41 struct {

	// Date/time at which positions are struck at the end of the day to note which parties will receive the relevant amount of entitlement, due to be distributed on payment date.
	RecordDate *DateFormat19Choice `xml:"RcrdDt,omitempty"`

	// Date/time as from which trading (including exchange and OTC trading) occurs on the underlying security without the benefit.
	ExDividendDate *DateFormat19Choice `xml:"ExDvddDt,omitempty"`

	// Date/time on which the lottery is run and applied to the holder's positions. This is also applicable to partial calls.
	LotteryDate *DateFormat19Choice `xml:"LtryDt,omitempty"`
}

func (c *CorporateActionDate41) AddRecordDate() *DateFormat19Choice {
	c.RecordDate = new(DateFormat19Choice)
	return c.RecordDate
}

func (c *CorporateActionDate41) AddExDividendDate() *DateFormat19Choice {
	c.ExDividendDate = new(DateFormat19Choice)
	return c.ExDividendDate
}

func (c *CorporateActionDate41) AddLotteryDate() *DateFormat19Choice {
	c.LotteryDate = new(DateFormat19Choice)
	return c.LotteryDate
}
