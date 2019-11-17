package model

// Closing balance for the statement period (final closing balance) or of this page (intermediary closing balance).
type ClosingBalance3 struct {

	// Indication that the position is short or long.
	ShortLongIndicator *ShortLong1Code `xml:"ShrtLngInd"`

	// Closing balance for the statement period (final closing balance) or of this page (intermediary closing balance).
	ClosingBalance *ClosingBalance4Choice `xml:"ClsgBal"`
}

func (c *ClosingBalance3) SetShortLongIndicator(value string) {
	c.ShortLongIndicator = (*ShortLong1Code)(&value)
}

func (c *ClosingBalance3) AddClosingBalance() *ClosingBalance4Choice {
	c.ClosingBalance = new(ClosingBalance4Choice)
	return c.ClosingBalance
}
