package model

// Specifies corporate action dates.
type CorporateActionDate50 struct {

	// Date/time at which positions are struck at the end of the day to note which parties will receive the relevant amount of entitlement, due to be distributed on payment date.
	RecordDate *DateFormat34Choice `xml:"RcrdDt,omitempty"`

	// Date/time as from which trading (including exchange and OTC trading) occurs on the underlying security without the benefit.
	ExDividendDate *DateFormat34Choice `xml:"ExDvddDt,omitempty"`
}

func (c *CorporateActionDate50) AddRecordDate() *DateFormat34Choice {
	c.RecordDate = new(DateFormat34Choice)
	return c.RecordDate
}

func (c *CorporateActionDate50) AddExDividendDate() *DateFormat34Choice {
	c.ExDividendDate = new(DateFormat34Choice)
	return c.ExDividendDate
}
