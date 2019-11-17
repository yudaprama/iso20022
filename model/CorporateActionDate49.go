package model

// Specifies corporate action dates.
type CorporateActionDate49 struct {

	// Date/time at which positions are struck at the end of the day to note which parties will receive the relevant amount of entitlement, due to be distributed on payment date.
	RecordDate *DateFormat31Choice `xml:"RcrdDt,omitempty"`

	// Date/time as from which trading (including exchange and OTC trading) occurs on the underlying security without the benefit.
	ExDividendDate *DateFormat31Choice `xml:"ExDvddDt,omitempty"`
}

func (c *CorporateActionDate49) AddRecordDate() *DateFormat31Choice {
	c.RecordDate = new(DateFormat31Choice)
	return c.RecordDate
}

func (c *CorporateActionDate49) AddExDividendDate() *DateFormat31Choice {
	c.ExDividendDate = new(DateFormat31Choice)
	return c.ExDividendDate
}
